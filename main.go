package main

import (
	"fmt"
	"os"
	"time"
)

// Simple cli tool that accepts a time duration and prints it in a more human readable format.
func main() {
	// get the duration from the command line
	duration := os.Args[1]

	// parse the duration
	d, err := time.ParseDuration(duration)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// print the duration
	fmt.Println(d)

}
