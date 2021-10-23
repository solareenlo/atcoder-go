package main

import (
	"fmt"
	"strings"
)

func main() {
	var s, t string
	fmt.Scan(&s, &t)

	ss := s + s
	res := -1
	n := len(t)
	for i := 0; i < n; i++ {
		start := (res + 1) % len(s)
		sub := ss[start:]
		index := strings.IndexByte(sub, t[i])
		if index == -1 {
			fmt.Println(-1)
			return
		}
		res += index + 1
	}
	fmt.Println(res + 1)
}
