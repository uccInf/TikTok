package main

import (
	"TikTok/utils"
	"fmt"
)

func main() {
	var t = utils.GetTrie()
	fmt.Println(t.FilterString("傻逼，sb, sdf, 尼玛 ,哈哈哈哈哈"))
}
