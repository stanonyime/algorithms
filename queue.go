package algorithms

import (
	"errors"
)

// Queue defines a struct for a Queue data structure
type Queue struct {
	items []interface{}
}

func (s *Queue) Enqueue(item interface{}) {
	s.items = append(s.items, item)
}

func (s *Queue) Dequeue() (interface{}, error) {
	if len(s.items) < 1 {
		return 0, errors.New("Queue is empty")
	}
	item := s.items[0]
	s.items = s.items[1:]
	return item, nil
}
