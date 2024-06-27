package server

import "FastSocketsTool/pkg/option"

func NewTcpServer(option *option.OperationOption) IServer {
	return &TCPServer{option: option}
}

func NewUdpServer(option *option.OperationOption) IServer {
	return UdpServer{option}
}
