package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	n := len(s)
	max := strings.Repeat("a", n)
	min := strings.Repeat("z", n)

	s += s
	for i := 0; i < len(s)-n; i++ {
		sub := s[i : i+n]
		if sub > max {
			max = sub
		}
		if sub < min {
			min = sub
		}
	}
	fmt.Println(min)
	fmt.Println(max)
}
