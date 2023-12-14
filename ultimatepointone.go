package piscine

func UltimatePointOne(n ***int) {
	***n += 1
}

/*func UltimatePointOne(n ***int) {

}
Usage
Here is a possible program to test your function :

package main

import (
	"fmt"
	"piscine"
)

func main() {
	a := 0
	b := &a
	n := &b
	piscine.UltimatePointOne(&n)
	fmt.Println(a)
}
And its output :

$ go run .
1
$*/
