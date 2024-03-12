package main

import (
	"log"
	"math"
	"strconv"

	"github.com/atifcppprogrammer/advent-of-golang/internal/utilities"
)

func main() {
	filepath := utilities.GetInputFile()
	lines, _ := utilities.ReadLines(filepath)

	result := solution(lines)
	utilities.PrintSolution("2021", 01, "Sonar Sweep", strconv.Itoa(result))
}

func solution(lines []string) (result int) {
	var finalIndex, windowSize int
	switch utilities.SolutionPart {
	case 1:
		finalIndex = len(lines) - 1
		windowSize = 1
	case 2:
		finalIndex = len(lines) - 3
		windowSize = 3
	default:
		log.Fatal(`"PART" environment variable must be either 1 or 2`)
	}

	var (
		currMeasure  int
		prevMeasure  int   = math.MaxInt
		measurements []int = parseMeasurements(lines)
	)

	for i := 0; i <= finalIndex; i++ {
		currMeasure = findWindowSum(measurements, windowSize, i)
		if currMeasure > prevMeasure {
			result++
		}
		prevMeasure = currMeasure
	}
	return
}

func parseMeasurements(lines []string) (measurements []int) {
	measurements = make([]int, len(lines))
	for i, line := range lines {
		measurement, _ := strconv.Atoi(line)
		measurements[i] = measurement
	}
	return
}

func findWindowSum(measurements []int, size, start int) int {
	if size == 1 {
		return measurements[start]
	} else {
		return measurements[start] + measurements[start+1] + measurements[start+2]
	}
}
