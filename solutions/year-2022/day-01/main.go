package main

import (
	"adventofgolang/packages/utilities"
	"strconv"
)

func main() {
	filepath := utilities.GetInputFile()
	lines, _ := utilities.ReadLines(filepath)

	calorieCount := solution(lines)
	utilities.PrintSolution(
		"2022", 1, "Calorie Counting",
		strconv.Itoa(calorieCount),
	)
}

func solution(lines []string) int {
	var currentSum, maxCalories int

	for index, line := range lines {
		isLast := index == len(lines)-1
		if line == "" {
			maxCalories = max(maxCalories, currentSum)
			currentSum = 0
		} else {
			calories, _ := strconv.Atoi(line)
			currentSum += calories
			if isLast {
				maxCalories = max(maxCalories, currentSum)
			}
		}
	}

	return maxCalories
}

func max(a, b int) int {
	if a <= b {
		return b
	} else {
		return a
	}
}
