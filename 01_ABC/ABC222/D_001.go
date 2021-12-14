package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	b := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}
	for i := range b {
		fmt.Scan(&b[i])
	}

	MAX := 3000
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, MAX+1)
	}
	dp[0][0] = 1

	mod := 998244353
	for i := 0; i < n+1; i++ {
		for j := 0; j < MAX; j++ {
			dp[i][j+1] += dp[i][j]
			dp[i][j+1] %= mod
		}
		if i != n {
			for j := a[i]; j <= b[i]; j++ {
				dp[i+1][j] += dp[i][j]
				dp[i+1][j] %= mod
			}
		}
	}

	fmt.Println(dp[n][MAX])
}
