package piscine

func ToLower(s string) string {
	result := []rune(s)

	for i := 0; i < len(result); i++ {
		if result[i] >= 'A' && result[i] <= 'Z' {
			// Convert lowercase letter to uppercase
			result[i] = result[i] - 'A' + 'a'
		}
	}

	return string(result)
}
