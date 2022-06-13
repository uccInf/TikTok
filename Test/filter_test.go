package main

import (
	"TikTok/utils"
	"fmt"
	"testing"
)

func TestFilter(t *testing.T) {
	var trie = utils.GetTrie()
	fmt.Println(trie.FilterString("傻逼，sb, sdf, 尼玛 ,哈哈哈哈哈"))
}
