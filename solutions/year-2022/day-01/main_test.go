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

		testCases := []struct {
			solutionPart string
			expected     int
		}{
			{"1", 24000},
			{"2", 45000},
		}

		for _, testCase := range testCases {
			setSolutionPartEnv(t, testCase.solutionPart)
			expected, received := testCase.expected, solution(lines)
			if expected != received {
				t.Errorf("Expected maxCalorieCount %q received %q", expected, received)
			}
		}
	})
}

func TestSetMaxCalories(t *testing.T) {
	t.Run("returns a sorted list when a < sum", func(t *testing.T) {
		var (
			currentSum       = 5
			maxCalories      = []int{4, 3, 2}
			expectedCalories = []int{5, 4, 3}
		)
		setMaxCalories(maxCalories, currentSum)
		for index, calorie := range expectedCalories {
			expected, received := calorie, maxCalories[index]
			if expected != received {
				t.Errorf("Expected calorie %q at index %q, received %q", expected, index, received)
			}
		}
	})

	t.Run("returns a sorted list when b < sum < a", func(t *testing.T) {
		var (
			currentSum       = 4
			maxCalories      = []int{4, 3, 2}
			expectedCalories = []int{4, 4, 3}
		)
		setMaxCalories(maxCalories, currentSum)
		for index, calorie := range expectedCalories {
			expected, received := calorie, maxCalories[index]
			if expected != received {
				t.Errorf("Expected calorie %q at index %q, received %q", expected, index, received)
			}
		}
	})

	t.Run("returns a sorted list when c < sum < b", func(t *testing.T) {
		var (
			currentSum       = 3
			maxCalories      = []int{4, 3, 2}
			expectedCalories = []int{4, 3, 3}
		)
		setMaxCalories(maxCalories, currentSum)
		for index, calorie := range expectedCalories {
			expected, received := calorie, maxCalories[index]
			if expected != received {
				t.Errorf("Expected calorie %q at index %q, received %q", expected, index, received)
			}
		}
	})
}
