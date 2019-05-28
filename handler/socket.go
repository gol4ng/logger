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

func NewTCPSocket(network string, address string, formatter logger.FormatterInterface) (*Socket, error) {
	tcpAddr, err := net.ResolveTCPAddr(network, address)
	if err != nil {
		return nil, err
	}

	connection, err := net.DialTCP(network, nil, tcpAddr)
	if err != nil {
		return nil, err
	}

	return &Socket{connection: connection, formatter: formatter}, nil
}

func NewUDPSocket(network string, address string, formatter logger.FormatterInterface) (*Socket, error) {
	udpAddr, err := net.ResolveUDPAddr(network, address)
	if err != nil {
		return nil, err
	}

	connection, err := net.DialUDP(network, nil, udpAddr)
	if err != nil {
		return nil, err
	}

	return &Socket{connection: connection, formatter: formatter}, nil
}
