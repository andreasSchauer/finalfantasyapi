package dataFormattingMixes

import(
	"strings"
)


type MixData struct {
	Name               		string     		`json:"name"`
	Category				string			`json:"category"`
	UniqueCombinations 		int        		`json:"unique_combinations"`
	BestCombinations		[]Combination	`json:"best_combinations"`
	PossibleCombinations	[]Combination	`json:"possible_combinations"`
}


func getMixData(mix string) MixData {
	mixSplit := strings.Split(mix, "\n")
	mixData := mixSplit[0] // name / category / best mixes
	dataSplit := strings.Split(mixData, " / ")
	mixName := dataSplit[0]
	mixCategory := dataSplit[1]
	bestCombosString := strings.Split(dataSplit[2], "; ")
	bestCombos := getCombinations(bestCombosString)
	
	combinationsString := mixSplit[2:]   // the combinationStrings of the mix (= item 1 + item 2)
	combinations := getCombinations(combinationsString)

	return MixData{
		Name: mixName,
		Category: mixCategory,
		UniqueCombinations: len(combinations),
		BestCombinations: bestCombos,
		PossibleCombinations: combinations,
	}
}

