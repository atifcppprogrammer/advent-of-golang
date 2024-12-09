package main

import (
	"fmt"
	"github.com/atifcppprogrammer/advent-of-golang/internal/utilities"
	"sort"
	"strconv"
	"strings"
)

func main() {
	filepath := utilities.GetInputFile()
	lines, _ := utilities.ReadLines(filepath)

	result := solution(lines)
	utilities.PrintSolution("2024", 02, "Red-Nosed Reports", strconv.Itoa(result))
}

func solution(lines []string) (result int) {
outerloop:
	for _, line := range lines {
		var (
			trimmed    = strings.Split(line, " ")
			report     = make(sort.IntSlice, len(trimmed))
			diffs      = make([]int, len(trimmed))
			diffCounts = map[int]int{0: 0, 1: 1, -1: 0}
		)

		for i, val := range trimmed {
			report[i], _ = strconv.Atoi(val)
			if i >= 1 {
				diffs[i] = report[i] - report[i-1]
			}
		}

		if utilities.SolutionPart == 1 {
			if !isSorted(report) {
				continue outerloop
			}

			for _, diff := range diffs[1:] {
				if abs(diff) < 1 || abs(diff) > 3 {
					continue outerloop
				}
			}
		}

		if utilities.SolutionPart == 2 {
			for _, diff := range diffs[1:] {
				if diff == 0 {
					diffCounts[0]++
				} else if diff < 0 {
					diffCounts[-1]++
				} else {
					diffCounts[1]++
				}
			}

			fmt.Println(diffCounts, line)
		}

		result++
	}

	return
}

func isSorted(d sort.Interface) bool {
	if r := sort.Reverse(d); sort.IsSorted(d) || sort.IsSorted(r) {
		return true
	}
	return false
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}
