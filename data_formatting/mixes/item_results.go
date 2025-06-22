package dataFormattingMixes

import(
	"fmt"
	"slices"
	"strings"
)


type ItemResult struct {
	Mixname		string
	FirstItem	string
	SecondItem	string
}


func appendItemResults(itemResultsPtr *[]ItemResult, mixName string, combinations []Combination) {
	for _, combination := range combinations {
		itemResult := ItemResult{
			Mixname: mixName,
			FirstItem: combination.item1,
			SecondItem: combination.item2,
		}

		*itemResultsPtr = append(*itemResultsPtr, itemResult)
	}
}


func sortItemResults(itemResultsPtr *[]ItemResult) {
	slices.SortStableFunc(*itemResultsPtr, func(i, j ItemResult) int {
		if i.FirstItem == j.FirstItem {
			return compare(i.SecondItem, j.SecondItem)
		}

		return compare(i.FirstItem, j.FirstItem)
	})
}


func formatItemResults(itemResultsPtr *[]ItemResult) string {
	var stringSlice []string

	for _, itemResult := range *itemResultsPtr {
		resultString := fmt.Sprintf("%s + %s = %s", itemResult.FirstItem, itemResult.SecondItem, itemResult.Mixname)

		stringSlice = append(stringSlice, resultString)
	}

	return strings.Join(stringSlice, "\n")
}