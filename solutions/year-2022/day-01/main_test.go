package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("solves example input correctly", func(t *testing.T) {
		lines := []string{
			"1000",
			"2000",
			"3000",
			"",
			"4000",
			"",
			"5000",
			"6000",
			"",
			"7000",
			"8000",
			"9000",
			"",
			"10000",
		}

		expected, received := 24000, solution(lines)
		if expected != received {
			t.Errorf("Expected maxCalorieCount %q received %q", expected, received)
		}
	})
}

func TestMax(t *testing.T) {
	t.Run("returns the greater value of the provided arguments", func(t *testing.T) {
		expected, received := 2, max(1, 2)
		if expected != received {
			t.Errorf("Expected maxCalorieCount %q received %q", expected, received)
		}
	})
}
