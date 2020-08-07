package StrategyPattern

type Strategy interface {
	DoOperation(num1 int, num2 int) int
}

type OperationAdd struct {
}

func (o *OperationAdd) DoOperation(num1 int, num2 int) int {
	return num1 + num2
}

type OperationSubstract struct {
}

func (o *OperationSubstract) DoOperation(num1 int, num2 int) int {
	return num1 - num2
}

type OperationMultiply struct {
}

func (o *OperationMultiply) DoOperation(num1 int, num2 int) int {
	return num1 * num2
}
