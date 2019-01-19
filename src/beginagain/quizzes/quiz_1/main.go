package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s", *csvFilename))
		os.Exit(1)
	}
	// create csv reader
	r := csv.NewReader(file)
	// We want to parse the csv
	lines, err := r.ReadAll()
	// see if there is an error
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}
	problems := parseLines(lines)
	// iterate over all problems and print them out, get response from user, check to see if user is correct
	// keep track of how many problems we get correct:
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s \n", i+1, p.q)
		// next step, read in an answer. Need variable to store it in.
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

// we need some way to create problems via 2d slice

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	// fill in every value of slice with one of the lines
	for i, line := range lines {
		ret[i] = problem{
			// doing this because first column is question, second column is answer.
			q: line[0],
			// trim spaces to ignore them!
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

// take lines, build problem type
type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
