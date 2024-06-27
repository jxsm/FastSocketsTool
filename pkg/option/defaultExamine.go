package option

import (
	"FastSocketsTool/prompt"
)

type DefaultExamine struct {
}

// ExamineOperationOption Default check operation
func (d DefaultExamine) ExamineOperationOption(info *OperationOption) bool {
	if *info.ConnectionAddress == "" {
		prompt.Prompt("connection_address_IsEmpty")
		return true
	}

	if *info.Port == 0 {
		prompt.Prompt("port_IsEmpty")
		return true
	}

	if *info.Port >= 65536 {
		prompt.Prompt("port_is_OutOfRange")
		return true
	}

	return false
}
