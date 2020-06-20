package algorithms

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := new(Stack)

	t.Run("Stack Push()", func(t *testing.T) {
		s.Push(1)
		expects := 1

		if len(s.items) != expects {
			t.Errorf("got %v, want %v", len(s.items), expects)
		}

		s.Push(2)
		expects = 2

		if len(s.items) != expects {
			t.Errorf("got %v, want %v", len(s.items), expects)
		}
	})

	t.Run("Stack Pop()", func(t *testing.T) {
		got, _ := s.Pop()
		expects := 2

		if got != expects {
			t.Errorf("got %v, want %v", got, expects)
		}

		got, _ = s.Pop()
		expects = 1

		if got != expects {
			t.Errorf("got %v, want %v", got, expects)
		}
	})

}
