package Week_03

func pown(x float64, n int) float64 {
	//2、递归
	//时间复杂度：O(logN)

	//n < 0 情况
	if n < 0 {
		n *= -1
		x = 1 / x
	}
	return fast(x, n)

}

func fast(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}

	subResult := fast(x, n/2)

	//判断奇偶决定是否多乘一个X
	if n&1 == 1 {
		return subResult * subResult * x
	}
	return subResult * subResult

}
