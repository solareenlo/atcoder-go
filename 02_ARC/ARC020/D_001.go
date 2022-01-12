package main

import "fmt"

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)

	d := make([]int, n)
	for i := 0; i < n-1; i++ {
		fmt.Scan(&d[i])
	}

	dn := make([]int, 1<<10)
	for i := 0; i < 1<<k; i++ {
		for j := 0; j < k-1; j++ {
			dn[i] += ((i >> j) ^ (i >> (j + 1))) & 1
		}
	}

	const mod = 1000000007
	dp := [101][1 << 10][30]int{}
	dp[0][0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < 1<<k; j++ {
			for l := 0; l < m; l++ {
				dp[i+1][j][(l+d[i]*dn[j])%m] += dp[i][j][l]
				dp[i+1][j][(l+d[i]*dn[j])%m] %= mod
				for x := 0; x < k; x++ {
					if ((j >> x) & 1) == 0 {
						dp[i+1][j|(1<<x)][(l+d[i]*dn[j|(1<<x)])%m] += dp[i][j][l]
						dp[i+1][j|(1<<x)][(l+d[i]*dn[j|(1<<x)])%m] %= mod
					}
				}
			}
		}
	}

	fmt.Println(dp[n][(1<<k)-1][0])
}
