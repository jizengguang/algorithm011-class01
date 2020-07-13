package Week_04

import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	/*
		1、哈希表法
		    // 时间复杂度：O(M)+O(N)
		    // 空间复杂度：O(min(M,N))
			var record = make(map[int]int, 0)
			var result = make([]int, 0)

			for i := 0; i < len(nums1); i++ {
				if _, ok := record[nums1[i]]; ok {
					record[nums1[i]] += 1
					continue
				}
				record[nums1[i]] = 1
			}
			for j := 0; j < len(nums2); j++ {
				if value, ok := record[nums2[j]]; ok && value != 0 {
					result = append(result, nums2[j])
					record[nums2[j]] -= 1
				}
			}
			return result
	*/

	//2、sort
	//时间复杂度 O(mlogm+nlogn)
	sort.Ints(nums1)
	sort.Ints(nums2)
	var result = make([]int, 0)
	var s1Index, s2Index = 0, 0
	for s1Index < len(nums1) && s2Index < len(nums2) {
		if nums1[s1Index] == nums2[s2Index] {
			result = append(result, nums1[s1Index])
			s1Index++
			s2Index++
		} else if nums1[s1Index] < nums2[s2Index] {
			s1Index++
		} else {
			s2Index++
		}
	}
	return result
}
