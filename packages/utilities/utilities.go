package utilities

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var SolutionPart int

func init() {
	value, exists := os.LookupEnv("PART")
	if exists {
		if valueInt, err := strconv.Atoi(value); err == nil {
			SolutionPart = valueInt
		} else {
			panic(err)
		}
	}
}

func PrintSolution(event string, day int, name, solution string) {
	fmt.Printf(
		"Advent of Code (%s) | Day %02d - %s | Solution: %s \n",
		event, day, name, solution,
	)
}

func GetInputFile() string {
	if len(os.Args) < 2 {
		log.Fatal("Please provided path to puzzle's input file as argument")
	}

	exists, err := FileExists(os.Args[1])
	if !exists {
		log.Fatal(err)
	}

	return os.Args[1]
}

func FileExists(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	return err == nil, err
}

func ReadLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var (
		scanner = bufio.NewScanner(file)
		lines   = make([]string, 0)
	)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}
