package main

import "fmt"

const mod = 1e9 + 7

func main() {
	var N, K, A int
	fmt.Scan(&N, &K)
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, 256)
	}
	dp[0][0] = 1
	for i := 0; i < N; i++ {
		fmt.Scan(&A)
		for j := i; j >= 0; j-- {
			for k := 0; k < 256; k++ {
				dp[j+1][k^A] = (dp[j+1][k^A] + dp[j][k]*(int(j+1))) % mod
			}
		}
	}
	var ans int
	for j := 0; j <= N; j++ {
		ans += dp[j][K]
	}
	fmt.Println(ans % mod)
}
