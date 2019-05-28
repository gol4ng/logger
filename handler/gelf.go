package handler

import (
	"github.com/gol4ng/logger/formatter"
)

func NewGelfTCP(network string, address string) (*Socket, error) {
	return NewTCPSocket(network, address, formatter.NewGelfEncoder())
}

// TODO
//func NewGelfUDP(network string, address string) (*Socket, error) {
//	return NewUDPSocket(network, address, formatter.NewGelfEncoder())
//}
