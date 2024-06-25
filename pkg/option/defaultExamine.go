package option

type DefaultExamine struct {
}

// ExamineOperationOption Default check operation
func (d DefaultExamine) ExamineOperationOption(info *OperationOption) bool {
	isErr := false
	if *info.ConnectionAddress == "" {
		Remind("ConnectionUrlIsEmpty")
		isErr = true
	}

	if *info.Port == 0 {
		Remind("PortIsEmpty")
		isErr = true
	}

	if *info.Port >= 65536 {
		Remind("PortIsOutOfRange")
		isErr = true
	}

	return isErr
}
