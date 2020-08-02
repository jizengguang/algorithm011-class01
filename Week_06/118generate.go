package Week_06

func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}

	var res = make([][]int, 0)
	res = append(res, []int{1})

	for i := 2; i <= numRows; i++ {
		var level = make([]int, i)
		level[0] = 1
		level[i-1] = 1
		for j := 1; j <= i-2; j++ {
			var pre = res[i-2]
			level[j] = pre[j-1] + pre[j]
		}
		res = append(res, level)
	}
	return res

}
