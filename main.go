package main

import (
	"log"
	mixes "github.com/andreasSchauer/finalfantasyapi/data_formatting/mixes"
)





func main() {
	err := mixes.FormatMixCombinations()
	if err != nil {
		log.Fatal(err)
	}
}