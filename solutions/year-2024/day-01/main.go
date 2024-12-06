package main

import (
	"github.com/atifcppprogrammer/advent-of-golang/internal/utilities"
	"sort"
	"strconv"
	"strings"
)

func main() {
	filepath := utilities.GetInputFile()
	lines, _ := utilities.ReadLines(filepath)

	result := solution(lines)
	utilities.PrintSolution("2024", 01, "", strconv.Itoa(result))
}

func solution(lines []string) (result int) {
	var (
		count = len(lines)
		l1    = make(sort.IntSlice, count)
		l2    = make(sort.IntSlice, count)
	)

	for i, line := range lines {
		val1, val2, _ := strings.Cut(line, "   ")
		num1, _ := strconv.Atoi(val1)
		num2, _ := strconv.Atoi(val2)
		l1[i] = num1
		l2[i] = num2
	}

	if utilities.SolutionPart == 1 {
		sort.Sort(l1)
		sort.Sort(l2)

		for i := 0; i < count; i++ {
			diff := l1[i] - l2[i]
			if diff < 0 {
				result += -diff
			} else {
				result += +diff
			}
		}

		return
	}

	if utilities.SolutionPart == 2 {
		numCounts := make(map[int]int)

		for _, num := range l2 {
			if _, exists := numCounts[num]; !exists {
				numCounts[num] = 1
			} else {
				numCounts[num]++
			}
		}

		for _, num := range l1 {
			if count, exists := numCounts[num]; exists {
				result += num * count
			}
		}
	}

	return
}
