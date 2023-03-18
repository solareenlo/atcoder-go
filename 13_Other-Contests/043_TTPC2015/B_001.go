package main

import (
	"fmt"
)

func main() {
	var n, b, c, s int
	fmt.Scan(&n, &b, &c)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
	}
	for i := n; i >= 1; i-- {
		s += a[i] * min(b, c)
		c -= min(b, c)
	}
	fmt.Println(s)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
