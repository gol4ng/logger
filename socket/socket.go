package socket

import "net"

// TCPConnection will return a TCPConn or panic
func TCPConnection(network string, address string) *net.TCPConn {
	tcpAddr, err := net.ResolveTCPAddr(network, address)
	if err != nil {
		panic(err)
	}
	connection, err := net.DialTCP(network, nil, tcpAddr)
	if err != nil {
		panic(err)
	}
	return connection
}

// UDPConnection will return a UDPConn or panic
func UDPConnection(network string, address string) *net.UDPConn {
	udpAddr, err := net.ResolveUDPAddr(network, address)
	if err != nil {
		panic(err)
	}
	connection, err := net.DialUDP(network, nil, udpAddr)
	if err != nil {
		panic(err)
	}
	return connection
}
