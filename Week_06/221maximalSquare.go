package Week_06

func maximalSquare(matrix [][]byte) int {
	//动态规划方程
	//1、找到重复子问题
	//2、定义状态
	//3、方程

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	ans, d := 0, make([][]int, len(matrix))
	for i, m, n := 0, len(matrix), len(matrix[0]); i < m; i++ {
		d[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				if i == 0 || j == 0 {
					d[i][j] = 1
				} else {
					d[i][j] = min(d[i-1][j], d[i][j-1], d[i-1][j-1]) + 1
				}
			}
			if d[i][j] > ans {
				ans = d[i][j]
			}
		}
	}
	return ans * ans
}

func min(x, y, z int) int {
	if x > y {
		x = y
	}
	if x > z {
		return z
	}
	return x
}
