package main

import (
	"log"
	format "github.com/andreasSchauer/finalfantasyapi/data_formatting"
)





func main() {
	err := format.FormatMixes()
	if err != nil {
		log.Fatal(err)
	}
}