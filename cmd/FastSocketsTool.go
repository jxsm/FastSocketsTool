package main

import (
	"FastSocketsTool/pkg/client"
	"FastSocketsTool/pkg/option"
)

func main() {
	operationOption := option.InitOperationOption()
	client.Start(operationOption)
}
