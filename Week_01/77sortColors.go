package Week_01

func sortColors(nums []int) {
	if len(nums) < 1 {
		return
	}

	var zeroRight, twoLeft, curr = 0, len(nums)-1, 0
	for curr < twoLeft {
		if nums[curr] == 0 {
			nums[curr], nums[zeroRight] = nums[zeroRight], 0
			zeroRight++
			curr++
		} else if nums[curr] == 2 {
			nums[curr], nums[twoLeft] = nums[twoLeft], 2
			twoLeft--
		} else {
			curr++
		}
	}

}
