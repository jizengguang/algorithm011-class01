package Week_01

func climbStairs(n int) int {
	/**
	 滑动数组
	 step1 step2 result
	  0 0 1
	<-
	  0 1 1
	<-
	  1 1 2
	 */
	step1, step2, result := 0, 0, 1
	for i := 1; i <= n; i++ {
		step1, step2 = step2, result
		result = step1 + step2
	}
	return result
}
