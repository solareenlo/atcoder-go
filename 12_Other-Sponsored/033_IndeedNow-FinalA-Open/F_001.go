package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 1000000007

	var n, S, N, K int
	fmt.Fscan(in, &n, &S, &N, &K)
	if S+N+K < n {
		fmt.Println(0)
		return
	}
	var co [151][151]int
	for i := 0; i < 151; i++ {
		co[i][0] = 1
		co[i][i] = 1
		for j := 1; j < i; j++ {
			co[i][j] = (co[i-1][j] + co[i-1][j-1]) % MOD
		}
	}
	var dp [151][301][301]int
	for i := 0; i < n+1; i++ {
		for j := 0; j < n+1; j++ {
			for k := 0; k < n+1; k++ {
				if i+j+k <= n {
					dp[i+j+k][i][j] = co[i+j+k][i] * co[j+k][j] % MOD
				}
			}
		}
	}
	for i := 0; i < 151; i++ {
		for j := 0; j < 301; j++ {
			for k := 0; k < 301; k++ {
				if j > 0 {
					dp[i][j][k] = (dp[i][j][k] + dp[i][j-1][k]) % MOD
				}
				if k > 0 {
					dp[i][j][k] = (dp[i][j][k] + dp[i][j][k-1]) % MOD
				}
				if j > 0 && k > 0 {
					dp[i][j][k] = (dp[i][j][k] - dp[i][j-1][k-1] + MOD) % MOD
				}
			}
		}
	}
	res := 0
	for i := 0; i < S+1; i++ {
		for j := 0; j < N+1; j++ {
			if i+j <= n {
				for k := 0; k < K+1; k++ {
					if i+j+k <= n {
						for sn := 0; sn < S+N-i-j+1; sn++ {
							if i+j+k+sn <= n {
								res = (res + ((((co[n][i] * co[n-i][j] % MOD) * co[n-i-j][k] % MOD) * co[n-i-j-k][sn] % MOD) * dp[n-i-j-k-sn][N+K-j-k][K+S-k-i] % MOD)) % MOD
							}
						}
					}
				}
			}
		}
	}
	fmt.Println(res)
}
