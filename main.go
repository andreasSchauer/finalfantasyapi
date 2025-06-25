package main

import (
	"log"
	// mixes "github.com/andreasSchauer/finalfantasyapi/data_formatting/mixes"
	monsters "github.com/andreasSchauer/finalfantasyapi/data_formatting/monsters"
)





func main() {
	err := monsters.FormatMonsterJson()
	if err != nil {
		log.Fatal(err)
	}
}