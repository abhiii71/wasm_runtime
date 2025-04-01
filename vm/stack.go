package vm

import "errors"

type Stack struct {
	data []int
}

// NewStack creates a new empty stack
func NewStack() *Stack {
	return &Stack{data: []int{}}
}

// Push adds a value to the stack
func (s *Stack) Push(value int) {
	s.data = append(s.data, value)
}

// Pop removes and returns the top value of the stack
func (s *Stack) Pop() (int, error) {
	if len(s.data) == 0 {
		return 0, errors.New("stack underflow")
	}
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val, nil
}

func (s *Stack) Peek() (int, error) {
	if len(s.data) == 0 {
		return 0, errors.New("Stack is empty")
	}
	return s.data[len(s.data)-1], nil
}
