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
	csvFilename := flag.String("csv", "problems.csv", "A csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "The time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	problems := parseLines(lines)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct := 0
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

// Problem is a data structure that holds a question and answer from a line in a CSV file
type Problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []Problem {
	result := make([]Problem, len(lines))
	for i, line := range lines {
		result[i] = Problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return result
}
