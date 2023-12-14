package piscine

func CollatzCountdown(start int) int {
	if start < 1 {
		return -1
	}
	nb := start
	count := 0
	for nb > 1 {
		if nb%2 == 0 {
			nb = nb / 2
			count++
		} else if nb%2 == 1 {
			nb = nb*3 + 1
			count++
		}
	}

	return count
}
