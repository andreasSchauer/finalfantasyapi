package dataFormattingMixes

import(
	"fmt"
	"strings"
)


func getMixTable(mixName string, combinations []Combination) string {
	combinationsRejoined := rejoinCombinations(combinations)
	sortedString := strings.Join(combinationsRejoined, "\n")
	mixTable := fmt.Sprintf("%s\n\n%s", mixName, sortedString)

	return mixTable
}



func rejoinCombinations(combinations []Combination) []string {
	var rejoinedCombos []string

	for _, combination := range combinations {
		joinedString := fmt.Sprintf("%s + %s", combination.item1, combination.item2)
		rejoinedCombos = append(rejoinedCombos, joinedString)
	}

	return rejoinedCombos
}
