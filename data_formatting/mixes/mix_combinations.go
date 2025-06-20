package dataFormatting

import(
	"fmt"
	"slices"
	"strings"
)



func getMixTable(mixName string, combinations []Combination) string {
	combinationsRejoined := rejoinCombinations(combinations)
	sortedString := strings.Join(combinationsRejoined, "\n")
	mixTable := fmt.Sprintf("%s\n\n%s", mixName, sortedString)

	return mixTable
}



func rejoinCombinations(combinationsSlice []Combination) []string {
	var rejoinedCombos []string

	for _, combination := range combinationsSlice {
		joinedString := fmt.Sprintf("%s + %s", combination.item1, combination.item2)
		rejoinedCombos = append(rejoinedCombos, joinedString)
	}

	return rejoinedCombos
}



func getCombinationSlice(combinations []string) []Combination {
	var combinationSlice []Combination
	for _, comboString := range combinations {
		if comboString == "" {
			continue
		}
		item1, item2 := splitCombination(comboString)
		combination := Combination{
			item1: item1,
			item2: item2,
		}
		combinationSlice = append(combinationSlice, combination)
	}

	return sortCombinationSlice(combinationSlice)
}


func splitCombination(combination string) (string, string) {
	items := strings.Split(combination, " + ")

	if slices.Index(itemNames, items[0]) <= slices.Index(itemNames, items[1]) {
		return items[0], items[1]
	}

	return items[1], items[0]
}



func sortCombinationSlice(combinationSlice []Combination) []Combination {
	slices.SortStableFunc(combinationSlice, func(i, j Combination) int {
		if i.item1 == j.item1 {
			return compare(i.item2, j.item2)
		}

		return compare(i.item1, j.item1)
	})

	return combinationSlice
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