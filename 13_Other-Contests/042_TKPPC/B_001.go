package main

import (
	"fmt"
)

func main() {
	var n, m int
	l := 0
	fmt.Scan(&n)

	ans := -int(1e18)
	for i := 1; i <= n; i++ {
		fmt.Scan(&m)
		if m > ans {
			ans = m
			l = i
		}
	}

	fmt.Println(l)
}
