package Week_03

func generateMatrix(n int) [][]int {
	//设定边界值
	top, bo, left, right := 0, n-1, 0, n-1
	var res = [][]int{}
	for i := 0; i < n; i++ {
		res = append(res, make([]int, n))
	}

	var num = 1
	var tar = n * n

	for num <= tar {
		for i := left; i <= right; i++ {
			res[top][i] = num
			num += 1
		}
		top++

		for i := top; i <= bo; i++ {
			res[i][right] = num
			num += 1
		}
		right--

		for i := right; i >= left; i-- {
			res[bo][i] = num
			num += 1
		}
		bo--

		for i := bo; i >= top; i-- {
			res[i][left] = num
			num += 1
		}
		left++
	}
	return res
}
