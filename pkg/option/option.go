package option

import (
	"flag"
	"os"
	"strconv"
	"strings"
)

type OperationOption struct {
	ConnectionType    string
	ConnectionAddress *string
	Encode            *string
	Port              *int
}

type ExamineOperationOption interface {
	ExamineOperationOption(info *OperationOption) bool
}

// InitOperationOption Initialization Creation OperationOption
func InitOperationOption() *OperationOption {
	data := &OperationOption{}
	b := flag.Bool("u", false, "udp Mod")
	data.ConnectionType = "tcp"
	if *b {
		data.ConnectionType = "udp"
	}
	data.ConnectionAddress = flag.String("a", "", "Connection Url")
	data.Encode = flag.String("e", "utf-8", "Encode")
	data.Port = flag.Int("p", 0, "Port Number - 1 to 65535")
	flag.Parse()
	examineOperationOption(FactoryOption{}.NewExamineOperationOption(), data)
	return data
}

// examineOperationOption Check the validity of the operation option
func examineOperationOption(info ExamineOperationOption, operationOption *OperationOption) {
	if info.ExamineOperationOption(operationOption) {
		os.Exit(1)
	}
}

// GetConnectionAddress  Get Connection Address
func (operationOption *OperationOption) GetConnectionAddress() string {
	builder := strings.Builder{}
	builder.WriteString(*operationOption.ConnectionAddress)
	builder.WriteString(":")
	builder.WriteString(strconv.Itoa(*operationOption.Port))
	return builder.String()
}
