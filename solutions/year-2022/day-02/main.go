package main

import (
	"adventofgolang/internal/utilities"
	"strconv"
)

const (
	roundWon      = 6
	roundLost     = 0
	roundDraw     = 3
	rockPoints    = 1
	paperPoints   = 2
	scissorPoints = 3
)

var partOneRoundOutcomeMap = map[string]int{
	"A X": roundDraw + rockPoints,
	"A Y": roundLost + paperPoints,
	"A Z": roundWon + scissorPoints,
	"B X": roundWon + rockPoints,
	"B Y": roundDraw + paperPoints,
	"B Z": roundLost + scissorPoints,
	"C X": roundLost + rockPoints,
	"C Y": roundWon + paperPoints,
	"C Z": roundDraw + scissorPoints,
}

var partTwoRoundOutcomeMap = map[string]int{
	"A X": roundLost + scissorPoints,
	"A Y": roundDraw + rockPoints,
	"A Z": roundWon + paperPoints,
	"B X": roundLost + rockPoints,
	"B Y": roundDraw + paperPoints,
	"B Z": roundWon + scissorPoints,
	"C X": roundLost + paperPoints,
	"C Y": roundDraw + scissorPoints,
	"C Z": roundWon + rockPoints,
}

func main() {
	filepath := utilities.GetInputFile()
	lines, _ := utilities.ReadLines(filepath)

	outcomeSum := solution(lines)
	utilities.PrintSolution("2022", 2, "Rock Paper Scissors", strconv.Itoa(outcomeSum))
}

func solution(lines []string) (outcomeSum int) {
	var solutionPartOutcomeMap map[string]int
	if utilities.SolutionPart == 1 {
		solutionPartOutcomeMap = partOneRoundOutcomeMap
	} else if utilities.SolutionPart == 2 {
		solutionPartOutcomeMap = partTwoRoundOutcomeMap
	}

	for _, line := range lines {
		outcomeSum += solutionPartOutcomeMap[line]
	}
	return
}
