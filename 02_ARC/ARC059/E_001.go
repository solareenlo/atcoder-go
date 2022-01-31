package main

import "fmt"

func main() {
	var n, c int
	fmt.Scan(&n, &c)

	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
	}
	b := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&b[i])
	}

	mod := int(1e9 + 7)
	pw := [404][404]int{}
	for i := 1; i <= 400; i++ {
		pw[i][0] = 1
		for j := 1; j <= 400; j++ {
			pw[i][j] = pw[i][j-1] * i % mod
		}
	}
	for i := 1; i <= 400; i++ {
		for j := 0; j <= 400; j++ {
			pw[i][j] = (pw[i][j] + pw[i-1][j]) % mod
		}
	}

	dp := [404][404]int{}
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j <= c; j++ {
			for k := 0; k <= j; k++ {
				dp[i][j] = (dp[i][j] + dp[i-1][j-k]*(pw[b[i]][k]-pw[a[i]-1][k])%mod) % mod
			}
		}
	}

	fmt.Println((dp[n][c] + mod) % mod)
}
