package main

import "fmt"

func main() {
	const MOD = 998244353
	const N = 3030

	var dp [N][N]int

	var s string
	fmt.Scan(&s)
	n := len(s)
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if s[i-1] != ')' {
				dp[i][j+1] = (dp[i][j+1] + dp[i-1][j]) % MOD
			}
			if s[i-1] != '(' && j != 0 {
				dp[i][j-1] = (dp[i][j-1] + dp[i-1][j]) % MOD
			}
		}
	}
	fmt.Println(dp[n][0])
}
