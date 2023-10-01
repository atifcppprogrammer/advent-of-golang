package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("solves example input correctly", func(t *testing.T) {
		lines := []string{
			"2-4,6-8",
			"2-3,4-5",
			"5-7,7-9",
			"2-8,3-7",
			"6-6,4-6",
			"2-6,4-8",
		}
		received, expected := solution(lines), 2
		if expected != received {
			t.Errorf("Expected pairCount %q but received %q", expected, received)
		}
	})
}

func TestIsPairSatisfactory(t *testing.T) {
	t.Run("returns true if either pair is contained in the other", func(t *testing.T) {
		line := "2-8,3-3"
		expected, received := true, isPairSatisfactory(line)
		if expected != received {
			t.Errorf("expected %t received %t", expected, received)
		}

		line = "3-6,5-6"
		expected, received = true, isPairSatisfactory(line)
		if expected != received {
			t.Errorf("expected %t received %t", expected, received)
		}
	})

	t.Run("returns false if pair does not contain the other", func(t *testing.T) {
		line := "2-7,8-9"
		expected, received := false, isPairSatisfactory(line)
		if expected != received {
			t.Errorf("expected %t received %t", expected, received)
		}
	})
}
