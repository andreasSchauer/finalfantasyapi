package dataFormatting

import(
	"slices"
	"sort"
)



func createMixMap(mixName string, combinations []Combination) MixMap {
	mixMap := MixMap{
		Name:            	mixName,
		UniqueCombinations: len(combinations),
		MixCombinations: 	make(map[string][]string),
	}

	for _, combination := range combinations {
		mixMap.populateMixCombinations(combination.item1, combination.item2)
		mixMap.populateMixCombinations(combination.item2, combination.item1)
	}

	for _, combinationSlice := range mixMap.MixCombinations {
		sort.SliceStable(combinationSlice, func(i, j int) bool {
			return slices.Index(itemNames, combinationSlice[i]) < slices.Index(itemNames, combinationSlice[j])
		})
	}

	return mixMap
}



func (m MixMap) populateMixCombinations(itemA, itemB string) {
	_, ok := m.MixCombinations[itemA]
	if !ok {
		m.MixCombinations[itemA] = []string{}
	}

	if itemA != itemB && slices.Index(m.MixCombinations[itemA], itemB) == -1 {
		m.MixCombinations[itemA] = append(m.MixCombinations[itemA], itemB)
	}

	if itemA == itemB && slices.Index(m.MixCombinations[itemA], itemB) == -1 {
		m.MixCombinations[itemA] = append(m.MixCombinations[itemA], itemB)
	}
}