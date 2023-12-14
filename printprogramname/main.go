package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Get the program name from os.Args[0]
	programName := os.Args[0]

	// Find the index of the last '/' character in the programName
	lastSlashIndex := 0
	for i := len(programName) - 1; i >= 0; i-- {
		if programName[i] == '/' {
			lastSlashIndex = i + 1
			break
		}
	}

	// Extract the base name of the executable
	programName = programName[lastSlashIndex:]

	// Iterate over the characters of the programName and print each one
	for _, char := range programName {
		z01.PrintRune(char)
	}

	// Print a newline after printing the programName
	z01.PrintRune('\n')
}
