package handler

import (
	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/formatter"
	"github.com/gol4ng/logger/socket"
	"github.com/pkg/errors"

	"net"
)

// http://docs.graylog.org/en/3.0/pages/gelf.html#gelf-via-tcp
type GelfTCPFormatter struct {
	formatter *formatter.Gelf
}

func (g *GelfTCPFormatter) Format(entry logger.Entry) string {
	return g.formatter.Format(entry) + "\x00"
}

// helper that creates a TCP connection and return a handlerInterface
// CAUTION: you will not be able to handle connection errors using that func
//          as it panics if their is any error during the connection
//
// use Gelf func if you want to handle the connection error and Closure
func TCPGelf(network string, address string) logger.HandlerInterface {
	return TCPSocket(
		socket.TCPConnection(network, address),
		&GelfTCPFormatter{formatter: formatter.NewGelf()},
	).Handle
}

// handles gelf message formatting and automatically chose the correct transport (UDP or TCP)
// regarding the provided connection
//
// connection errors can be handled before calling the func
// use TCPGelf func if you want your program to panic on connection errors
func Gelf(connection net.Conn) (logger.HandlerInterface, error) {
	switch connection.(type) {
	case *net.TCPConn:
		return TCPSocket(
			connection.(*net.TCPConn),
			&GelfTCPFormatter{formatter: formatter.NewGelf()},
		).Handle, nil
	case *net.UDPConn:
		//TODO implement UDP chuncks
		return nil, errors.New("gelf UDP transport is not implemented yet")
	}
	return nil, errors.New("gelf protocol only supports udp and tcp connections")
}
