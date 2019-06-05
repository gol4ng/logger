package handler_test

import (
	"net"
	"reflect"
	"testing"

	"bou.ke/monkey"
	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/handler"
	"github.com/gol4ng/logger/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTCPSocket_Handle(t *testing.T) {
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
		assert.Equal(t, []uint8("my formatter return"), b)
		return 99, nil
	})
	defer monkey.UnpatchAll()

	mockFormatter := mocks.FormatterInterface{}
	mockFormatter.On("Format", mock.AnythingOfType("logger.Entry")).Return("my formatter return")

	h, err := handler.TCPSocket("fake_network", "fake_address", &mockFormatter)

	assert.Nil(t, err)
	assert.Nil(t, h.Handle(logger.Entry{Message: "test message", Level: logger.WarningLevel, Context: nil}))
}

func TestUDPSocket_Handle(t *testing.T) {
	var UDPAddr *net.UDPAddr
	var UDPConn *net.UDPConn
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
	monkey.PatchInstanceMethod(reflect.TypeOf(UDPConn), "Write", func(conn *net.UDPConn, b []byte) (n int, err error) {
		assert.Equal(t, []uint8("my formatter return"), b)
		return 99, nil
	})
	defer monkey.UnpatchAll()

	mockFormatter := mocks.FormatterInterface{}
	mockFormatter.On("Format", mock.AnythingOfType("logger.Entry")).Return("my formatter return")

	h, err := handler.UdpSocket("fake_network", "fake_address", &mockFormatter)

	assert.Nil(t, err)
	assert.Nil(t, h.Handle(logger.Entry{Message: "test message", Level: logger.WarningLevel, Context: nil}))
}
