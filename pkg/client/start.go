package client

import (
	"FastSocketsTool/pkg/option"
	"fmt"
	"net"
)

func Start(option *option.OperationOption) {
	_, err := net.Dial(option.ConnectionType, option.GetConnectionAddress())
	if err != nil {
		fmt.Println("Failed to connect to the server")
	}
}
