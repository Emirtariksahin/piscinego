package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 21 {
		return 0 // Return 0 for values outside the valid range
	}

	result := int(1)

	for i := 1; i <= nb; i++ {
		result *= int(i)
	}

	return result
}
