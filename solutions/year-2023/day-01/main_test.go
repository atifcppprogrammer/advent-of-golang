package main

import (
	"strings"
	"testing"

	"github.com/atifcppprogrammer/advent-of-golang/internal/utilities"
)

func setSolutionPartEnv(t *testing.T, solutionPart string) {
	t.Setenv("PART", solutionPart)
	utilities.SetSolutionPart()
}

func TestSolution(t *testing.T) {
	t.Run("solves example input correctly for each part", func(t *testing.T) {
		testCases := []struct {
			solutionPart string
			expected     int
			lines        []string
		}{
			{
				solutionPart: "1",
				expected:     142,
				lines: []string{
					"1abc2",
					"pqr3stu8vwx",
					"a1b2c3d4e5f",
					"treb7uchet",
				},
			},
			{
				solutionPart: "2",
				expected:     281,
				lines: []string{
					"two1nine",
					"eightwothree",
					"abcone2threexyz",
					"xtwone3four",
					"4nineeightseven2",
					"zoneight234",
					"7pqrstsixteen",
				},
			},
		}

		for _, testCase := range testCases {
			setSolutionPartEnv(t, testCase.solutionPart)
			expected, received := testCase.expected, solution(testCase.lines)
			if expected != received {
				t.Errorf("expected %q received %q", expected, received)
			}
		}
	})
}

func TestExtractCalibrationVal(t *testing.T) {
	t.Run("extracts value correctly for each part", func(t *testing.T) {
		testCases := []struct {
			line     string
			expected []int
		}{
			{line: "seven898989seven", expected: []int{89, 77}},
			{line: "1one2six4689five", expected: []int{19, 15}},
			{line: "723three461eight", expected: []int{71, 78}},
			{line: "eight4eighteight", expected: []int{44, 88}},
		}

		for _, testCase := range testCases {
			for index, solutionPart := range []string{"1", "2"} {
				setSolutionPartEnv(t, solutionPart)
				received, _ := extractCalibrationVal(testCase.line)
				expected := testCase.expected[index]
				if expected != received {
					t.Errorf("expected %q received %q", expected, received)
				}
			}
		}
	})
}

func TestIsDigitString(t *testing.T) {
	t.Run("returns false if string doesn't represents digit", func(t *testing.T) {
		testCases := []string{"word", "digit", "10", "11", "323"}
		for _, testCase := range testCases {
			expected, received := false, isDigitString(testCase)
			if expected != received {
				t.Errorf("expected %t received %t", expected, received)
			}
		}
	})

	t.Run("returns true if string represents digit", func(t *testing.T) {
		testCases := strings.Split("0123456789", "")
		for _, testCase := range testCases {
			expected, received := true, isDigitString(testCase)
			if expected != received {
				t.Errorf("expected %t received %t", expected, received)
			}
		}
	})
}
