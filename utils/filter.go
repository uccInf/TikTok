package utils

type TrieNode struct {
	next  map[rune]*TrieNode
	isEnd bool
}

var trie = TrieNode{isEnd: false}

func (t *TrieNode) initTrie(words []string) {
	for _, s := range words {
		trie.Add(s)
	}
}

func (t *TrieNode) Add(word string) {
	if len(word) == 0 {
		return
	}
	cur := t
	for _, r := range word {
		if _, exist := cur.next[r]; !exist {
			cur.next[r] = &TrieNode{isEnd: false, next: map[rune]*TrieNode{}}
		}
		cur = cur.next[r]
	}
	cur.isEnd = true
}

func (t *TrieNode) Filter(s string) string {
	var (
		res      string
		begin    = 0
		position = 0
		//index    = 0
		length = len(s)
		first  = true
	)

	cur := t
	for position < length {
		c := s[position]
		if cur.isEnd {
			//replace
			first = false
		} else {
			if node, exist := cur.next[rune(c)]; exist {
				if first {
					begin = position
				}
				position++
				cur = node
			} else {
				position = begin + 1
				first = false
			}
		}

	}

	return res
}
