package piscine

func IsNumeric(s string) bool {
	length := len(s)
	for i := 0; i < length; i++ {
		if !(s[i] >= '0' && s[i] <= '9') {
			return false
		}
	}
	return true
}
