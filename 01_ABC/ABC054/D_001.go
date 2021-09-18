package main

import "fmt"

func main() {
	var n, ma, mb int
	fmt.Scan(&n, &ma, &mb)
	a := make([]int, n)
	b := make([]int, n)
	c := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i], &b[i], &c[i])
	}

	const inf int = int(1e9 + 7)
	dp := make([][][]int, 41)
	for i := range dp {
		dp[i] = make([][]int, 40*10+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, 40*10+1)
			for k := range dp[i][j] {
				dp[i][j][k] = inf
			}
		}
	}
	dp[0][0][0] = 0

	for i := 0; i < n; i++ {
		for ca := 0; ca <= 40*10; ca++ {
			for cb := 0; cb <= 40*10; cb++ {
				if dp[i][ca][cb] == inf {
					continue
				}
				dp[i+1][ca][cb] = min(dp[i+1][ca][cb], dp[i][ca][cb])
				dp[i+1][ca+a[i]][cb+b[i]] = min(dp[i+1][ca+a[i]][cb+b[i]], dp[i][ca][cb]+c[i])
			}
		}
	}

	res := inf
	for ca := 1; ca <= 40*10; ca++ {
		for cb := 1; cb <= 40*10; cb++ {
			if ca*mb == cb*ma {
				res = min(res, dp[n][ca][cb])
			}
		}
	}

	if res == inf {
		res = -1
	}
	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
