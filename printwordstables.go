package piscine

import (
	"github.com/01-edu/z01"
)

func PrintWordsTables(a []string) {
	var words []string
	harfler := ""

	for _, kelime := range a {
		for _, karakter := range kelime {
			if karakter == ' ' || karakter == '\t' || karakter == '\n' {
				if harfler != "" {
					words = append(words, harfler)
					harfler = ""
				}
			} else {
				harfler += string(karakter)
			}
		}
	}

	if harfler != "" {
		words = append(words, harfler)
	}

	// Ayırdığımız kelimeleri yazdıralım.
	for _, word := range a {
		for _, char := range word {
			z01.PrintRune(char)
		}

		z01.PrintRune('\n') // Add a newline after each word
	}
}
