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
		lines := []string{
			"00100",
			"11110",
			"10110",
			"10111",
			"10101",
			"01111",
			"00111",
			"11100",
			"10000",
			"11001",
			"00010",
			"01010",
		}

		tests := []struct {
			name         string
			solutionPart string
			want         int
		}{
			{name: "part-1", solutionPart: "1", want: 198},
			{name: "part-2", solutionPart: "2", want: 230},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				setSolutionPartEnv(t, tt.solutionPart)
				got := solution(lines)
				if got != tt.want {
					t.Errorf("got %d, want %d", got, tt.want)
				}
			})
		}
	})
}
