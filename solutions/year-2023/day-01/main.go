package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/atifcppprogrammer/advent-of-golang/internal/utilities"
)

const (
	regexStrOne = "[0-9]{1}"
	regexStrTwo = "([0-9]{1})|(one)|(two)|(three)|(four)|(five)|(six)|(seven)|(eight)|(nine)"
	intRegex    = "^[0-9]{1}$"
)

var digitWordMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

var correctionsMap = map[string]string{
	"oneight":   "oneeight",
	"threeight": "threeeight",
	"fiveight":  "fiveeight",
	"nineight":  "nineeight",
	"twone":     "twoone",
	"sevenine":  "sevennine",
	"eightwo":   "eighttwo",
}

func main() {
	filepath := utilities.GetInputFile()
	lines, _ := utilities.ReadLines(filepath)

	result := solution(lines)
	utilities.PrintSolution("2023", 01, "Trebuchet?!", strconv.Itoa(result))
}

func solution(lines []string) (result int) {
	makeCorrections(lines)
	for _, line := range lines {
		calValue, err := extractCalibrationVal(line)
		if err != nil {
			log.Fatal(err)
		}
		result += calValue
	}
	return
}

func makeCorrections(lines []string) {
	for index := range lines {
		for overlap, correction := range correctionsMap {
			lines[index] = strings.ReplaceAll(
				lines[index], overlap, correction,
			)
		}
	}
}

func extractCalibrationVal(line string) (int, error) {
	var regexStr = regexStrOne
	if utilities.SolutionPart == 2 {
		regexStr = regexStrTwo
	}

	var (
		regex   = regexp.MustCompile(regexStr)
		matches = regex.FindAllString(line, -1)
		first   = matches[0]
		last    = matches[len(matches)-1]
	)

	if utilities.SolutionPart == 2 {
		if !isDigitString(first) {
			first = digitWordMap[first]
		}
		if !isDigitString(last) {
			last = digitWordMap[last]
		}
	}

	joined := strings.Join([]string{first, last}, "")
	return strconv.Atoi(joined)
}

func isDigitString(numStr string) bool {
	regex := regexp.MustCompile(intRegex)
	return regex.MatchString(numStr)
}
