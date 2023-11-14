package main

import (
	"adventofgolang/internal/utilities"
	"log"
	"strconv"
	"strings"
	"sync"
)

func main() {
	filepath := utilities.GetInputFile()
	lines, _ := utilities.ReadLines(filepath)
	pairCount := solution(lines)
	utilities.PrintSolution("2022", 4, "Camp Cleanup", strconv.Itoa(pairCount))
}

func solution(lines []string) int {
	pairStatus := make([]bool, len(lines))

	var wg sync.WaitGroup
	for index, line := range lines {
		wg.Add(1)
		go func(i int, l string) {
			pairStatus[i] = isPairSatisfactory(l)
			wg.Done()
		}(index, line)
	}
	wg.Wait()

	pairCount := 0
	for _, isSatisFactory := range pairStatus {
		if isSatisFactory {
			pairCount++
		}
	}

	return pairCount
}

func isPairSatisfactory(line string) bool {
	var getPairFrom = func(segment string) Pair {
		before, after, found := strings.Cut(segment, "-")
		if !found {
			log.Fatalln("Failed to parse line", line)
		}

		start, _ := strconv.Atoi(before)
		end, _ := strconv.Atoi(after)
		return Pair{start, end}
	}

	before, after, found := strings.Cut(line, ",")
	if !found {
		log.Fatalln("Failed to parse line", line)
	}

	pOne := getPairFrom(before)
	pTwo := getPairFrom(after)
	return pOne.Contains(pTwo) || pTwo.Contains(pOne)
}
