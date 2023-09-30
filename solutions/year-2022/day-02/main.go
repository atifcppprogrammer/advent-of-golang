package main

import (
	"adventofgolang/packages/utilities"
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

var roundOutcomeMap = map[string]int{
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

func main() {
	filepath := utilities.GetInputFile()
	lines, _ := utilities.ReadLines(filepath)

	outcomeSum := solution(lines)
	utilities.PrintSolution("2022", 2, "Rock Paper Scissors", strconv.Itoa(outcomeSum))
}

func solution(lines []string) (outcomeSum int) {
	for _, line := range lines {
		outcomeSum += roundOutcomeMap[line]
	}
	return
}
