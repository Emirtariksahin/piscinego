package piscine

func StrRev(s string) string {
	sayi := []byte(s)

	for i, j := 0, len(sayi)-1; i < j; i, j = i+1, j-1 {
		sayi[i], sayi[j] = sayi[j], sayi[i]
	}

	return string(sayi)
}

/*Write a function that reverses a string.

This function will return the reversed string.

Expected function
func StrRev(s string) string {

}
Usage
Here is a possible program to test your function :

package main

import (
	"fmt"
	"piscine"
)

func main() {
	s := "Hello World!"
	s = piscine.StrRev(s)
	fmt.Println(s)
}
And its output :

$ go run .
!dlroW olleH
$*/
