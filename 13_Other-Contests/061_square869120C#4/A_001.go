package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	s := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}

	var t string
	fmt.Scan(&t)

	l, r := 1, n+1
	for i := 0; i < n; i++ {
		a := strings.Replace(s[i], "?", "a", -1)
		z := strings.Replace(s[i], "?", "z", -1)
		if a > t {
			r--
		}
		if z < t {
			l++
		}
	}

	for i := l; i <= r; i++ {
		fmt.Print(i)
		if i == r {
			fmt.Println()
		} else {
			fmt.Print(" ")
		}
	}
}
