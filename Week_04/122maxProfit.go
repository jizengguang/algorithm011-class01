package Week_04

func maxProfit(prices []int) int {
	//1、贪心算法
	//时间复杂度O(N)
	//空间复杂度O(1)
	var res int = 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}

	return res

}
