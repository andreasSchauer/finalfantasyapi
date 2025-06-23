package dataFormattingMixes

import (
	"slices"
	"strings"
)


type Combination struct {
	Item1 string `json:"first_item"`
	Item2 string `json:"second_item"`
}


func getCombinations(combinationStrings []string) []Combination {
	var combinations []Combination
	for _, comboString := range combinationStrings {
		if comboString == "" {
			continue
		}
		combination := createCombination(comboString)
		combinations = append(combinations, combination)
	}

	return sortCombinations(combinations)
}


func createCombination(combinationString string) Combination {
	items := strings.Split(combinationString, " + ")
	combination := Combination{
		Item1: items[0],
		Item2: items[1],
	}

	if slices.Index(itemNames, items[0]) > slices.Index(itemNames, items[1]) {
		combination.Item1 = items[1]
		combination.Item2 = items[0]
	}

	return combination
}


func sortCombinations(combinations []Combination) []Combination {
	slices.SortStableFunc(combinations, func(i, j Combination) int {
		if i.Item1 == j.Item1 {
			return compare(i.Item2, j.Item2)
		}

		return compare(i.Item1, j.Item1)
	})

	return combinations
}


func compare(a, b string) int {
	if slices.Index(itemNames, a) < slices.Index(itemNames, b) {
		return -1
	} else if slices.Index(itemNames, a) > slices.Index(itemNames, b) {
		return 1
	} else {
		return 0
	}
}
