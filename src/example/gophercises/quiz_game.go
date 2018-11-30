package main

import (
	"flag"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of a 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		fmt.Printff("Failed to open the CSV file: %s", *csvFilename)
		os.Exit(1)
	}
	_ = file
}
