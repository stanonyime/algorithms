package algorithms

import (
	"testing"
)

func TestQueue(t *testing.T) {
	s := new(Queue)

	t.Run("Enqueue", func(t *testing.T) {
		s.Enqueue(1)
		expects := 1

		if len(s.items) != expects {
			t.Errorf("got %v, want %v", len(s.items), expects)
		}

		s.Enqueue(2)
		expects = 2

		if len(s.items) != expects {
			t.Errorf("got %v, want %v", len(s.items), expects)
		}
	})

	t.Run("Dequeue", func(t *testing.T) {
		got, _ := s.Dequeue()
		expects := 1

		if got != expects {
			t.Errorf("got %v, want %v", got, expects)
		}

		got, _ = s.Dequeue()
		expects = 2

		if got != expects {
			t.Errorf("got %v, want %v", got, expects)
		}
	})

}
