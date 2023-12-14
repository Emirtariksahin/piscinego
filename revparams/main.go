package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var programekle string

	for i := len(os.Args) - 1; i > 0; i-- {
		programekle = os.Args[i]

		for _, char := range programekle {
			z01.PrintRune(char)
		}

		if i != 1 {
			z01.PrintRune('\n')
		}
	}
	z01.PrintRune('\n')
}
