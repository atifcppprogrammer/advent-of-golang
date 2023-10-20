package main

import (
	"adventofgolang/packages/utilities"
	"adventofgolang/solutions/year-2022/day-05/structures"
	"fmt"
	"testing"
)

func getListFrom(other []rune) *structures.List {
	list := new(structures.List)
	for _, element := range other {
		list.Push(element)
	}
	return list
}

func setSolutionPartEnv(t *testing.T, solutionPart string) {
	t.Setenv("PART", solutionPart)
	utilities.SetSolutionPart()
}

func assertListSlice(t *testing.T, exp, rec []structures.List) {
	if len(exp) != len(rec) {
		t.Errorf("expected %q received %q", len(exp), len(rec))
	}

	for index := 0; index < len(rec); index++ {
		expected := exp[index]
		received := rec[index]
		if equal, reason := expected.IsEqual(received); !equal {
			t.Errorf(fmt.Sprintf("lists @ %02d differ: %s", index, reason))
		}
	}
}

func TestSolution(t *testing.T) {
	t.Run("solves example input correctly", func(t *testing.T) {
		testCases := []struct{ solutionPart, expected string }{
			{"1", "CMZ"},
			{"2", "MCD"},
		}

		lines := []string{
			"    [D]",
			"[N] [C]",
			"[Z] [M] [P]",
			" 1   2   3",
			"",
			"move 1 from 2 to 1",
			"move 3 from 1 to 3",
			"move 2 from 2 to 1",
			"move 1 from 1 to 2",
		}

		for _, testCase := range testCases {
			setSolutionPartEnv(t, testCase.solutionPart)
			expected, received := testCase.expected, solution(lines)
			if expected != received {
				t.Errorf("expected %q received %q", expected, received)
			}
		}
	})
}

func TestExecInstr(t *testing.T) {

	t.Run("processes instruction correctly for part-1", func(t *testing.T) {
		receivedStacks, _ := getStacks([]string{
			"[D]",
			"[N] [C]",
			"[Z] [M] [P]",
			" 1   2   3",
			"",
		})

		expectedStacks, _ := getStacks([]string{
			"        [Z]",
			"        [N]",
			"    [C] [D]",
			"    [M] [P]",
			" 1   2   3",
			"",
		})

		setSolutionPartEnv(t, "1")
		execInstr(receivedStacks, "move 3 from 1 to 3")
		assertListSlice(t, expectedStacks, receivedStacks)
	})

	t.Run("processes instruction correctly for part-2", func(t *testing.T) {
		t.Setenv("PART", "1")
		utilities.SetSolutionPart()
		receivedStacks, _ := getStacks([]string{
			"[D]",
			"[N] [C]",
			"[Z] [M] [P]",
			" 1   2   3",
			"",
		})

		expectedStacks, _ := getStacks([]string{
			"        [D]",
			"        [N]",
			"    [C] [Z]",
			"    [M] [P]",
			" 1   2   3",
			"",
		})

		setSolutionPartEnv(t, "2")
		execInstr(receivedStacks, "move 3 from 1 to 3")
		assertListSlice(t, expectedStacks, receivedStacks)
	})
}

func TestGetStacks(t *testing.T) {
	t.Run("parses stacks correctly", func(t *testing.T) {
		lines := []string{
			"    [D]",
			"[N] [C]",
			"[Z] [M] [P]",
			" 1   2   3",
			"",
		}

		receivedStacks, instrIndex := getStacks(lines)
		expectedStacks := []structures.List{
			*getListFrom([]rune{'Z', 'N'}),
			*getListFrom([]rune{'M', 'C', 'D'}),
			*getListFrom([]rune{'P'}),
		}

		assertListSlice(t, expectedStacks, receivedStacks)
		expected, received := 5, instrIndex
		if expected != received {
			t.Errorf("expected %q received %q", expected, received)
		}
	})
}
