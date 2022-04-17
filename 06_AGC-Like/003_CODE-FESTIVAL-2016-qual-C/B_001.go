package main

import (
	"fmt"
	"sort"
)

func main() {
	var k, n int
	fmt.Scan(&k, &n)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
	}
	sort.Ints(a)
	fmt.Println(max(a[n]*2-1-k, 0))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
