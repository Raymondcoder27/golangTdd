package adt

type Stack struct {
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Empty() bool {
	return true
}
