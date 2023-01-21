package main

import "fmt"

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)

	const N = 5050
	var f, dp [N]int
	for i := 1; i <= m; i++ {
		f[i] = 1
		dp[i] = i
	}

	const mod = 998244353
	for i := 2; i <= n; i++ {
		for j := 1; j <= m; j++ {
			x := f[j]
			f[j] = (((dp[max(j-k, 0)]-dp[min(j+k-1, m)])%mod+mod)%mod + dp[m]) % mod
			if k == 0 {
				f[j] = ((f[j]-x)%mod + mod) % mod
			}
		}
		for j := 1; j <= m; j++ {
			dp[j] = (dp[j-1] + f[j]) % mod
		}
	}

	fmt.Println((dp[m]%mod + mod) % mod)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
