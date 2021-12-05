package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)

	mod := 998244353
	dp := [1024][1024][10]int{}
	for i := 1; i < n+1; i++ {
		x := int(s[i-1] - 'A')
		for u := 0; u < 1024; u++ {
			for j := 0; j < 10; j++ {
				dp[i][u][j] = dp[i-1][u][j]
				if j == x {
					dp[i][u][j] += dp[i-1][u][j]
					dp[i][u][j] %= mod
				}
			}
		}
		for v := 0; v < 1024; v++ {
			if v&(1<<x) != 0 {
				continue
			}
			for j := 0; j < 10; j++ {
				dp[i][v|(1<<x)][x] += dp[i-1][v][j]
				dp[i][v|(1<<x)][x] %= mod
			}
		}
		dp[i][(1 << x)][x]++
		dp[i][(1 << x)][x] %= mod
	}

	res := 0
	for u := 0; u < 1024; u++ {
		for j := 0; j < 10; j++ {
			res += dp[n][u][j]
			res %= mod
		}
	}
	fmt.Println(res)
}
