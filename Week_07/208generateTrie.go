package Week_07

type Trie struct {
	IsEnd bool
	next  [26]*Trie
}

func Constructor() Trie {
	return Trie{}
}

func (trie *Trie) Insert(word string) {
	for _, v := range word {
		if trie.next[v-'a'] == nil {
			trie.next[v-'a'] = new(Trie)
		}
		trie = trie.next[v-'a']
	}
	trie.IsEnd = true
}

func (trie *Trie) Search(word string) bool {
	for _, v := range word {
		if trie.next[v-'a'] == nil {
			return false
		}
		trie = trie.next[v-'a']
	}
	if trie.IsEnd == false {
		return false
	}

	return true
}

func (trie *Trie) StartWith(prefix string) bool {
	for _, v := range prefix {
		if trie.next[v-'a'] == nil {
			return false
		}
		trie = trie.next[v-'a']
	}
	return true
}
