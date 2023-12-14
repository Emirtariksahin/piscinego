package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	// Special case for zero
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	// Convert the number to a slice of digits
	digits := make([]int, 0)
	for n > 0 {
		digit := n % 10
		digits = append(digits, digit)
		n /= 10
	}

	// Sort the digits in ascending order
	for i := 0; i < len(digits)-1; i++ {
		for j := i + 1; j < len(digits); j++ {
			if digits[i] > digits[j] {
				digits[i], digits[j] = digits[j], digits[i]
			}
		}
	}

	// Print the sorted digits
	for i := 0; i < len(digits); i++ {
		z01.PrintRune(rune(digits[i]) + '0')
	}
}
