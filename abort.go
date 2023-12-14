package piscine

func insertionSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		key := nums[i]
		j := i - 1

		for j >= 0 && nums[j] > key {
			nums[j+1] = nums[j]
			j = j - 1
		}
		nums[j+1] = key
	}
}

func Abort(a, b, c, d, e int) int {
	nums := []int{a, b, c, d, e}

	insertionSort(nums)

	return nums[2]
}
