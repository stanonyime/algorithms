package algorithms

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	s := new(LinkedList)
	s.Insert(1)
	s.Insert(2)
	s.Insert(3)

	t.Run("LinkedList Insert", func(t *testing.T) {
		if s.head.item != 3 {
			t.Errorf("got %v, want %v", s.head.item, 3)
		}

	})

	t.Run("LinkedList Search", func(t *testing.T) {
		got := s.Search(2)
		expects := 2

		if got.item != expects {
			t.Errorf("got %v, searched for %v", got.item, expects)
		}

		if got.next.item != 1 {
			t.Errorf("got %v, want %v", got.next.item, 1)
		}
	})

	t.Run("LinkedList Delete", func(t *testing.T) {
		//got := s.Search(2)

		s.Delete(2)

		if s.head.next.item != 1 {
			t.Errorf("got %v, want %v", s.head.next.item, 1)
		}

	})

}
