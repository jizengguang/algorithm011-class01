package Week_01

func moveZeroes(nums []int) {


	var length = len(nums)
	var j = 0

	/**
	j 代表当前为0的数组下标
	第i个元素不等于0时，将第i个元素赋给当前为0的数组下标j处，如果i和j不相同
	那就需要把当前的元素换为0 （因为当前元素已经转移到前面为0的位置去了，当前当然应该赋值为0），
	原本记录0的位置因为已经不为0了，则往后挪动一位。
	 */
	for i := 0; i < length; i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			if i != j {
				nums[i] = 0
			}
			j++
		}

	}
}

