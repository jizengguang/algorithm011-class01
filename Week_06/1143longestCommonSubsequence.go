package Week_06

func longestCommonSubsequence(text1 string, text2 string) int {
	//1、暴力轮询
	//O(m*n)

	//2、动态规划
	//a.寻找重复子问题
	//b.定义状态
	//dp 符合s1和s2的最长子序列
	//索引为0的全部表示空串，则dp[0][..] = 0, dp[..][0] = 0
	//c.动规方程
	//两个指针 i，j 如果s[i]=s[j] 那么一定在lcs里，如果不等于，那么至少有一个字符不在lcs中。就要丢弃一个字母。

	//初始化dp数组
	/**
		var dp = make([][]int, len(text1)+1)
	for i := 0; i < len(text1)+1; i++ {
		dp[i] = make([]int, len(text2)+1)
	}

	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = maxDp(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	return dp[len(text1)][len(text2)]

	*/

	//一维优化

	var dp = make([]int, len(text2)+1)

	for i := 1; i <= len(text1); i++ {
		var last int
		for j := 1; j <= len(text2); j++ {
			var tmp = dp[j]
			if text1[i-1] == text2[j-1] {
				dp[j] = last + 1
			} else {
				dp[j] = maxDp(dp[j-1], tmp)
			}
			last = tmp
		}
	}
	return dp[len(text2)]

}
func maxDp(a, b int) int {
	if a > b {
		return a
	}
	return b
}
