package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/atifcppprogrammer/advent-of-golang/internal/utilities"
)

type Position struct {
	X, Y, Aim int
}

func main() {
	filepath := utilities.GetInputFile()
	lines, _ := utilities.ReadLines(filepath)

	result := solution(lines)
	utilities.PrintSolution("2021", 02, "Dive", strconv.Itoa(result))
}

func solution(lines []string) int {
	position := Position{}
	for _, line := range lines {
		updatePos(&position, strings.Split(line, " "))
	}
	return position.X * position.Y
}

func updatePos(pos *Position, line []string) {
	moveType := line[0]
	diff, _ := strconv.Atoi(line[1])

	if utilities.SolutionPart == 1 {
		switch moveType {
		case "forward":
			pos.X += diff
		case "up":
			pos.Y -= diff
		case "down":
			pos.Y += diff
		default:
			log.Fatal("unrecognized movement type")
		}
		return
	}

	if utilities.SolutionPart == 2 {
		switch moveType {
		case "forward":
			pos.X += diff
			pos.Y += pos.Aim * diff
		case "up":
			pos.Aim -= diff
		case "down":
			pos.Aim += diff
		default:
			log.Fatal("unrecognized movement type")
		}
		return
	}

	log.Fatal("unrecognized solution part")
}
