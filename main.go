package main

import (
	"log"
	// mixes "github.com/andreasSchauer/finalfantasyapi/data_formatting/mixes"
	// monstersOld "github.com/andreasSchauer/finalfantasyapi/data_formatting/monstersOld"
	monstersNew "github.com/andreasSchauer/finalfantasyapi/data_formatting/monstersNew"
)

func main() {
	err := monstersNew.FormatNewMonsterJson()
	if err != nil {
		log.Fatal(err)
	}
}
