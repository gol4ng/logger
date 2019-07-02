package handler_test

import (
	"net"
	"os"
	"reflect"
	"testing"
	"time"

	"bou.ke/monkey"

	"github.com/stretchr/testify/assert"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/handler"
)

func initTCP(t *testing.T, netw string, addr string, data []byte) func() {
	var TCPAddr *net.TCPAddr
	var TCPConn *net.TCPConn
	monkey.Patch(net.ResolveTCPAddr, func(network, address string) (*net.TCPAddr, error) {
		assert.Equal(t, netw, network)
		assert.Equal(t, addr, address)
		return TCPAddr, nil
	})
	monkey.Patch(net.DialTCP, func(network string, laddr, raddr *net.TCPAddr) (*net.TCPConn, error) {
		assert.Equal(t, netw, network)
		assert.Nil(t, laddr)
		assert.Equal(t, TCPAddr, raddr)
		return TCPConn, nil
	})
	monkey.PatchInstanceMethod(reflect.TypeOf(TCPConn), "Write", func(conn *net.TCPConn, b []byte) (n int, err error) {
		assert.Equal(t, data, b)
		return 99, nil
	})
	monkey.Patch(time.Now, func() time.Time { return time.Unix(513216000, 0) })
	monkey.Patch(os.Hostname, func() (string, error) { return "my_fake_hostname", nil })

	return monkey.UnpatchAll
}

func initUDP(t *testing.T, netw string, addr string, data []byte) func() {
	var UDPAddr *net.UDPAddr
	var UDPConn *net.UDPConn
	monkey.Patch(net.ResolveUDPAddr, func(network, address string) (*net.UDPAddr, error) {
		assert.Equal(t, netw, network)
		assert.Equal(t, addr, address)
		return UDPAddr, nil
	})
	monkey.Patch(net.DialUDP, func(network string, laddr, raddr *net.UDPAddr) (*net.UDPConn, error) {
		assert.Equal(t, netw, network)
		assert.Nil(t, laddr)
		assert.Equal(t, UDPAddr, raddr)
		return UDPConn, nil
	})
	monkey.PatchInstanceMethod(reflect.TypeOf(UDPConn), "Write", func(conn *net.UDPConn, b []byte) (n int, err error) {
		assert.Equal(t, data, b)
		return 99, nil
	})
	monkey.Patch(time.Now, func() time.Time { return time.Unix(513216000, 0) })
	monkey.Patch(os.Hostname, func() (string, error) { return "my_fake_hostname", nil })
	return monkey.UnpatchAll
}

func TestGelf_TCP_Handle(t *testing.T) {
	defer initTCP(t, "tcp", "fake_address", []uint8("{\"version\":\"1.1\",\"host\":\"my_fake_hostname\",\"level\":4,\"timestamp\":513216000.000,\"short_message\":\"test message\",\"full_message\":\"<warning> test message\"}\x00"))()
	h := handler.Gelf("tcp", "fake_address")

	assert.Nil(t, h(logger.Entry{Message: "test message", Level: logger.WarningLevel, Context: nil}))
}

func TestGelf_UDP_Handle(t *testing.T) {
	defer initUDP(t, "udp", "fake_address", []byte{0x1f, 0x8b, 0x8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x4, 0xff, 0x4, 0xc0, 0xc1, 0x8a, 0xc2, 0x30, 0x10, 0x6, 0xe0, 0xfb, 0x3e, 0x45, 0xf9, 0xcf, 0xa1, 0x24, 0xbb, 0xab, 0x87, 0x20, 0xbe, 0x4a, 0x98, 0xc3, 0xdf, 0xb4, 0x98, 0x49, 0xa4, 0x33, 0x56, 0x44, 0x7c, 0x77, 0xbf, 0x37, 0xe, 0xee, 0xb6, 0x8d, 0x8e, 0x8c, 0x34, 0x27, 0x4, 0xac, 0xc3, 0x1c, 0x19, 0xfa, 0x2a, 0x8b, 0xdc, 0x58, 0xd6, 0x61, 0xde, 0x45, 0x89, 0x80, 0xc6, 0x83, 0xd, 0xf9, 0x3f, 0xc0, 0x37, 0xa5, 0xb9, 0xe8, 0x1d, 0xf9, 0x94, 0xfe, 0x7e, 0xd3, 0x39, 0xc6, 0x38, 0xc7, 0x18, 0x3, 0x6c, 0x1d, 0xbb, 0x17, 0xa5, 0x99, 0x54, 0x22, 0xc3, 0x69, 0x3e, 0x29, 0xcd, 0xa4, 0x12, 0x1, 0xcb, 0xa3, 0xb5, 0xa2, 0x34, 0x93, 0x4a, 0x64, 0x5c, 0x9e, 0xb2, 0xf7, 0xad, 0xd7, 0xeb, 0xe4, 0x34, 0x9f, 0x94, 0x66, 0x52, 0x89, 0xcf, 0xcf, 0x37, 0x0, 0x0, 0xff, 0xff, 0x82, 0xfb, 0xfd, 0x81, 0x97, 0x0, 0x0, 0x0})()
	h := handler.Gelf("udp", "fake_address")

	assert.Nil(t, h(logger.Entry{Message: "test message", Level: logger.WarningLevel, Context: nil}))
}

func TestGelfTCP_Handle(t *testing.T) {
	defer initTCP(t, "fake_network", "fake_address", []uint8("{\"version\":\"1.1\",\"host\":\"my_fake_hostname\",\"level\":4,\"timestamp\":513216000.000,\"short_message\":\"test message\",\"full_message\":\"<warning> test message\"}\x00"))()
	h := handler.GelfTCP("fake_network", "fake_address")

	assert.Nil(t, h(logger.Entry{Message: "test message", Level: logger.WarningLevel, Context: nil}))
}

func TestGelfUDP_Handle(t *testing.T) {
	defer initUDP(t, "fake_network", "fake_address", []byte{0x1f, 0x8b, 0x8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x4, 0xff, 0x4, 0xc0, 0xc1, 0x8a, 0xc2, 0x30, 0x10, 0x6, 0xe0, 0xfb, 0x3e, 0x45, 0xf9, 0xcf, 0xa1, 0x24, 0xbb, 0xab, 0x87, 0x20, 0xbe, 0x4a, 0x98, 0xc3, 0xdf, 0xb4, 0x98, 0x49, 0xa4, 0x33, 0x56, 0x44, 0x7c, 0x77, 0xbf, 0x37, 0xe, 0xee, 0xb6, 0x8d, 0x8e, 0x8c, 0x34, 0x27, 0x4, 0xac, 0xc3, 0x1c, 0x19, 0xfa, 0x2a, 0x8b, 0xdc, 0x58, 0xd6, 0x61, 0xde, 0x45, 0x89, 0x80, 0xc6, 0x83, 0xd, 0xf9, 0x3f, 0xc0, 0x37, 0xa5, 0xb9, 0xe8, 0x1d, 0xf9, 0x94, 0xfe, 0x7e, 0xd3, 0x39, 0xc6, 0x38, 0xc7, 0x18, 0x3, 0x6c, 0x1d, 0xbb, 0x17, 0xa5, 0x99, 0x54, 0x22, 0xc3, 0x69, 0x3e, 0x29, 0xcd, 0xa4, 0x12, 0x1, 0xcb, 0xa3, 0xb5, 0xa2, 0x34, 0x93, 0x4a, 0x64, 0x5c, 0x9e, 0xb2, 0xf7, 0xad, 0xd7, 0xeb, 0xe4, 0x34, 0x9f, 0x94, 0x66, 0x52, 0x89, 0xcf, 0xcf, 0x37, 0x0, 0x0, 0xff, 0xff, 0x82, 0xfb, 0xfd, 0x81, 0x97, 0x0, 0x0, 0x0})()
	h := handler.GelfUDP("fake_network", "fake_address")

	assert.Nil(t, h(logger.Entry{Message: "test message", Level: logger.WarningLevel, Context: nil}))
}
