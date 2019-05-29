package handler

import (
	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/formatter"
)

// http://docs.graylog.org/en/3.0/pages/gelf.html#gelf-via-tcp
type GelfTCPFormatter struct {
	formatter *formatter.Gelf
}

func (g *GelfTCPFormatter) Format(entry logger.Entry) string {
	return g.formatter.Format(entry) + "\x00"
}

func NewGelfTCP(network string, address string) (*Socket, error) {
	return NewTCPSocket(network, address, &GelfTCPFormatter{formatter: formatter.NewGelf()})
}

// TODO
//func NewGelfUDP(network string, address string) (*Socket, error) {
//	return NewUDPSocket(network, address, formatter.NewGelf())
//}
