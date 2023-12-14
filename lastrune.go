package piscine

func LastRune(s string) rune {
	sonharf := []rune(s)

	for i := len(s) - 1; i < len(s); i++ {
		return sonharf[i]
	}

	return 0
}
