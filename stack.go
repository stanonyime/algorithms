package algorithms

import (
	"errors"
)

// Stack defines a struct for a stack data structure
type Stack struct {
	items []interface{}
}

func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (interface{}, error) {
	if len(s.items) < 1 {
		return 0, errors.New("Stack is empty")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}
