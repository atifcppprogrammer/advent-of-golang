package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("solves example input correctly", func(t *testing.T) {
		lines := []string{
			"A Y",
			"B X",
			"C Z",
		}

		expected, received := 15, solution(lines)
		if expected != received {
			t.Errorf("expected outcomeSum %q received %q", expected, received)
		}
	})
}
