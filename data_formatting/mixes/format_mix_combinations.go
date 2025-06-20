package dataFormatting

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)



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
		mixSplit := strings.Split(mix, "\n") // a mix and its combinations
		mixName := mixSplit[0]               // the name of the mix
		combinationsString := mixSplit[2:]   // the combinations of the mix (item 1 + item 2)
		combinations := getCombinationSlice(combinationsString)

		mixTable := getMixTable(mixName, combinations)
		sortedMixes = append(sortedMixes, mixTable)

		mixMap := createMixMap(mixName, combinations)
		mixMapSlice = append(mixMapSlice, mixMap)
	}

	mixCombinations := strings.Join(sortedMixes, "\n\n\n\n\n* ")

	return mixCombinations, mixMapSlice
}






