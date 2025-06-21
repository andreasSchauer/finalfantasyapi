package dataFormattingMixes

import (
	"slices"
	"strings"
)

type Combination struct {
	item1 string
	item2 string
}

func getCombinations(combinationStrings []string) []Combination {
	var combinations []Combination
	for _, comboString := range combinationStrings {
		if comboString == "" {
			continue
		}
		item1, item2 := splitCombinationString(comboString)
		combination := Combination{
			item1: item1,
			item2: item2,
		}
		combinations = append(combinations, combination)
	}

	return sortCombinations(combinations)
}

func splitCombinationString(combinationString string) (string, string) {
	items := strings.Split(combinationString, " + ")

	if slices.Index(itemNames, items[0]) <= slices.Index(itemNames, items[1]) {
		return items[0], items[1]
	}

	return items[1], items[0]
}

func sortCombinations(combinations []Combination) []Combination {
	slices.SortStableFunc(combinations, func(i, j Combination) int {
		if i.item1 == j.item1 {
			return compare(i.item2, j.item2)
		}

		return compare(i.item1, j.item1)
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
