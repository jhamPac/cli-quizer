package main

import "flag"

func main() {
	csvFilename := flag.String("csv", "problems.csv", "A csv file in the format of 'question,answer'")
	flag.Parse()
	_ = csvFilename
}
