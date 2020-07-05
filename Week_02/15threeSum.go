package Week_02

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	var sum int
	//index为中间值
	var index, start, end int
	for index = 1; index < len(nums)-1; index++ {
		start, end = 0, len(nums)-1
		//去重
		if index > 1 && nums[index] == nums[index-1] {
			start = index - 1
		}
		for start < index && index < end {
			sum = nums[start] + nums[index] + nums[end]
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}
			if end < len(nums)-1 && nums[end] == nums[end+1] {
				end--
				continue
			}
			if sum == 0 {
				result = append(result, []int{nums[start], nums[index], nums[end]})
				start++
				end--
			} else if sum > 0 {
				end--
			} else {
				start++
			}
		}
	}
	return result
}
