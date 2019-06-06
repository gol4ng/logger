package handler

import (
	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/formatter"
	"github.com/gol4ng/logger/socket"
	"net"
)

func Gelf(network string, addrerss string) logger.HandlerInterface {
	switch network {
	case "tcp", "tcp4", "tcp6":
		return GelfTCP(network, addrerss)
	case "udp", "udp4", "udp6":
		//TODO implement GelfUDP
		panic("gelf UDP transport is not implemented yet")
	}
	panic("gelf protocol only supports udp and tcp connections")
}

// handles gelf message formatting and automatically chose the correct transport (UDP or TCP)
// regarding the provided connection
//
// connection errors can be handled before calling the func
// use Gelf func if you want your program to panic on connection errors
func GelfFromConnection(connection net.Conn) logger.HandlerInterface {
	switch connection.(type) {
	case *net.TCPConn:
		return GelfTCPFromConnection(connection.(*net.TCPConn))
	case *net.UDPConn:
		//TODO implement UDP chuncks
		panic("gelf UDP transport is not implemented yet")
	}
	panic("gelf protocol only supports udp and tcp connections")
}

// helper that creates a TCP connection and return a handlerInterface
// CAUTION: you will not be able to handle connection errors using that func
//          as it panics if their is any error during the connection
//
// use Gelf func if you want to handle the connection error and Closure
func GelfTCP(network string, address string) logger.HandlerInterface {
	return GelfTCPFromConnection(socket.TCPConnection(network, address))
}

func GelfTCPFromConnection(connection *net.TCPConn) logger.HandlerInterface {
	return Socket(
		connection,
		formatter.NewGelfTCP(),
	)
}
