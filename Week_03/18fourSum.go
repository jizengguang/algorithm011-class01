package Week_03

import "sort"

func fourSum(nums []int, target int) [][]int {
	length := len(nums)
	if length < 4 {
		return nil
	}
	var result = make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < length-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		r := threeSum(nums[i+1:], target-nums[i])
		for _, v := range r {
			result = append(result, append(v, nums[i]))
		}
	}

	return result

}

func threeSum(nums []int, target int) [][]int {
	length := len(nums)
	var result = make([][]int, 0)
	var sum int
	sort.Ints(nums)

	for index := 1; index < length-1; index++ {
		var start, end = 0, length-1
		if index > 1 && nums[index] == nums[index-1] {
			start = index - 1
		}

		for start < index && index < end {
			sum = nums[index] + nums[start] + nums[end]
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}

			if end < length-1 && nums[end] == nums[end+1] {
				end--
				continue
			}

			if sum == target {
				result = append(result, []int{nums[index], nums[start], nums[end]})
				start++
				end--
			} else if sum < target {
				start++
			} else {
				end--
			}

		}

	}
	return result
}
