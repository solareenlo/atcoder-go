package main

import "fmt"

func main() {
	const MOD = 998244353

	var f func(string) int
	f = func(s string) int {
		n := len(s)
		dp := make([]int, n+2)
		dp[1] = 1
		var p [128]int
		for i := 0; i < n; i++ {
			x := (dp[i+1] - dp[p[s[i]]] + MOD) % MOD
			dp[i+2] = (dp[i+1] + x) % MOD
			p[s[i]] = i + 1
		}
		return dp[n+1] - 1
	}

	var s, t string
	fmt.Scan(&s, &t)
	n := len(s)
	m := len(t)

	var dp [1002][1002]int
	dp[1][1] = 1
	for i := 0; i < n; i++ {
		dp[i+2][1] = 1
	}
	for j := 0; j < m; j++ {
		dp[1][j+2] = 1
	}

	var p [128]int
	for i := 0; i < n; i++ {
		var q [128]int
		for j := 0; j < m; j++ {
			x := 0
			if s[i] == t[j] {
				x = (((dp[i+1][j+1]-dp[p[s[i]]][j+1]+MOD)%MOD-dp[i+1][q[t[j]]]+MOD)%MOD + dp[p[s[i]]][q[t[j]]]) % MOD
			}
			dp[i+2][j+2] = (((dp[i+1][j+2]+dp[i+2][j+1])%MOD-dp[i+1][j+1]+MOD)%MOD + x) % MOD
			q[t[j]] = j + 1
		}
		p[s[i]] = i + 1
	}
	fmt.Println(((f(s)+f(t))%MOD - (dp[n+1][m+1] - 1) + MOD) % MOD)
}
