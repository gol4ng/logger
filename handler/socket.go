package handler

import (
	"bytes"
	"net"

	"github.com/gol4ng/logger"
)

type Socket struct {
	connection net.Conn
	formatter  logger.FormatterInterface
}

func (g *Socket) Handle(entry logger.Entry) error {
	buffer := bytes.NewBuffer([]byte(g.formatter.Format(entry)))
	_, err := g.connection.Write(buffer.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func TCPSocket(tcpConn *net.TCPConn, formatter logger.FormatterInterface) *Socket {
	return &Socket{connection: tcpConn, formatter: formatter}
}

func UdpSocket(udpConn *net.UDPConn, formatter logger.FormatterInterface) *Socket {
	return &Socket{connection: udpConn, formatter: formatter}
}
