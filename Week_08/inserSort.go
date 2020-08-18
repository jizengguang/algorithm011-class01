package Week_08

import "fmt"

func insertSort(arr []int) {
	l := len(arr)
	for i := 1; i < l; i++ {
		for j := 0; j < i; j++ {
			if arr[j] > arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
}

func main() {
	var nums = []int{9,8,7,6,5,4,3,2,1}
	insertSort(nums)
	fmt.Println(nums)
}
