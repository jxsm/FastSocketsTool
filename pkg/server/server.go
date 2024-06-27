package server

import (
	"FastSocketsTool/pkg/option"
)

type Server struct {
	server IServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start(option *option.OperationOption) {
	if option.ConnectionType == "tcp" {
		s.server = NewTcpServer(option)
	} else {
		s.server = NewUdpServer(option)
	}
	s.server.StartServer()
}

func GetNetworkString(network string, isIpv6 bool) string {
	if network != "tcp" && network != "udp" {
		return network
	}
	if isIpv6 {
		return network + "6"
	}
	return network + "4"
}
