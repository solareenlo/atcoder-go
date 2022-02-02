package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)

	cnt := 0
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			cnt++
		}
	}
	fmt.Println(cnt)
}
