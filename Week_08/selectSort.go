package Week_08

import "fmt"

func selectSort(nums []int) {
	l := len(nums)
	for i := 0; i < l-1; i++ {
		var min = i
		for j := i + 1; j < l; j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}
}

func main() {
	var nums = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	selectSort(nums)
	fmt.Println(nums)
}
