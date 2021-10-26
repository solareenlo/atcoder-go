package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, t int
	fmt.Scan(&n, &t)

	type pair struct{ a, b int }
	ab := make([]pair, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ab[i].a, &ab[i].b)
	}
	sort.Slice(ab, func(i, j int) bool {
		return ab[i].a < ab[j].a
	})

	dp := make([]int, 3001)
	for i := range ab {
		for j := t - 1; j >= 0; j-- {
			k := min(j+ab[i].a, t)
			dp[k] = max(dp[k], dp[j]+ab[i].b)
		}
	}

	fmt.Println(dp[t])
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
