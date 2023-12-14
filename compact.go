package piscine

func Compact(ptr *[]string) int {
	slice := *ptr
	var nonZeroCount int

	// Compact the slice by removing zero values
	for i := 0; i < len(slice); i++ {
		if slice[i] != "" {
			slice[nonZeroCount] = slice[i]
			nonZeroCount++
		}
	}

	// Resize the slice to the number of non-zero elements
	*ptr = slice[:nonZeroCount]

	return nonZeroCount
}
