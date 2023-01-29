package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	t := "atcoder"
	cnt := 0
	for i := 0; i < len(t); i++ {
		k := strings.Index(s, string(t[i]))
		s = s[:k] + s[k+1:]
		cnt += k
	}
	fmt.Println(cnt)
}
