package Week_01

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	var slow = 0
	for faster := 1; faster < length; faster++ {
		if nums[faster] != nums[slow]{
			slow++
			nums[slow] = nums[faster]
		}
	}
	return slow + 1
}
