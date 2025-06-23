package dataFormattingMixes

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type MixCombinationsData struct {
	MixesData []MixData `json:"mixes_data"`
}



func FormatMixCombinations() error {
	combinationsFilePath := "./data/mix_data/mix_combinations_copy.txt"
	jsonFilePath := "./data/mixes_copy.json"
	// itemResultsFilePath := "./data/mix_data/item_results_copy.txt"

	file, err := os.Open(combinationsFilePath)
	if err != nil {
		return fmt.Errorf("couldn't open file: %v", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("couldn't read file: %v", err)
	}

	mixTables, mixDataSlice := processMixCombinations(data)
	var bytesMixTables []byte
	bytesMixTables = fmt.Appendf(bytesMixTables, "* %s", mixTables)
	os.WriteFile(combinationsFilePath, bytesMixTables, 0666)
	
	/*
	var bytesItemResults []byte
	bytesItemResults = fmt.Append(bytesItemResults, itemResults)
	os.WriteFile(itemResultsFilePath, bytesItemResults, 0666)
	
	mixJsonData := getMixJsonData(mixDataSlice)
	*/

	mixCombinationsData := MixCombinationsData{
		MixesData: mixDataSlice,
	}

	json, err := json.MarshalIndent(mixCombinationsData, "", "    ")
	if err != nil {
		return fmt.Errorf("couldn't encode JSON: %v", err)
	}

	os.WriteFile(jsonFilePath, json, 0666)

	return nil
}

func processMixCombinations(data []byte) (string, []MixData) {
	mixDataRaw := string(data)
	mixes := strings.Split(mixDataRaw, "* ") // a slice of the individual mixes and their combinations
	var mixTables []string
	// var itemResults []ItemResult
	// itemResultsPtr := &itemResults
	// var mixMapSlice []MixMap
	var mixDataSlice []MixData

	for _, mix := range mixes {
		if mix == "" {
			continue
		}
		mixData := getMixData(mix)
		mixDataSlice = append(mixDataSlice, mixData)

		mixTable := mixData.getMixTable()
		mixTables = append(mixTables, mixTable)
		
		/*
		appendItemResults(itemResultsPtr, mixName, combinations)

		mixMap := createMixMap(mixName, combinations)
		mixMapSlice = append(mixMapSlice, mixMap)
		*/
	}

	mixTablesJoined := strings.Join(mixTables, "\n\n\n\n\n* ")

	/*
	sortItemResults(itemResultsPtr)
	formattedItemResults := formatItemResults(itemResultsPtr)
	*/

	return mixTablesJoined, mixDataSlice
}
