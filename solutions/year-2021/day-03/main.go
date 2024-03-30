package main

import (
	"log"
	"strconv"

	"github.com/atifcppprogrammer/advent-of-golang/internal/utilities"
)

func main() {
	filepath := utilities.GetInputFile()
	lines, _ := utilities.ReadLines(filepath)

	result := solution(lines)
	utilities.PrintSolution("2021", 03, "Binary Diagnostic", strconv.Itoa(result))
}

func solution(lines []string) (result int) {
	switch utilities.SolutionPart {
	case 1:
		result = solutionPartOne(lines)
	case 2:
		result = solutionPartTwo(lines)
	default:
		log.Fatalf("%q environment variable must be either 1 or 2 \n", "PART")
	}
	return
}

func solutionPartOne(lines []string) int {
	var (
		lineCount      = len(lines)
		counts0        = make([]int, len(lines[0]))
		gammaRateStr   string
		epsilonRateStr string
	)

	for _, line := range lines {
		for i, character := range line {
			if character == '0' {
				counts0[i]++
			}
		}
	}

	for _, count0 := range counts0 {
		count1 := lineCount - count0
		if count0 <= count1 {
			gammaRateStr += "1"
			epsilonRateStr += "0"
		} else {
			gammaRateStr += "0"
			epsilonRateStr += "1"
		}
	}

	gammaRate, _ := strconv.ParseInt(gammaRateStr, 2, 0)
	epsilonRate, _ := strconv.ParseInt(epsilonRateStr, 2, 0)
	return int(gammaRate * epsilonRate)
}

func solutionPartTwo(lines []string) int {
	var (
		lineCount           = len(lines)
		oxygenGenRateStr    string
		co2GenRateStr       string
		oxygenGenCandidates = make(map[int]struct{}, lineCount)
		co2GenCandidates    = make(map[int]struct{}, lineCount)
	)

	for i := range lines {
		oxygenGenCandidates[i] = struct{}{}
		co2GenCandidates[i] = struct{}{}
	}

	oxygenGenRateStr = reduceCandidates(
		lines,
		oxygenGenCandidates,
		func(a, b int) byte {
			if a <= b {
				return '1'
			} else {
				return '0'
			}
		},
	)

	co2GenRateStr = reduceCandidates(
		lines,
		co2GenCandidates,
		func(a, b int) byte {
			if a <= b {
				return '0'
			} else {
				return '1'
			}
		},
	)

	oxygenGenRate, _ := strconv.ParseInt(oxygenGenRateStr, 2, 0)
	co2GenRate, _ := strconv.ParseInt(co2GenRateStr, 2, 0)
	return int(oxygenGenRate * co2GenRate)
}

func reduceCandidates(lines []string, candidateSet map[int]struct{}, findCommon func(a, b int) byte) string {
	for i := range lines {
		candidateSet[i] = struct{}{}
	}

	for i := 0; len(candidateSet) != 1; i++ {
		counts0 := 0
		for candidateIndex := range candidateSet {
			candidate := lines[candidateIndex]
			if candidate[i] == '0' {
				counts0++
			}
		}

		commonValue := findCommon(counts0, len(candidateSet)-counts0)
		for c := range candidateSet {
			candidate := lines[c]
			if candidate[i] != byte(commonValue) {
				delete(candidateSet, c)
			}
		}
	}

	return lines[getFirstKey(candidateSet)]
}

func getFirstKey(m map[int]struct{}) int {
	for key := range m {
		return key
	}
	return 0
}
