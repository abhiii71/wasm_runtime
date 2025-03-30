package vm

import "errors"

type Stack struct {
	data []int64
}

func NewStack() *Stack {
	return &Stack{data: make([]int64, 0)}
}

func (s *Stack) Push(value int64) {
	s.data = append(s.data, value)
}

func (s *Stack) Pop() (int64, error) {
	if len(s.data) == 0 {
		return 0, errors.New("stack underflow")
	}
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val, nil
}

func (s *Stack) Peek() (int64, error) {
	if len(s.data) == 0 {
		return 0, errors.New("Stack is empty")
	}
	return s.data[len(s.data)-1], nil
}
