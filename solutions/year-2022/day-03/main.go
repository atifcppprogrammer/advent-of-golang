package main

import (
	"adventofgolang/packages/utilities"
	"log"
	"strconv"
	"sync"
)

var PriorityMap = make(map[rune]int)

func main() {
	filepath := utilities.GetInputFile()
	rucksacks, _ := utilities.ReadLines(filepath)

	prioritySum := solution(rucksacks)
	prioritySumStr := strconv.Itoa(prioritySum)
	utilities.PrintSolution("2022", 3, "Rucksack Reorganization", prioritySumStr)
}

func solution(rucksacks []string) int {
	prioritySum := 0
	priorities := make([]int, len(rucksacks))

	var wg sync.WaitGroup
	for index, rucksack := range rucksacks {
		wg.Add(1)
		go func(i int, r string) {
			priorities[i] = findSharedItemPriority(r)
			wg.Done()
		}(index, rucksack)
	}
	wg.Wait()

	for _, priority := range priorities {
		prioritySum += priority
	}

	return prioritySum
}

func findSharedItemPriority(rucksack string) int {
	mid := len(rucksack) / 2
	compartmentOne := rucksack[0:mid]
	mapOne := make(map[rune]bool)
	for _, letter := range compartmentOne {
		mapOne[letter] = true
	}

	compartmentTwo := rucksack[mid:]
	mapTwo := make(map[rune]bool)
	for _, letter := range compartmentTwo {
		mapTwo[letter] = true
	}

	for _, letter := range rucksack {
		if mapOne[letter] && mapTwo[letter] {
			return PriorityMap[letter]
		}
	}

	log.Fatalln("unexpected condition: no shared item found")
	return 0
}

func init() {
	priority := 1
	lowerCaseLetters := "abcdefghijklmnopqrstuvwxyz"
	for _, letter := range lowerCaseLetters {
		PriorityMap[letter] = priority
		priority++
	}

	upperCaseLetters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for _, letter := range upperCaseLetters {
		PriorityMap[letter] = priority
		priority++
	}
}
