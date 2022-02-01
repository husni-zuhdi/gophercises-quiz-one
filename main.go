package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	// Set flag to parse csv file.
	csvFileName := flag.String(
		"csv",
		"problems.csv",
		"a csv file in the format of 'question, answer'",
	)
	// Parse list of commands.
	flag.Parse()

	// Set and read the csv file
	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the csv file: %s\n", *csvFileName))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to load the csv file")
	}

	// Parse lines in csv file
	problems := parseLines(lines)

	// Iterate the problem
	for i, prob := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, prob.q)
	}
}

// exit to print exit message when something gone wrong.
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

// Breakdown the read lines into a new type of struct
type problem struct {
	q string
	a string
}

// parseLines to parse the csv file into question and answer
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines)) // Prepare variable ret to be returned
	for i, line := range lines {
		// Insert problem type struct into ret[i]
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return ret
}
