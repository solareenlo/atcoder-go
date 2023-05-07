package main

import "fmt"

func main() {
	const mod = 998244353

	var N, M int
	fmt.Scan(&N, &M)

	var dp [4005][2005]int
	dp[0][0] = 1
	for n := 1; n <= 2*N; n++ {
		for m := 0; m <= M; m++ {
			if m == 0 {
				dp[n][M] = dp[n-1][0]
			} else {
				dp[n][M-m] = (dp[n][M-m] + dp[n-1][m]) % mod
				if 0 <= M-m-2 {
					dp[n][M-m-2] = (dp[n][M-m-2] + dp[n-1][m]) % mod
				}
			}
		}
		for m := M; 0 <= m; m-- {
			dp[n][m] = (dp[n][m] + dp[n][m+2]) % mod
		}
	}
	fmt.Println(dp[2*N][0])
}
