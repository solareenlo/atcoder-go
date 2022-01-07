package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	b := make([][]int, n)
	for i := 0; i < n; i++ {
		b[i] = make([]int, 3)
		fmt.Scan(&b[i][0], &b[i][1], &b[i][2])
	}
	for i := 0; i < n; i++ {
		sort.Ints(b[i])
	}

	maxN, maxM, maxL := 0, 0, 0
	for i := 0; i < n; i++ {
		maxN = max(maxN, b[i][0])
		maxM = max(maxM, b[i][1])
		maxL = max(maxL, b[i][2])
	}

	fmt.Println(maxN * maxM * maxL)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
