package utils

import (
	"TikTok/constdef"
	"fmt"
	"unicode/utf8"
)

var invalidWords = []string{"傻逼", "sb", "赌博", "弱智"}
var trie *Trie

type Trie struct {
	child map[rune]*Trie
	word  string
}

func NewTrie() *Trie {
	return &Trie{
		child: make(map[rune]*Trie),
		word:  "",
	}
}

func (trie *Trie) insert(word string) *Trie {
	cur := trie
	for _, v := range []rune(word) {
		// 若存在，不做处理，若不存在，创建新的子树
		if _, ok := cur.child[v]; !ok {
			t := NewTrie()
			cur.child[v] = t
		}
		cur = cur.child[v]
	}
	cur.word = word
	return trie
}

func (trie *Trie) FilterString(word string) string {
	cur := trie

	for i, v := range []rune(word) {
		if _, ok := cur.child[v]; ok {
			cur = cur.child[v]
			if cur.word != "" {
				word = replaceStr(word, constdef.Replace, i+1-utf8.RuneCountInString(cur.word), i)
				cur = trie // ，符合条件，从头开始准备下一次遍历
			}
		} else {
			cur = trie // 不存在，则从头遍历
		}
	}
	return word
}

func replaceStr(word, replace string, left, right int) string {
	str := ""
	for i, v := range []rune(word) {
		if i >= left && i <= right {
			str = str + replace
		} else {
			str += string(v)
		}
	}
	return str
}

func init() {
	trie = NewTrie()
	for i := 0; i < len(invalidWords); i++ {
		trie.insert(invalidWords[i])
	}
}

func GetTrie() *Trie {
	return trie
}

func testFilter() {

	fmt.Println(GetTrie().FilterString("傻逼, sb, 我是你爹"))
}
