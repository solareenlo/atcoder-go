package main

import (
	"bufio"
	"fmt"
	"os"
)

const M = 305

var c [M][M]int
var n int

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 1000000007

	var p2 [M * M]int
	p2[0] = 1
	for i := 1; i < M*M; i++ {
		p2[i] = 2 * p2[i-1] % MOD
	}

	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &c[i][j])
		}
	}

	r := compr()
	var dp [M][M]int
	dp[r][r] = 1

	for i := 1; i <= r; i++ {
		dp[r][r] = dp[r][r] * (MOD + p2[n] - p2[i-1]) % MOD
	}

	for i := r; i <= n-1; i++ {
		for j := r; j <= n-1; j++ {
			dp[i+1][j] += dp[i][j] * p2[j-r]
			dp[i+1][j] %= MOD
			dp[i+1][j+1] += dp[i][j] * (MOD + p2[n] - p2[j])
			dp[i+1][j+1] %= MOD
		}
	}

	ans := 0
	for i := 0; i < n+1; i++ {
		ans += dp[n][i] * p2[n*(n-i)]
		ans %= MOD
	}

	fmt.Println(ans)
}

func compr() int {
	for i := 0; i < n; i++ {
		ind := -1
		m := n
		for j := i; j < n; j++ {
			for k := 0; k < n; k++ {
				if c[j][k] == 1 {
					if k < m {
						m = k
						ind = j
					}
					break
				}
			}
		}
		if ind == -1 {
			return i
		}
		for j := 0; j < n; j++ {
			c[i][j], c[ind][j] = c[ind][j], c[i][j]
		}
		for j := i + 1; j < n; j++ {
			if c[j][m] == 1 {
				for k := 0; k < n; k++ {
					c[j][k] ^= c[i][k]
				}
			}
		}
	}

	return n
}
