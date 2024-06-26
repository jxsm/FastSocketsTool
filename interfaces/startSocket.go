package interfaces

import "FastSocketsTool/pkg/option"

type StartSocket interface {
	Start(option *option.OperationOption)
}
