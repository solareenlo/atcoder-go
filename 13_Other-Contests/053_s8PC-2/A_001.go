package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	ans := 0
	v := 'O'
	for i := 0; i < len(s); i++ {
		if rune(v) != rune(s[i]) {
			ans++
			v = rune(s[i])
		}
	}
	if v == 'O' {
		ans--
	}
	fmt.Println(max(ans, 0))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
