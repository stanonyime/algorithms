package algorithms

import (
	"errors"
)

// Queue defines a struct for a Queue data structure
type Queue struct {
	items []int
}

func (s *Queue) Enqueue(item int) {
	s.items = append(s.items, item)
}

func (s *Queue) Dequeue() (int, error) {
	if len(s.items) < 1 {
		return 0, errors.New("Queue is empty")
	}
	item := s.items[0]
	s.items = s.items[1:]
	return item, nil
}
