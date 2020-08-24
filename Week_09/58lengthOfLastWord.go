package Week_09

func lengthOfLastWord(s string) int {

	var result int
	for len(s) >= 1 && string(s[len(s)-1]) == " " {
		s = s[:len(s)-1]
	}

	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) != " " {
			result++
			continue
		}
		break
	}
	return result
}
