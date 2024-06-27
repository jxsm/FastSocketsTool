package start

import (
	"FastSocketsTool/pkg/client"
	"FastSocketsTool/pkg/option"
	"FastSocketsTool/pkg/server"
)

var clientFunc = client.NewClient
var serverFunc = server.NewServer

func Start(option *option.OperationOption) {
	if *option.Server {
		newServer := serverFunc()
		newServer.Start(option)
	} else {
		newClient := clientFunc()
		newClient.Start(option)
	}
}
