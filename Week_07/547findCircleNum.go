package Week_07

func findCircleNum(M [][]int) int {
	n := len(M)
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}

	var f func(int)int; f = func(x int) int {
		if p[x] != x {
			p[x] = f(p[x])
		}
		return p[x]
	}

	ans := n
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if M[i][j] == 1{
				pi, pj := f(i), f(j)
				if pi != pj {
					p[pj] = pi
					ans--
				}
			}
		}
	}
	return ans
}