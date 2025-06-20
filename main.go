package main

import (
	"log"
	format "github.com/andreasSchauer/finalfantasyapi/data_formatting"
)





func main() {
	err := format.FormatMixCombinations()
	if err != nil {
		log.Fatal(err)
	}
}