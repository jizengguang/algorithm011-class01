package Week_04

var mutationMap = map[uint8][3]string{
	'A': {"T", "G", "C"},
	'T': {"A", "G", "C"},
	'G': {"A", "T", "C"},
	'C': {"A", "T", "G"},
}

func minMutation(start string, end string, bank []string) int {
	if isExist(end, bank) == -1 {
		return -1
	}

	minCount := len(bank) + 1
	isUsed := make([]bool, len(bank))

	var dfsMinMutation func(curr string, count int)

	dfsMinMutation = func(curr string, count int) {
		if count >= minCount {
			return
		}

		if curr == end {
			if count < minCount {
				minCount = count
				return
			}
		}

		for i := 0; i < len(curr); i++ {
			for _, value := range mutationMap[curr[i]] {
				gene := curr[:i] + value + curr[i+1:]
				if index := isExist(gene, bank); index != -1 && !isUsed[index] {
					isUsed[index] = true
					dfsMinMutation(gene, count+1)
					isUsed[index] = false
				}

			}
		}
	}

	dfsMinMutation(start, 0)

	if minCount <= len(bank) {
		return minCount
	}
	return -1

}

func isExist(end string, bank []string) int {

	for key, value := range bank {
		if value == end {
			return key
		}
	}
	return -1
}
