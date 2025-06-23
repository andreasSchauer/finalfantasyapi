package dataFormattingMixes

import (
	"slices"
	"sort"
)

type MixMap struct {
	Name               string
	UniqueCombinations int
	MixCombinations    map[string][]string
}


type Mix struct {
	Name               string     `json:"name"`
	UniqueCombinations int        `json:"unique_combinations"`
	MixCombinations    []MixItems `json:"mix_combinations"`
}


type MixItems struct {
	FirstItem           string   `json:"first_item"`
	PossibleSecondItems []string `json:"possible_second_items"`
}




func createMixMap(mixName string, combinations []Combination) MixMap {
	mixMap := MixMap{
		Name:               mixName,
		UniqueCombinations: len(combinations),
		MixCombinations:    make(map[string][]string),
	}

	for _, combination := range combinations {
		mixMap.populateMixCombinations(combination.Item1, combination.Item2)
		mixMap.populateMixCombinations(combination.Item2, combination.Item1)
	}

	for _, secondItemSlice := range mixMap.MixCombinations {
		sort.SliceStable(secondItemSlice, func(i, j int) bool {
			return slices.Index(itemNames, secondItemSlice[i]) < slices.Index(itemNames, secondItemSlice[j])
		})
	}

	return mixMap
}

func (m MixMap) populateMixCombinations(itemA, itemB string) {
	_, ok := m.MixCombinations[itemA]
	if !ok {
		m.MixCombinations[itemA] = []string{}
	}

	// if item is not in slice yet
	if slices.Index(m.MixCombinations[itemA], itemB) == -1 {
		m.MixCombinations[itemA] = append(m.MixCombinations[itemA], itemB)
	}
}

func (m MixMap) getJsonMixCombinations() []MixItems {
	var mixCombinations []MixItems

	for firstItem := range m.MixCombinations {
		mixItems := MixItems{
			FirstItem:           firstItem,
			PossibleSecondItems: m.MixCombinations[firstItem],
		}

		mixCombinations = append(mixCombinations, mixItems)
	}

	sort.SliceStable(mixCombinations, func(i, j int) bool {
		return slices.Index(itemNames, mixCombinations[i].FirstItem) < slices.Index(itemNames, mixCombinations[j].FirstItem)
	})

	return mixCombinations
}



/*
func getMixJsonData(mixMapSlice []MixMap) []Mix {
	var mixes []Mix

	for _, mixMap := range mixMapSlice {
		mixCombinations := mixMap.getJsonMixCombinations()

		mix := Mix{
			Name:               mixMap.Name,
			UniqueCombinations: mixMap.UniqueCombinations,
			MixCombinations:    mixCombinations,
		}

		mixes = append(mixes, mix)
	}

	return mixes
}
*/