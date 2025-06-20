package dataFormattingMixes

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type MixCombinationsData struct {
	MixCombinationsData []Mix `json:"mix_combinations_data"`
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

func FormatMixCombinations() error {
	combinationsFilePath := "./data/mix_data/mix_combinations_copy.txt"
	jsonFilePath := "./data/mix_combinations_copy.json"

	file, err := os.Open(combinationsFilePath)
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
	os.WriteFile(combinationsFilePath, bytes, 0666)

	mixJsonData := getMixJsonData(mixMapSlice)
	mixCombinationsData := MixCombinationsData{
		MixCombinationsData: mixJsonData,
	}

	json, err := json.MarshalIndent(mixCombinationsData, "", "    ")
	if err != nil {
		return fmt.Errorf("couldn't encode JSON: %v", err)
	}

	os.WriteFile(jsonFilePath, json, 0666)

	return nil
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
		mixSplit := strings.Split(mix, "\n") // a mix and its combinations as strings
		mixName := mixSplit[0]               // the name of the mix
		combinationsString := mixSplit[2:]   // the combinationStrings of the mix (= item 1 + item 2)
		combinations := getCombinations(combinationsString)

		mixTable := getMixTable(mixName, combinations)
		sortedMixes = append(sortedMixes, mixTable)

		mixMap := createMixMap(mixName, combinations)
		mixMapSlice = append(mixMapSlice, mixMap)
	}

	mixCombinations := strings.Join(sortedMixes, "\n\n\n\n\n* ")

	return mixCombinations, mixMapSlice
}

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
