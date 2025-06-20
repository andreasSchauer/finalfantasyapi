package dataFormatting

import(
	"slices"
	"sort"
)


func getMixJsonData(mixMapSlice []MixMap) []Mix {
	var mixes []Mix

	for _, mixMap := range mixMapSlice {
		mixCombinations := getMixCombinations(mixMap)

		mix := Mix{
			Name: mixMap.Name,
			UniqueCombinations: mixMap.UniqueCombinations,
			MixCombinations: mixCombinations,
		}

		mixes = append(mixes, mix)
	}

	return mixes
}



func getMixCombinations(mixMap MixMap) []MixItems {
	var mixCombinations []MixItems

	for firstItem := range mixMap.MixCombinations {
		mixItems := MixItems{
			FirstItem: firstItem,
			PossibleSecondItems: mixMap.MixCombinations[firstItem],
		}

		mixCombinations = append(mixCombinations, mixItems)
	}

	sort.SliceStable(mixCombinations, func(i, j int) bool {
		return slices.Index(itemNames, mixCombinations[i].FirstItem) < slices.Index(itemNames, mixCombinations[j].FirstItem)
	})

	return mixCombinations
}