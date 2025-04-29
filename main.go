package main

import (
	"log"
	format "github.com/andreasSchauer/finalfantasyapi/data_formatting"
)





func main() {
	err := format.FormatJson()
	if err != nil {
		log.Fatal(err)
	}
}