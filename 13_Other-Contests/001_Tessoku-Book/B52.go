package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, x int
	var t string
	fmt.Scan(&n, &x, &t)
	s := strings.Split(t, "")
	x--
	s[x] = "@"
	for j := x - 1; j >= 0 && s[j] == "."; j-- {
		s[j] = "@"
	}
	for j := x + 1; j < n && s[j] == "."; j++ {
		s[j] = "@"
	}
	fmt.Println(strings.Join(s, ""))
}
