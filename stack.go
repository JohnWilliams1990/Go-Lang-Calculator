package stack

import "errors"

var Underflow = errors.New("stack underflow");

type cell struct {
    next *cell
	value interface{}
}

type Stack struct {
    top *cell
}

func New() Stack {
    return Stack{nil}
}

func (s *Stack) Push(v interface{}) error {
    s.top = &cell{s.top, v}
	return nil
}
func (s *Stack) Pop() (interface{}, error) {
    if s.top == nil {
	    return nil, Underflow
	}
	v := s.top.value
	s.top = s.top.next
	return v, nil
}

func (s Stack) Top() interface{} {
    if s.top == nil {
	    return nil
	}
	return s.top.value
}

func (s Stack) IsEmpty() bool {
    return s.top == nil
}