package piscine

// DescendAppendRange, verilen max ve min değerleri arasındaki tüm değerleri içeren bir dilim oluşturur.
// max dahil, min hariç.
func DescendAppendRange(max, min int) []int {
	// Eğer max, min'den küçük veya eşitse, boş bir dilim döndür
	if max <= min {
		return []int{}
	}

	// Oluşturulacak dilimi tutacak slice
	var result []int

	// max'ten min'e kadar olan değerleri slice'e ekleyerek oluştur
	for i := max; i > min; i-- {
		result = append(result, i)
	}

	// Oluşturulan dilimi döndür
	return result
}
