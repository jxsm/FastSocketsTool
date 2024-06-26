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
	Server            *bool
	ServerEncode      *string
}

type ExamineOperationOption interface {
	ExamineOperationOption(info *OperationOption) bool
}

// InitOperationOption Initialization Creation OperationOption
func InitOperationOption() *OperationOption {
	data := &OperationOption{}
	b := flag.Bool("u", false, "udp Mod")
	data.ConnectionType = "tcp"

	data.ConnectionAddress = flag.String("h", "", "Connection address")
	data.Encode = flag.String("e", "utf-8", " The encoding used at the time of sending")
	data.Port = flag.Int("p", 0, "Port Number - 1 to 65535")
	data.Server = flag.Bool("s", false, "Server Mode")
	data.ServerEncode = flag.String("re", "utf-8", "The server-side information will be output through this encoding")
	flag.Parse()
	if *b {
		data.ConnectionType = "udp"
	}

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
