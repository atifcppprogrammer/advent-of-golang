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
		testCases := [][]struct {
			solutionPart string
			expected     int
		}{
			{{"1", 7}, {"2", 19}},
			{{"1", 5}, {"2", 23}},
			{{"1", 6}, {"2", 23}},
			{{"1", 10}, {"2", 29}},
			{{"1", 11}, {"2", 26}},
		}

		lines := []string{
			"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			"bvwbjplbgvbhsrlpgdmjqwftvncz",
			"nppdvjthqldpwncqszvftbrmjlhg",
			"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
		}

		for i, testCase := range testCases {
			for j := range []int{1, 2} {
				setSolutionPartEnv(t, testCase[j].solutionPart)
				expected, received := testCase[j].expected, solution([]string{lines[i]})
				if expected != received {
					t.Errorf("expected %q received %q", expected, received)
				}
			}
		}
	})
}

func TestGetMarkerLengthForSolutionPart(t *testing.T) {
	t.Run("returns 4 for part-1", func(t *testing.T) {
		setSolutionPartEnv(t, "1")
		expected, received := 4, getMarkerLengthForSolutionPart()
		if expected != received {
			t.Errorf("expected %q received %q", expected, received)
		}
	})

	t.Run("returns 14 for part-2", func(t *testing.T) {
		setSolutionPartEnv(t, "2")
		expected, received := 14, getMarkerLengthForSolutionPart()
		if expected != received {
			t.Errorf("expected %q received %q", expected, received)
		}
	})
}
