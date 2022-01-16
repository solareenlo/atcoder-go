package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	t := make([]int, 4)
	for i := 0; i < n; i++ {
		fmt.Scan(&t[i])
	}
	sort.Ints(t)
	fmt.Println(min(max(t[0]+t[3], t[1]+t[2]), max(t[3], t[0]+t[1]+t[2])))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
