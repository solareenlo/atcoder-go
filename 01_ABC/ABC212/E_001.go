package main

import (
	"fmt"
)

func main() {
	nv := make([][]int, 5100)
	var n, m, k int
	fmt.Scan(&n, &m, &k)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Scan(&u, &v)
		nv[u] = append(nv[u], v)
		nv[v] = append(nv[v], u)
	}

	mod := 998244353
	dp := [2][5005]int{}
	dp[0][1] = 1
	for l := 0; l < k; l++ {
		sum := 0
		for i := 1; i < n+1; i++ {
			sum += dp[0][i]
		}
		for i := 1; i < n+1; i++ {
			dp[1][i] = sum - dp[0][i]
			for _, j := range nv[i] {
				dp[1][i] -= dp[0][j]
			}
			dp[1][i] %= mod
		}
		dp[0], dp[1] = dp[1], dp[0]
	}

	fmt.Println(dp[0][1] % mod)
}
