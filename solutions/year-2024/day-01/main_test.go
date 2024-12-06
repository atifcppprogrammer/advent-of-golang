package main

import (
	"fmt"
	"testing"

	"github.com/atifcppprogrammer/advent-of-golang/internal/utilities"
)

func setSolutionPartEnv(t *testing.T, solutionPart string) {
	t.Setenv("PART", solutionPart)
	utilities.SetSolutionPart()
}

func TestSolution(t *testing.T) {
	t.Run("solves example input correctly", func(t *testing.T) {
		lines := []string{
			"3   4",
			"4   3",
			"2   5",
			"1   3",
			"3   9",
			"3   3",
		}

		tests := []struct {
			solutionPart string
			want         int
		}{
			{solutionPart: "1", want: 11},
			{solutionPart: "2", want: 31},
		}
		for _, tt := range tests {
			t.Run(fmt.Sprintf("solves example correct for part %s", tt.solutionPart), func(t *testing.T) {
				setSolutionPartEnv(t, tt.solutionPart)
				got, want := solution(lines), tt.want
				if got != want {
					t.Errorf("want %d got %d", want, got)
				}
			})
		}
	})
}
