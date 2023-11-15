package main

import (
	"testing"

	"github.com/atifcppprogrammer/advent-of-golang/internal/utilities"
)

func setSolutionPartEnv(t *testing.T, solutionPart string) {
	t.Setenv("PART", solutionPart)
	utilities.SetSolutionPart()
}

func TestSolution(t *testing.T) {
	t.Run("solves example input correctly", func(t *testing.T) {
		testCases := []struct {
			solutionPart string
			expected     int
		}{{"1", 150}, {"2", 900}}

		lines := []string{
			"forward 5",
			"down 5",
			"forward 8",
			"up 3",
			"down 8",
			"forward 2",
		}

		for _, testCase := range testCases {
			setSolutionPartEnv(t, testCase.solutionPart)
			expected, received := testCase.expected, solution(lines)
			if expected != received {
				t.Errorf("expected %q received %q", expected, received)
			}
		}
	})
}
