package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/atifcppprogrammer/advent-of-golang/internal/utilities"
)

func main() {
	filepath := utilities.GetInputFile()
	lines, _ := utilities.ReadLines(filepath)

	result := solution(lines)
	utilities.PrintSolution("2022", 5, "Supply Stacks", result)
}

func solution(lines []string) string {
	stacks, instrIndex := getStacks(lines)
	for i := instrIndex; i < len(lines); i++ {
		execInstr(stacks, lines[i])
	}

	var result string
	for _, stack := range stacks {
		result += fmt.Sprintf("%c", stack.Last())
	}

	return result
}

func execInstr(stacks []List, instr string) {
	pattern := regexp.MustCompile("[0-9]+")
	matches := pattern.FindAllString(instr, -1)

	var (
		diff, _ = strconv.Atoi(string(matches[0]))
		from, _ = strconv.Atoi(string(matches[1]))
		to, _   = strconv.Atoi(string(matches[2]))
	)

	var (
		destination = &stacks[to-1]
		source      = &stacks[from-1]
	)

	var execPartOne = func() {
		for ; diff > 0; diff-- {
			destination.Push(source.Pop())
		}
	}

	var execPartTwo = func() {
		if diff == 1 {
			destination.Push(source.Pop())
			return
		}

		revStack := List{}
		for ; diff > 0; diff-- {
			revStack.Push(source.Pop())
		}

		for !revStack.IsEmpty() {
			destination.Push(revStack.Pop())
		}
	}

	switch utilities.SolutionPart {
	case 1:
		execPartOne()
	case 2:
		execPartTwo()
	}
}

func getStacks(lines []string) ([]List, int) {
	var descriptionIndex, instructionIndex int
	for index, line := range lines {
		if line == "" {
			descriptionIndex = index - 1
			instructionIndex = index + 1
			break
		}
	}

	var populateStack = func(j int) List {
		stack := List{}
		for i := descriptionIndex - 1; true; i-- {
			if i == -1 || j >= len(lines[i]) || lines[i][j] == 32 {
				break
			} else {
				crate := rune(lines[i][j])
				stack.Push(crate)
			}
		}

		return stack
	}

	stacks := make([]List, 0)
	for index, char := range lines[descriptionIndex] {
		if index%2 != 0 && char != ' ' {
			stacks = append(stacks, populateStack(index))
		}
	}

	return stacks, instructionIndex
}
