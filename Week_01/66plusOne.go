package Week_01

func plusOne(digits []int) []int {
	/**
	 末尾为9的情况下，加上1 即为0
	 下一层循环到前一位加1
	 其他情况加1返回即可
	 */
	length := len(digits)
	for i := length - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] += 1
			return digits
		}
		digits[i] = 0
	}
	digits = append([]int{1}, digits...)
	return digits
}
