package Week_02

func removeOuterParentheses(s string) string {
	k := 0
	ret := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		if s[i] == ')' {
			k -= 1
		}
		if k != 0 {
			ret = append(ret, s[i])
		}
		//第一次遇到（ 后，将k+1进入下一个循环中去，保存下一个括号。当遇到原语右侧的）时，k减去了1，不入到结果中。
		if s[i] == '(' {
			k += 1
		}
	}
	return string(ret)
}

func main()  {
	removeOuterParentheses("(()())(())(()(()))")
}