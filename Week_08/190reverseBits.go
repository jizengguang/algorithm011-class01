package Week_08

func reverseBits(num uint32) uint32 {
	var res = uint32(0)
	var power = uint32(31)

	for num != 0 {
		res += (num & 1) << power
		num = num >> 1
		power -= 1
	}
	return res
}
