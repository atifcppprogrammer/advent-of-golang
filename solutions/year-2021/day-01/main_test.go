package main

import (
	"fmt"
	"testing"

	"github.com/atifcppprogrammer/advent-of-golang/internal/utilities"
)

func setSolutionPartEnv(t *testing.T, solutionPart string) {
	t.Helper()
	t.Setenv("PART", solutionPart)
	utilities.SetSolutionPart()
}

func TestSolution(t *testing.T) {
	t.Run("solves example input correctly", func(t *testing.T) {
		lines := []string{"199", "200", "208", "210", "200", "207", "240", "269", "260", "263"}
		tests := []struct {
			name         string
			solutionPart int
			lines        []string
			want         int
		}{
			{name: "solution-part-01", solutionPart: 1, lines: lines, want: 7},
			{name: "solution-part-02", solutionPart: 2, lines: lines, want: 5},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				setSolutionPartEnv(t, fmt.Sprintf("%d", tt.solutionPart))
				got, want := solution(tt.lines), tt.want
				if got != want {
					t.Errorf("got %q want %q", got, want)
				}
			})
		}
	})
}
