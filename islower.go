package piscine

func IsLower(s string) bool {
	length := len(s)

	for i := 0; i < length; i++ {
		if !(s[i] >= 'a' && s[i] <= 'z') {
			return false
		}
	}

	return true
}
