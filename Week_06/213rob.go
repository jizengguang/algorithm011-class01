package Week_06

func rob(nums []int) int {
	//动态规划去写
	l := len(nums)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return nums[0]
	}
	if l == 2 {
		return max(nums[0], nums[1])
	}

	var dp1 = make([]int, l-1)
	var dp2 = make([]int, l-1)
	var dpNum1 = nums[:l-1]
	var dpNum2 = nums[1:]
	dp1[0] = dpNum1[0]
	dp1[1] = max(dpNum1[0], dpNum1[1])

	dp2[0] = dpNum2[0]
	dp2[1] = max(dpNum2[0], dpNum2[1])

	for i := 2; i < l-1; i++ {
		dp1[i] = max(dp1[i-1], dp1[i-2]+dpNum1[i])
		dp2[i] = max(dp2[i-1], dp2[i-2]+dpNum2[i])
	}
	return max(dp1[l-1], dp2[l-1])

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
