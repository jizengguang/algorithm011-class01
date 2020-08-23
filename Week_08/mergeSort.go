package Week_08

import (
	"fmt"
)

func mergeSort(arr []int, begin, end int) []int {
	l := len(arr)
	if l < 2 {
		return arr
	}
	var medium = begin + end>>1
	var left = arr[:medium]
	var right = arr[medium:]

	return merge(mergeSort(left, 0, len(left)-1), mergeSort(right, 0, len(right)-1))

}

func merge(left, right []int) []int {
	var result []int
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
		}
		result = append(result, right[0])
		right = right[1:]

	}

	for len(left) > 0 {
		result = append(result, left[0])
		left = left[1:]
	}

	for len(right) > 0 {
		result = append(result, right[0])
		right = right[1:]
	}

	return result
}

func main() {
	var nums = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(mergeSort(nums, 0, len(nums)-1))

}
