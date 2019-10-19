package handler

import (
	"compress/gzip"
	"net"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/formatter"
	"github.com/gol4ng/logger/socket"
	"github.com/gol4ng/logger/writer"
)

// Gelf will send Entry into gelf writer(depending on network, address)
func Gelf(network string, address string) logger.HandlerInterface {
	switch network {
	case "tcp", "tcp4", "tcp6":
		return GelfTCP(network, address)
	case "udp", "udp4", "udp6":
		return GelfUDP(network, address)
	}
	panic("gelf protocol only supports udp and tcp connections")
}

// GelfFromConnection will send Entry into gelf writer(depending on GelfFromConnection)
func GelfFromConnection(connection net.Conn) logger.HandlerInterface {
	switch connection.(type) {
	case *net.TCPConn:
		return GelfTCPFromConnection(connection.(*net.TCPConn))
	case *net.UDPConn:
		return GelfUDPFromConnection(connection.(*net.UDPConn))
	}
	panic("gelf protocol only supports udp and tcp connections")
}

// GelfTCP will send Entry into TCP gelf writer(depending on network, address)
//
// use GelfTCPFromConnection func if you want to handle the connection error and Closure
func GelfTCP(network string, address string) logger.HandlerInterface {
	return GelfTCPFromConnection(socket.TCPConnection(network, address))
}

// GelfTCPFromConnection will return GelfTCP socket with the current TCPConn
func GelfTCPFromConnection(connection *net.TCPConn) logger.HandlerInterface {
	return Socket(
		connection,
		formatter.NewGelfTCP(),
	)
}

// GelfUDP will send Entry into UDP gelf writer(depending on network, address)
//
// use GelfUDPFromConnection func if you want to handle the connection error and Closure
//
// CAUTION: Logstash only supports Gzip compression
// https://github.com/elastic/logstash/issues/2387
func GelfUDP(network string, address string) logger.HandlerInterface {
	return GelfUDPFromConnection(socket.UDPConnection(network, address))
}

// GelfUDPFromConnection will return GelfUDP socket with the current TCPConn
func GelfUDPFromConnection(connection *net.UDPConn) logger.HandlerInterface {
	return Stream(
		writer.NewCompressWriter(writer.NewGelfChunkWriter(connection), writer.CompressionType(writer.CompressGzip), writer.CompressionLevel(gzip.BestSpeed)),
		formatter.NewGelf(),
	)
}
