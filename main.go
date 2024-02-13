package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define command-line flags
	countBytes := flag.Bool("c", false, "count bytes")

	// Parse command-line flags
	flag.Parse()

	// Get the filename from the command-line arguments
	filename := flag.Arg(0)

	// Check if the -c flag is provided
	if *countBytes {
		// Read the contents of the file
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
			os.Exit(1)
		}

		// Get the number of bytes
		byteCount := len(data)

		// Print the result
		fmt.Printf("%d %s\n", byteCount, filename)
	} else {
		fmt.Println("Usage: ccwc -c filename")
	}
}
