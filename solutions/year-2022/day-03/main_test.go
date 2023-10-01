package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("solves example input correctly", func(t *testing.T) {
		rucksacks := []string{
			"vJrwpWtwJgWrhcsFMMfFFhFp",
			"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
			"PmmdzqPrVvPwwTWBwg",
			"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
			"ttgJtRGJQctTZtZT",
			"CrZsJsPPZsGzwwsLwLmpwMD",
		}
		received, expected := solution(rucksacks), 157
		if expected != received {
			t.Errorf("Expected prioritySum %q but received %q", expected, received)
		}
	})
}

func TestFindSharedItemPriority(t *testing.T) {
	t.Run("computes correct priority for shared item", func(t *testing.T) {
		rucksack := "gggzzzpprrrlll"
		received := findSharedItemPriority(rucksack)
		expected := PriorityMap['p']
		if received != expected {
			t.Errorf("Expected priority %q but received %q", expected, received)
		}
	})
}
