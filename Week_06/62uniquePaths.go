package Week_06

func uniquePaths(m int, n int) int {
	//1、寻找重复子问题
	// 第i，j 只能来自 第i-1，j 和 第i,j-1的路径
	//2、定义状态
	// 到i,j的路径和等于第i-1,j 和第i,j-1的路径和
	//3、动规方程
	//dp[i][j] = dp[i-1][j] + dp[i][j-1]

	var dp = make([]int, m*n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i*n+j] = 1
			} else {
				dp[i*n+j] = dp[(i-1)*n+j] + dp[i*n+j-1]
			}

		}
	}
	return dp[m*n-1]

}
