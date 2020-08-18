package Week_08

import "fmt"

func bubbleSort(nums []int) {
	l := len(nums)
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

func main() {
	var nums = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	bubbleSort(nums)
	for _, v := range nums {
		fmt.Println(v)
	}
}
