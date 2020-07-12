package Week_03

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	i, j, k := len(nums)-2, len(nums)-1, len(nums)-1

	//从j到end为降序的
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}

	//到nums[i] < nums[j]时，从k到j查找第一个小于nums[i]的数，交换。

	if i >= 0 {
		for nums[i] >= nums[k] {
			k--
		}
		nums[i], nums[k] = nums[k], nums[i]
	}

	//从j到end转做一个升序
	for i, j := j, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}

}
