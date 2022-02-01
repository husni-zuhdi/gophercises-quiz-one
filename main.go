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
	fmt.Println(lines)
}

// exit to print exit message when something gone wrong.
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
