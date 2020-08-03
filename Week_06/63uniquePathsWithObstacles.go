package Week_06

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	//1、找到重复子问题
	//2、定义状态
	//3、动规方程
	var m = len(obstacleGrid)
	var n = len(obstacleGrid[0])

	var dp = make([]int, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				continue
			}
			if i == 0 && j == 0 {
				dp[0] = 1
			} else if i == 0  {
				dp[i*n+j] = dp[i*n+j-1]
			} else if j == 0  {
				dp[i*n+j] = dp[(i-1)*n+j]
			} else {
				dp[i*n+j] = dp[(i-1)*n+j] + dp[i*n+j-1]
			}
		}
	}
	return dp[m*n-1]

}
