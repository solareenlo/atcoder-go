package main

import "fmt"

func main() {
	const N = 410
	const MOD = 998244353

	can := [N][N]bool{}
	dp := [N][(N / 2)]int{}
	binom := [N][N]int{}
	binom[0][0] = 1

	for i := 0; i < N; i++ {
		binom[i][0] = 1
		binom[i][i] = 1
		for j := 0; j < i-1; j++ {
			binom[i][j+1] = (binom[i-1][j] + binom[i-1][j+1]) % MOD
		}
	}

	var n, m int
	fmt.Scan(&n, &m)
	for i := 0; i < 2*n; i++ {
		for j := 0; j < 2*n; j++ {
			can[i][j] = false
		}
	}

	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		can[a-1][b-1] = true
		can[b-1][a-1] = true
	}

	for i := 0; i < 2*n+1; i++ {
		dp[i][0] = 1
	}

	for j := 1; j < n+1; j++ {
		for i := 0; i < 2*(n-j)+1; i++ {
			dp[i][j] = 0
			for k := 0; k < j; k++ {
				if can[i][i+(2*k)+1] {
					x := (dp[i+1][k] * dp[i+(2*k)+2][j-k-1]) % MOD
					dp[i][j] = ((x * binom[j][k+1]) + dp[i][j]) % MOD
				}
			}
		}
	}

	fmt.Println(dp[0][n])
}
