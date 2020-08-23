package Week_09

func toLowerCase(str string) string {
	var result string
	for i := 0; i < len(str)-1; i++ {
		var s = str[i]
		if s >= 'A' && s <= 'Z' {
			s += 'a' - 'A'
		}
		result += string(s)
	}
	return result
}
