package Week_02

import (
	"sync"
)

func isAnagram(s string, t string) bool {
	var sn, tn = [26]int{}, [26]int{}
	var sl, tl = len(s), len(t)
	var wg sync.WaitGroup
	if sl != tl {
		return false
	}
	wg.Add(2)
	go func() {
		createTable(&sn, s)
		wg.Done()
	}()
	go func() {
		createTable(&tn, t)
		wg.Done()
	}()
	wg.Wait()
	return sn == tn

}

func createTable(sn *[26]int, s string) {
	for i := 0; i < len(s); i++ {
		index := s[i] - 'a'
		//保证数量一致
		sn[index] ++
	}
}