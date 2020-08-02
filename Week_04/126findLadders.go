package Week_04

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	dict := make(map[string]bool, 0)
	for _, v := range wordList {
		dict[v] = true
	}

	if _, ok := dict[endWord]; !ok {
		return nil
	}



}
