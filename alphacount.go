package piscine

func AlphaCount(s string) int {
	nb := 0
	length := len(s)
	harf := []rune(s)

	for i := 0; i < length; i++ {
		if (harf[i] >= 'a' && harf[i] <= 'z') || (harf[i] >= 'A' && harf[i] <= 'Z') {
			nb++
		}
	}

	return nb
}

/*Usage
Here is a possible program to test your function :

package main

import (
	"fmt"
	"piscine"
)

func main() {
	s := "Hello 78 World!    4455 /"
	nb := piscine.AlphaCount(s)
	fmt.Println(nb)
}
And its output :

$ go run .
10
$*/
