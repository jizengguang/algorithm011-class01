package Week_04

import "sort"

func findContentChildren(g []int, s []int) int {
	//g胃口值 s饼干值

	sort.Ints(g)
	sort.Ints(s)

	var res = 0
	var gStart, sStart = 0, 0
	for gStart < len(g) && sStart < len(s) {
		if g[gStart] <= s[sStart] {
			gStart++
			sStart++
			res++
			continue
		} else {
			sStart++
		}

	}
	return res
}
