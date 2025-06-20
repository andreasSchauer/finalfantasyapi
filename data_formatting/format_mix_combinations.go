package dataFormatting

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"slices"
	"sort"
	"strings"
)

var itemNames = []string{
	"potion", "hi-potion", "x-potion", "mega-potion", "ether", "turbo ether", "elixir", "megalixir", "phoenix down", "mega phoenix", "antidote", "soft", "eye drops", "echo screen", "holy water", "remedy", "power distiller", "mana distiller", "speed distiller", "ability distiller", "al bhed potion", "healing water", "tetra elemental", "antarctic wind", "arctic wind", "ice gem", "bomb fragment", "bomb core", "fire gem", "electro marble", "lightning marble", "lightning gem", "fish scale", "dragon scale", "water gem", "grenade", "frag grenade", "sleeping powder", "dream powder", "silence grenade", "smoke bomb", "shadow gem", "shining gem", "blessed gem", "supreme gem", "poison fang", "silver hourglass", "gold hourglass", "candle of life", "petrify grenade", "farplane shadow", "farplane wind", "dark matter", "chocobo feather", "chocobo wing", "lunar curtain", "light curtain", "star curtain", "healing spring", "mana spring", "stamina spring", "soul spring", "purifying salt", "stamina tablet", "mana tablet", "stamina tonic", "mana tonic", "twin stars", "three stars", "power sphere", "mana sphere", "speed sphere", "ability sphere", "fortune sphere", "attribute sphere", "special sphere", "skill sphere", "wht magic sphere", "blk magic sphere", "master sphere", "lv. 1 key sphere", "lv. 2 key sphere", "lv. 3 key sphere", "lv. 4 key sphere", "hp sphere", "mp sphere", "strength sphere", "defense sphere", "magic sphere", "magic def sphere", "agility sphere", "evasion sphere", "accuracy sphere", "luck sphere", "clear sphere", "return sphere", "friend sphere", "teleport sphere", "warp sphere", "map", "rename card", "musk", "hypello potion", "shining thorn", "pendulum", "amulet", "designer wallet", "door to tomorrow", "wings to discovery", "gamblers spirit", "underdogs secret", "winning formula",
}

func FormatMixCombinations() error {

	const filepath = "./data/mix_data/mix_combinations_copy.txt"

	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("couldn't open file: %v", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("couldn't read file: %v", err)
	}

	mixCombinations, mixMapSlice := processMixCombinations(data)
	var bytes []byte
	bytes = fmt.Appendf(bytes, "* %s", mixCombinations)
	os.WriteFile(filepath, bytes, 0666)

	mixJsonData := getMixJsonData(mixMapSlice)
	mixCombinationsData := MixCombinationsData{
		MixCombinationsData: mixJsonData,
	}

	json, err := json.MarshalIndent(mixCombinationsData, "", "    ")
	if err != nil {
		return fmt.Errorf("couldn't encode JSON: %v", err)
	}

	os.WriteFile("./data/mix_combinations_copy.json", json, 0666)

	return nil
}


func getMixJsonData(mixMapSlice []MixMap) []Mix {
	var mixes []Mix

	for _, mixMap := range mixMapSlice {
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

		mix := Mix{
			Name: mixMap.Name,
			UniqueCombinations: mixMap.UniqueCombinations,
			MixCombinations: mixCombinations,
		}

		mixes = append(mixes, mix)
	}

	return mixes
}


type Combination struct {
	item1 string
	item2 string
}

type MixMap struct {
	Name            	string
	UniqueCombinations	int
	MixCombinations 	map[string][]string
}


type MixCombinationsData struct {
	MixCombinationsData []Mix `json:"mix_combinations_data"`
}


type Mix struct {
	Name            	string    	`json:"name"`
	UniqueCombinations	int			`json:"unique_combinations"`
	MixCombinations 	[]MixItems	`json:"mix_combinations"`
}

type MixItems struct {
	FirstItem           string   `json:"first_item"`
	PossibleSecondItems []string `json:"possible_second_items"`
}

func processMixCombinations(data []byte) (string, []MixMap) {
	mixData := string(data)
	mixes := strings.Split(mixData, "* ") // a slice of the individual mixes and their combinations
	var sortedMixes []string
	var mixMapSlice []MixMap

	for _, mix := range mixes {
		if mix == "" {
			continue
		}
		mixSplit := strings.Split(mix, "\n") // a mix and its combinations
		mixName := mixSplit[0]               // the name of the mix
		combinationsString := mixSplit[2:]   // the combinations of the mix (item 1 + item 2)
		combinations := getCombinationSlice(combinationsString)

		mixTable := getMixTable(mixName, combinations)
		sortedMixes = append(sortedMixes, mixTable)

		mixMap := MixMap{
			Name:            mixName,
			UniqueCombinations: len(combinations),
			MixCombinations: make(map[string][]string),
		}

		for _, combination := range combinations {
			_, ok := mixMap.MixCombinations[combination.item1]
			if !ok {
				mixMap.MixCombinations[combination.item1] = []string{}
			}

			mixMap.MixCombinations[combination.item1] = append(mixMap.MixCombinations[combination.item1], combination.item2)

			_, ok = mixMap.MixCombinations[combination.item2]
			if !ok {
				mixMap.MixCombinations[combination.item2] = []string{}
			}

			if combination.item1 != combination.item2 {
				mixMap.MixCombinations[combination.item2] = append(mixMap.MixCombinations[combination.item2], combination.item1)
			}
		}

		for _, combinationSlice := range mixMap.MixCombinations {
			sort.SliceStable(combinationSlice, func(i, j int) bool {
				return slices.Index(itemNames, combinationSlice[i]) < slices.Index(itemNames, combinationSlice[j])
			})
		}

		mixMapSlice = append(mixMapSlice, mixMap)
	}

	mixCombinations := strings.Join(sortedMixes, "\n\n\n\n\n* ")

	return mixCombinations, mixMapSlice
}

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

func splitCombination(combination string) (string, string) {
	items := strings.Split(combination, " + ")

	if slices.Index(itemNames, items[0]) <= slices.Index(itemNames, items[1]) {
		return items[0], items[1]
	}

	return items[1], items[0]
}
