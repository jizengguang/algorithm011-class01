package Week_01

func maxArea(height []int) int {
	var left, right = 0, len(height)-1
	var max int

	for left < right {
		area := min(height[left], height[right]) * (right - left)
		if max < area {
			max = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}

	}
	return max

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
