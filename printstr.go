package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for _, ekle := range s {
		z01.PrintRune(ekle)
	}
	z01.PrintRune('\n')
}

/*Write a function that prints one by one the characters of a string on the screen.
Expected function
func PrintStr(s string) {

}
Hints
Here is a possible program to test your function :

package main

import "piscine"

func main() {
	piscine.PrintStr("Hello World!")
}
And its output :

go run . | cat -e
Hello World!*/
