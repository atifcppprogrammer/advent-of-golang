package main

import (
	"adventofgolang/internal/utilities"
	"testing"
)

func setSolutionPartEnv(t *testing.T, solutionPart string) {
	t.Setenv("PART", solutionPart)
	utilities.SetSolutionPart()
}

func TestSolution(t *testing.T) {
	t.Run("solves example input correctly for part-1", func(t *testing.T) {
		setSolutionPartEnv(t, "1")
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

	t.Run("solves example input correctly for part-2", func(t *testing.T) {
		setSolutionPartEnv(t, "2")
		lines := []string{
			"A Y",
			"B X",
			"C Z",
		}

		expected, received := 12, solution(lines)
		if expected != received {
			t.Errorf("expected outcomeSum %q received %q", expected, received)
		}
	})
}
