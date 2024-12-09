package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("solves example input correctly", func(t *testing.T) {
		lines := []string{ /* provide example input here */ }

		expected, received := 0, solution(lines)
		if expected != received {
			t.Errorf("expected %q received %q", expected, received)
		}
	})
}
