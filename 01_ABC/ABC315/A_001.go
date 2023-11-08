package main

import (
	"fmt"
	"regexp"
)

func main() {
	var S string
	fmt.Scan(&S)

	// 正規表現パターンを作成
	re := regexp.MustCompile("[aeiou]")

	// 文字列内の一致を削除し、置換
	result := re.ReplaceAllString(S, "")
	fmt.Println(result)
}
