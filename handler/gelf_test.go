package handler_test

import (
	"io"
	"net"
	"os"
	"reflect"
	"testing"
	"time"

	"bou.ke/monkey"

	"github.com/stretchr/testify/assert"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/handler"
	"github.com/gol4ng/logger/writer"
)

func TestGelf(t *testing.T) {
	tests := []struct {
		name    string
		network string
		address string
		method  func(network string, address string) logger.HandlerInterface
	}{
		// TCPs
		{name: `with_tcp_network`, network: "tcp", address: "fake addr", method: handler.GelfTCP},
		{name: `with_tcp4_network`, network: "tcp4", address: "fake addr", method: handler.GelfTCP},
		{name: `with_tcp6_network`, network: "tcp6", address: "fake addr", method: handler.GelfTCP},
		// UDPs
		{name: `with_udp_network`, network: "udp", address: "fake addr", method: handler.GelfUDP},
		{name: `with_udp4_network`, network: "udp4", address: "fake addr", method: handler.GelfUDP},
		{name: `with_udp6_network`, network: "udp6", address: "fake addr", method: handler.GelfUDP},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			monkey.Patch(tt.method, func(network, address string) logger.HandlerInterface {
				assert.Equal(t, network, tt.network)
				assert.Equal(t, address, tt.address)
				return logger.NopHandler
			})
			assert.IsType(t, logger.NopHandler, handler.Gelf(tt.network, tt.address))
			monkey.UnpatchAll()
		})
	}
}

func TestGelf_withWrongNetwork(t *testing.T) {
	assert.PanicsWithValue(t, "gelf protocol only supports udp and tcp connections", func() {
		handler.Gelf("bad network", "fake addr")
	})
}

type WrongConn struct{}

func (w *WrongConn) Read(_ []byte) (n int, err error)   { return 0, nil }
func (w *WrongConn) Write(_ []byte) (n int, err error)  { return 0, nil }
func (w *WrongConn) Close() error                       { return nil }
func (w *WrongConn) LocalAddr() net.Addr                { return nil }
func (w *WrongConn) RemoteAddr() net.Addr              { return nil }
func (w *WrongConn) SetDeadline(_ time.Time) error     { return nil }
func (w *WrongConn) SetReadDeadline(_ time.Time) error { return nil }
func (w *WrongConn) SetWriteDeadline(_ time.Time) error { return nil }

func TestGelfFromConnection(t *testing.T) {

	t.Run("TestGelfFromConnection for TCPConn", func(t *testing.T) {
		monkey.Patch(handler.GelfTCPFromConnection, func(c *net.TCPConn) logger.HandlerInterface {
			//assert.Equal(t, conn, c)
			return logger.NopHandler
		})
		assert.IsType(t, logger.NopHandler, handler.GelfFromConnection(&net.TCPConn{}))
		monkey.UnpatchAll()
	})

	t.Run("TestGelfFromConnection for UDPConn", func(t *testing.T) {
		monkey.Patch(handler.GelfUDPFromConnection, func(c *net.UDPConn) logger.HandlerInterface {
			//assert.Equal(t, conn, c)
			return logger.NopHandler
		})
		assert.IsType(t, logger.NopHandler, handler.GelfFromConnection(&net.UDPConn{}))
		monkey.UnpatchAll()
	})
}

func TestGelfFromConnection_withWrongConnection(t *testing.T) {
	assert.PanicsWithValue(t, "gelf protocol only supports udp and tcp connections", func() {
		handler.GelfFromConnection(&WrongConn{})
	})
}

func TestGelfTCP_Handle(t *testing.T) {
	var TCPAddr *net.TCPAddr
	var TCPConn *net.TCPConn
	monkey.Patch(net.ResolveTCPAddr, func(network, address string) (*net.TCPAddr, error) {
		assert.Equal(t, "fake_network", network)
		assert.Equal(t, "fake_address", address)
		return TCPAddr, nil
	})
	monkey.Patch(net.DialTCP, func(network string, laddr, raddr *net.TCPAddr) (*net.TCPConn, error) {
		assert.Equal(t, "fake_network", network)
		assert.Nil(t, laddr)
		assert.Equal(t, TCPAddr, raddr)
		return TCPConn, nil
	})
	monkey.PatchInstanceMethod(reflect.TypeOf(TCPConn), "Write", func(conn *net.TCPConn, b []byte) (n int, err error) {
		assert.Equal(t, []byte(`{"version":"1.1","host":"my_fake_hostname","level":4,"timestamp":513216000.000,"short_message":"test message","full_message":"<warning> test message"}`+"\x00"), b)
		return 99, nil
	})
	monkey.Patch(time.Now, func() time.Time { return time.Unix(513216000, 0) })
	monkey.Patch(os.Hostname, func() (string, error) { return "my_fake_hostname", nil })
	defer monkey.UnpatchAll()

	h := handler.GelfTCP("fake_network", "fake_address")

	assert.Nil(t, h(logger.Entry{Message: "test message", Level: logger.WarningLevel, Context: nil}))
}

func TestGelfUDP_Handle(t *testing.T) {
	osHostname, _ := os.Hostname()
	var UDPAddr *net.UDPAddr
	var UDPConn *net.UDPConn
	var CompressWriter *writer.CompressWriter

	monkey.Patch(net.ResolveUDPAddr, func(network, address string) (*net.UDPAddr, error) {
		assert.Equal(t, "fake_network", network)
		assert.Equal(t, "fake_address", address)
		return UDPAddr, nil
	})
	monkey.Patch(net.DialUDP, func(network string, laddr, raddr *net.UDPAddr) (*net.UDPConn, error) {
		assert.Equal(t, "fake_network", network)
		assert.Nil(t, laddr)
		assert.Equal(t, UDPAddr, raddr)
		return UDPConn, nil
	})
	monkey.Patch(writer.NewCompressWriter, func(w io.Writer, options ...writer.CompressOption) *writer.CompressWriter {
		assert.IsType(t, &writer.GelfChunkWriter{}, w, "writer passed to CompressWriter must be a GelfCompressWriter")
		assert.Len(t, options, 2)
		return CompressWriter
	})
	monkey.PatchInstanceMethod(reflect.TypeOf(CompressWriter), "Write", func(w *writer.CompressWriter, p []byte) (int, error) {
		assert.Equal(t, []byte(`{"version":"1.1","host":"`+osHostname+`","level":4,"timestamp":513216000.000,"short_message":"test message","full_message":"<warning> test message"}`+"\n"), p)
		return 99, nil
	})
	monkey.PatchInstanceMethod(reflect.TypeOf(UDPConn), "Write", func(conn *net.UDPConn, b []byte) (n int, err error) {
		assert.Equal(t, []byte(`{"version":"1.1","host":"`+osHostname+`","level":4,"timestamp":513216000.000,"short_message":"test message","full_message":"<warning> test message"}`+"\n"), b)
		return 99, nil
	})
	monkey.Patch(time.Now, func() time.Time { return time.Unix(513216000, 0) })
	defer monkey.UnpatchAll()

	h := handler.GelfUDP("fake_network", "fake_address")
	assert.Nil(t, h(logger.Entry{Message: "test message", Level: logger.WarningLevel, Context: nil}))
}
