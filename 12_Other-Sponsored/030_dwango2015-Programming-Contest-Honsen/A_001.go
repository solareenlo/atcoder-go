package main

import "fmt"

func main() {
	const MOD = 1000000007
	var N, K int
	var S string
	fmt.Scan(&N, &K, &S)
	var dp [255][255][255]int
	dp[0][0][0] = 1
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			for k := 0; k < K+1; k++ {
				if dp[i][j][k] == 0 {
					continue
				}
				for l := 0; l < 10; l++ {
					if S[i] != '?' && (int('0')+l) != int(S[i]) {
						continue
					}
					if j%2 == 0 {
						if l == 2 {
							dp[i+1][j+1][k] += dp[i][j][k]
							dp[i+1][j+1][k] %= MOD
						} else {
							dp[i+1][0][k] += dp[i][j][k]
							dp[i+1][0][k] %= MOD
						}
					} else {
						if l == 5 {
							dp[i+1][j+1][min(K, k+(j+1)/2)] += dp[i][j][k]
							dp[i+1][j+1][min(K, k+(j+1)/2)] %= MOD
						} else if l == 2 {
							dp[i+1][1][k] += dp[i][j][k]
							dp[i+1][1][k] %= MOD
						} else {
							dp[i+1][0][k] += dp[i][j][k]
							dp[i+1][0][k] %= MOD
						}
					}
				}
			}
		}
	}
	sum := 0
	for i := 0; i < N+1; i++ {
		sum += dp[N][i][K]
	}
	fmt.Println(sum % MOD)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
