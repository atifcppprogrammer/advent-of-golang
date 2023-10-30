package main

import (
	"adventofgolang/packages/utilities"
	"strconv"
)

func main() {
	filepath := utilities.GetInputFile()
	lines, _ := utilities.ReadLines(filepath)

	calorieCount := solution(lines)
	utilities.PrintSolution("2022", 1, "Calorie Counting", strconv.Itoa(calorieCount))
}

func solution(lines []string) int {
	var (
		maxCalories = []int{0, 0, 0}
		currentSum  int
	)

	for index, line := range lines {
		isLast := index == len(lines)-1
		if line == "" {
			setMaxCalories(maxCalories, currentSum)
			currentSum = 0
		} else {
			calories, _ := strconv.Atoi(line)
			currentSum += calories
			if isLast {
				setMaxCalories(maxCalories, currentSum)
			}
		}
	}

	var result int
	if utilities.SolutionPart == 1 {
		result = maxCalories[0]
	} else if utilities.SolutionPart == 2 {
		result = maxCalories[0] + maxCalories[1] + maxCalories[2]
	}

	return result
}

func setMaxCalories(maxCalories []int, currentSum int) {
	var (
		a = maxCalories[0]
		b = maxCalories[1]
		c = maxCalories[2]
	)

	if a <= currentSum {
		maxCalories[2] = b
		maxCalories[1] = a
		maxCalories[0] = currentSum
	} else if b <= currentSum && currentSum < a {
		maxCalories[2] = maxCalories[1]
		maxCalories[1] = currentSum
	} else if c <= currentSum && currentSum < b {
		maxCalories[2] = currentSum
	}
}
