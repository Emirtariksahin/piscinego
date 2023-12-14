package piscine

func IsUpper(s string) bool {
	length := len(s)

	for i := 0; i < length; i++ {
		if !(s[i] >= 'A' && s[i] <= 'Z') {
			return false
		}
	}

	return true
}
