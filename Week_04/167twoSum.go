package Week_04

func twoSum(numbers []int, target int) []int {
	/**
	1、暴力解法
	时间复杂度O(N)
	空间复杂度O(1)
	var index1,index2 = 0,0
	for ; index1 <= len(numbers)-2; index1++ {
		if numbers[index1] > target {
			return nil
		}

		tmp := target - numbers[index1]
		if tmp < numbers[index1] {
			return nil
		}

		index2 = index1 + 1
		for ; index2 <= len(numbers)-1; index2++ {
			if tmp == numbers[index2] {
				return []int{index1+1, index2+1}
			}
		}
	}
	return nil

	*/

	//2、二分查找
	//时间复杂度 O(nlogN)
	//空间复杂度 O(1)
	for k, v := range numbers {
		if v > target {
			return nil
		}
		var start, end = k+1, len(numbers)-1
		for start <= end {
			medium := (end-start)/2 + start
			if numbers[medium] == target-v {
				return []int{k + 1, medium + 1}
			}
			if numbers[medium] > target-v {
				end = medium - 1
			} else {
				start = medium + 1
			}
		}
	}
	return nil

}
