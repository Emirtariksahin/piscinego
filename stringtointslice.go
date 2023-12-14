package piscine

func StringToIntSlice(str string) []int {
	sonuc := []int{}
	if str == "" {
		return nil
	}
	for _, v := range str {
		sonuc = append(sonuc, int(v))
	}
	return sonuc
}
