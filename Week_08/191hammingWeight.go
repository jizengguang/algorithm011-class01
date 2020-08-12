package Week_08

func hammingWeight(num uint32) int {

	/**
	1、循环和位移动
	var mask uint32 = 1
	var res = 0

	for i := 0; i < 32; i++ {
		if (num & mask) != 0 {
			res++
		}
		mask <<= 1
	}
	return  res
	*/

	/**
	2、代码优化
	*/
	var res = 0

	for num != 0 {
		res++
		num &= num - 1
	}
	return res

}
