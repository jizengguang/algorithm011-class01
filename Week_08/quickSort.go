package Week_08

import (
	"fmt"
)

func quickSort(arr []int, begin, end int) {
	if begin >= end {
		return
	}
	//确定基线
	var pivot = arr[begin]
	var i = begin + 1

	for j := begin; j <= end; j++ {
		if pivot > arr[j] {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[begin], arr[i-1] = arr[i-1], pivot

	quickSort(arr, begin, i-2)
	quickSort(arr, i, end)

}

func main() {
	var nums = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
