package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	input := flag.String("input_file", "input.json", "input file")
	windowsize := flag.Uint("window_size", 10, "window size")
	flag.Parse()

	f, err := os.Open(*input)
	if err != nil {
		fmt.Printf("It was not possible to open the file '%s': %s", *input, err)
		return
	}

	if err := calculateAvg(f, os.Stdout, *windowsize); err != nil {
		fmt.Printf("It was not possible to calculate the average: %s", err)
		return
	}
}
