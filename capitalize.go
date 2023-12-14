package piscine

func isAlphanumeric(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}

func Capitalize(s string) string {
	result := make([]rune, len(s))
	previousIsAlphanumeric := false

	for i, char := range s {
		if isAlphanumeric(char) {
			if !previousIsAlphanumeric {
				if char >= 'a' && char <= 'z' {
					result[i] = char - 'a' + 'A'
				} else {
					result[i] = char
				}
			} else {
				if char >= 'A' && char <= 'Z' {
					result[i] = char - 'A' + 'a'
				} else {
					result[i] = char
				}
			}
			previousIsAlphanumeric = true
		} else {
			result[i] = char
			previousIsAlphanumeric = false
		}
	}

	return string(result)
}
