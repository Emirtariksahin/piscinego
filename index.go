package piscine

func Index(s string, toFind string) int {
	for i := 0; i <= len(s)-len(toFind); i++ {
		if s[i:i+len(toFind)] == toFind {
			return i
		}
	}

	return -1
}

/*func Index(s string, toFind string) int {
	harf := []rune(s)
	bul := []rune(toFind)
	for j := 0; j <= len(s)-len(toFind); j++ {
		found := true
		for i := 0; i < len(toFind); i++ {
			if bul[i] != harf[j+i] {
				found = false
				break
			}
		}
		if found {
			return j
		}
	}
	return -1
}*/
