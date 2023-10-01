package main

import (
	"testing"
)

func TestContains(t *testing.T) {
	t.Run("returns true if either pair is contained in the other", func(t *testing.T) {
		pOne, pTwo := Pair{2, 8}, Pair{3, 3}
		expected, received := true, pOne.Contains(pTwo)
		if expected != received {
			t.Errorf("expected %t received %t", expected, received)
		}

		pOne, pTwo = Pair{3, 6}, Pair{5, 6}
		expected, received = true, pOne.Contains(pTwo)
		if expected != received {
			t.Errorf("expected %t received %t", expected, received)
		}
	})

	t.Run("returns false if pair does not contain the other", func(t *testing.T) {
		pOne, pTwo := Pair{2, 7}, Pair{8, 9}
		expected, received := false, pOne.Contains(pTwo)
		if expected != received {
			t.Errorf("expected %t received %t", expected, received)
		}
	})
}
