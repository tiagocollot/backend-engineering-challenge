package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Declare a variable "input" which is a pointer to a string, initialized with the value "input.json"
	//  and a string describing the input file
	input := flag.String("input_file", "input.json", "input file")
	// Declare a variable "windowsize" which is a pointer to a uint, initialized with the value 10
	// and a string describing the window size
	windowsize := flag.Uint("window_size", 10, "window size")
	// parse the command-line flags
	flag.Parse()

	// Open the file specified by the input flag
	f, err := os.Open(*input)
	// If there is an error opening the file, print an error message and return
	if err != nil {
		fmt.Printf("It was not possible to open the file '%s': %s", *input, err)
		return
	}

	// call the calculateAvg function passing the file, the standard output and the window size
	if err := calculateAvg(f, os.Stdout, *windowsize); err != nil {
		fmt.Printf("It was not possible to calculate the average: %s", err)
		return
	}
}
