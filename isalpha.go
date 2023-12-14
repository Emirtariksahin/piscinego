package piscine

func IsAlpha(s string) bool {
	length := len(s)

	for i := 0; i < length; i++ {
		if !(s[i] >= 'a' && s[i] <= 'z' || s[i] >= 'A' && s[i] <= 'Z' || (s[i] >= '0' && s[i] <= '9')) {
			return false
		}
	}
	return true
}

/*Usage
Here is a possible program to test your function :

package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.IsAlpha("Hello! How are you?"))
	fmt.Println(piscine.IsAlpha("HelloHowareyou"))
	fmt.Println(piscine.IsAlpha("What's this 4?"))
	fmt.Println(piscine.IsAlpha("Whatsthis4"))

}
And its output :

$ go run .
false
true
false
true
$*/
