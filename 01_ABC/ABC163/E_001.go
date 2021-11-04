package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	type pair struct{ x, y int }
	a := make([]pair, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i].x)
		a[i].y = i
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i].x < a[j].x
	})

	dp := [2020][2020]int{}
	for i := n; i > 0; i-- {
		for j := i; j <= n; j++ {
			l := j - i + 1
			dp[i][j] = max(abs(a[l].y-i)*a[l].x+dp[i+1][j], abs(a[l].y-j)*a[l].x+dp[i][j-1])
		}
	}

	fmt.Println(dp[1][n])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
