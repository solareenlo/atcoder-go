package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, p int
	fmt.Scan(&n, &p)

	v := make([]pair, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&v[i].x, &v[i].y)
	}
	sort.Slice(v, func(i, j int) bool {
		return v[i].x < v[j].x
	})
	v = reverseOrder(v)

	dp := [5005][5005]int{}
	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, dp[i][p]+v[i].y)
		for j := 0; j < p+1; j++ {
			dp[i+1][j] = dp[i][j]
		}
		for j := v[i].x; j < p+1; j++ {
			dp[i+1][j] = max(dp[i+1][j], dp[i][j-v[i].x]+v[i].y)
		}
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type pair struct{ x, y int }

func reverseOrder(a []pair) []pair {
	n := len(a)
	res := make([]pair, n)
	n = copy(res, a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
