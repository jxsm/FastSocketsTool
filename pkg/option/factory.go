package option

type FactoryOption struct {
}

func (FactoryOption) NewExamineOperationOption() ExamineOperationOption {
	return DefaultExamine{}
}
