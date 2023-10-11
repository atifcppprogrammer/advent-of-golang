package structures

import (
	"testing"
)

func TestIsEmpty(t *testing.T) {
	t.Run("returns true if list is empty, returns false otherwise", func(t *testing.T) {
		list := List{}
		expected, received := true, list.IsEmpty()
		if expected != received {
			t.Errorf("expected %t received %t", expected, received)
		}

		list.Push('A')
		expected, received = false, list.IsEmpty()
		if expected != received {
			t.Errorf("expected %t received %t", expected, received)
		}
	})
}

func TestLast(t *testing.T) {
	t.Run("returns last element added to list", func(t *testing.T) {
		list := List{}
		elements := []rune{'0', 'A', 'B', 'C', 'D'}
		for _, element := range elements {
			if element != '0' {
				list.Push(element)
			}

			expected, received := element, list.Last()
			if expected != received {
				t.Errorf("expected %q received %q", expected, received)
			}
		}
	})
}

func TestIsEqual(t *testing.T) {
	t.Run("returns true if lists are same", func(t *testing.T) {
		listOne := List{}
		listTwo := List{}

		elements := []rune{'A', 'B', 'C', 'D'}
		for _, element := range elements {
			listOne.Push(element)
			listTwo.Push(element)
		}

		expected := true
		received, _ := listOne.IsEqual(listTwo)
		if expected != received {
			t.Errorf("expected %t received %t", expected, received)
		}
	})

	t.Run("returns false with reason if lists have different lengths", func(t *testing.T) {
		listOne := List{}
		listTwo := List{}

		elements := []rune{'A', 'B', 'C', 'D'}
		for index, element := range elements {
			listOne.Push(element)
			if index == 2 {
				continue
			}
			listTwo.Push(element)
		}

		expected := "len(this) = 4, len(other) = 3"
		_, received := listOne.IsEqual(listTwo)
		if expected != received {
			t.Errorf("expected %q received %q", expected, received)
		}
	})

	t.Run("returns false with reason if lists differ somewhere", func(t *testing.T) {
		listOne := List{}
		listTwo := List{}

		elements := []rune{'A', 'B', 'C', 'D'}
		for index, element := range elements {
			listOne.Push(element)
			if index == 2 {
				listTwo.Push('X')
			} else {
				listTwo.Push(element)
			}
		}

		expected := "this[2] = 'C', other[2] = 'X'"
		_, received := listOne.IsEqual(listTwo)
		if expected != received {
			t.Errorf("expected %q received %q", expected, received)
		}
	})
}

func TestPush(t *testing.T) {
	t.Run("adds elements to list once called", func(t *testing.T) {
		list := List{}
		elements := []rune{'A', 'B', 'C', 'D'}
		for index, element := range elements {
			list.Push(element)
			expected, received := index+1, list.length
			if expected != received {
				t.Errorf("expected %q received %q", expected, received)
			}
		}
	})
}

func TestPop(t *testing.T) {
	t.Run("remove elements from list once called", func(t *testing.T) {
		list := List{}
		list.Push('A')
		list.Push('B')
		list.Push('C')
		list.Push('D')

		for count := 4; count > 0; count-- {
			expected, received := list.GetLength(), count
			if expected != received {
				t.Errorf("expected %q received %q", expected, received)
			}
			list.Pop()
		}
	})
}
