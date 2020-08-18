package Week_08

func selectSort(nums []int) {
	l := len(nums)
	if l <= 1 {
		return
	}
	for i := 0; i < l-1; i++ {
		var min = i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}
}
