package main

import "fmt"

func main() {
	var n, ma, mb int
	fmt.Scan(&n, &ma, &mb)

	dp := make([][]int, 410)
	for i := range dp {
		dp[i] = make([]int, 410)
		for j := range dp[i] {
			dp[i][j] = int(1e9)
		}
	}
	dp[0][0] = 0

	var a, b, c int
	for k := 0; k < n; k++ {
		fmt.Scan(&a, &b, &c)
		for i := 400; i >= a; i-- {
			for j := 400; j >= b; j-- {
				dp[i][j] = min(dp[i][j], dp[i-a][j-b]+c)
			}
		}
	}

	res := int(1e9)
	for i := 1; i*ma <= 400 && i*mb <= 400; i++ {
		res = min(res, dp[i*ma][i*mb])
	}

	if res == int(1e9) {
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
