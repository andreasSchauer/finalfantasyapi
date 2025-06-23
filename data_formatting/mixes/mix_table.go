package dataFormattingMixes

import (
	"fmt"
	"strings"
)

func (m MixData)getMixTable() string {
	possCombosRejoined := rejoinCombinations(m.PossibleCombinations)
	sortedString := strings.Join(possCombosRejoined, "\n")
	firstRow := m.rejoinFirstRow()

	mixTable := fmt.Sprintf("%s\n\n%s", firstRow, sortedString)

	return mixTable
}


func (m MixData)rejoinFirstRow() string {
	bestCombosRejoined := rejoinCombinations(m.BestCombinations)
	bestCombosString := strings.Join(bestCombosRejoined, "; ")

	firstRow := fmt.Sprintf("%s / %s / %s", m.Name, m.Category, bestCombosString)
	return firstRow
}


func rejoinCombinations(combinations []Combination) []string {
	var rejoinedCombos []string

	for _, combination := range combinations {
		joinedString := fmt.Sprintf("%s + %s", combination.Item1, combination.Item2)
		rejoinedCombos = append(rejoinedCombos, joinedString)
	}

	return rejoinedCombos
}
