package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	f := make([]int, 100_002)
	for i := 0; i < n; i++ {
		fmt.Scan(&f[i+2])
		f[i+2]--
	}

	sum := make([]int, 100_002)
	dp := make([]int, 100_002)
	dp[1] = 1
	sum[1] = 1

	mod := int(1e9 + 7)
	s := 0
	c := make([]int, 100_002)
	for i := 2; i < n+2; i++ {
		if s < c[f[i]] {
			s = c[f[i]]
		}
		dp[i] = (sum[i-1] + mod - sum[s]) % mod
		sum[i] = (sum[i-1] + dp[i]) % mod
		c[f[i]] = i - 1
	}

	fmt.Println(dp[n+1])
}
