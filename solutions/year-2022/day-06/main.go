package main

import (
	"adventofgolang/internal/utilities"
	"log"
	"strconv"
)

func main() {
	filepath := utilities.GetInputFile()
	lines, _ := utilities.ReadLines(filepath)

	result := solution(lines)
	utilities.PrintSolution("2022", 06, "Tuning Trouble", strconv.Itoa(result))
}

func solution(lines []string) (result int) {
	var (
		markerMap    = make(map[byte]int)
		dataStream   = lines[0]
		markerLength = getMarkerLengthForSolutionPart()
	)

	start, end := 0, markerLength
	for end < len(dataStream) {
		for i := start; i < end; i++ {
			markerMap[dataStream[i]]++
		}

		if len(markerMap) == markerLength {
			return end
		}

		for i := start; i < end; i++ {
			delete(markerMap, dataStream[i])
		}

		start++
		end++
	}

	log.Fatalln("No marker found in data stream")
	return 0
}

func getMarkerLengthForSolutionPart() int {
	switch utilities.SolutionPart {
	case 1:
		return 4
	case 2:
		return 14
	default:
		panic("Solution part must be either 1 or 2")
	}
}
