package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	f := make([]int, 100_001)
	for i := 0; i < n; i++ {
		fmt.Scan(&f[i])
		f[i]--
	}

	sum, l, r, mod := 0, 0, 0, int(1e9+7)
	used := make([]bool, m)
	dp := make([]int, n+1)
	dp[0] = 1
	for r < n {
		for l < n && used[f[r]] {
			used[f[l]] = false
			sum = (sum - dp[l] + mod) % mod
			l++
		}
		used[f[r]] = true
		sum = (sum + dp[r]) % mod
		r++
		dp[r] = sum
	}
	fmt.Println(dp[n])
}
