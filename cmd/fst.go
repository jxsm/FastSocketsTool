package main

import (
	"FastSocketsTool/pkg/option"
	"FastSocketsTool/pkg/start"
)

func main() {
	operationOption := option.InitOperationOption()
	start.Start(operationOption)
}
