package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// Set csvFileName flag to parse csv file.
	csvFileName := flag.String(
		"csv",
		"problems.csv",
		"a csv file in the format of 'question, answer'",
	)
	// Set timer flag
	timeLimit := flag.Int(
		"limit",
		30,
		"the time limit for the quiz in seconds",
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

	// Parse lines in csv file.
	problems := parseLines(lines)

	// Set the timer.
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	<-timer.C

	correct := 0
	// Iterate the problem.
problemloop:
	for i, prob := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, prob.q)
		// Make answer channel
		answerCh := make(chan string)
		// Build goroutine with anonymous function and run it.
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		// Waiting for timer to end. If it end, print the score.
		case <-timer.C:
			// Print score
			fmt.Printf("Your score is %d out of %d.\n", correct, len(problems))
			break problemloop
		// Waiting for user to answer the question. Then check the answer
		case answer := <-answerCh:
			if answer == prob.a {
				correct++
			}
		}
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
			a: strings.TrimSpace(line[1]), // Trim space in case the answer have too much space
		}
	}
	return ret
}
