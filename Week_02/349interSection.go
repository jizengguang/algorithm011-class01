package Week_02

func intersection(nums1 []int, nums2 []int) []int {

	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}

	var keyMap = make(map[int]bool, 0)
	var result = make([]int, 0)

	for _, v := range nums1 {
		keyMap[v] = true
	}

	for _, value := range nums2 {
		if vl, ok := keyMap[value]; ok && vl {
			keyMap[value] = false
			result = append(result, value)
		}
	}

	return result
}
