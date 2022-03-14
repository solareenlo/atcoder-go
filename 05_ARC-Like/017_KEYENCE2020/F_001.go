package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 12
const mod = 998244353

func sgn(x int) int {
	if x&1 != 0 {
		return 1
	}
	return mod - 1
}

func main() {
	in := bufio.NewReader(os.Stdin)

	c := [N][N]int{}
	for i := 0; i < N; i++ {
		c[i][0] = 1
		for j := 1; j <= i; j++ {
			c[i][j] = (c[i-1][j] + c[i-1][j-1]) % mod
		}
	}

	dp := [2][N][N]int{}
	for t := 0; t < 2; t++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				if i == 0 || j == 0 {
					dp[t][i][j] = 1 << ((i + j) * t)
				}
			}
		}
		for i := 1; i < N; i++ {
			for j := 1; j < N; j++ {
				for k := 1; k <= i; k++ {
					dp[t][i][j] = (dp[t][i][j] + c[i][k]*dp[t][i-k][j]%mod*sgn(k)%mod*(1<<k)) % mod
				}
				for k := 1; k <= j; k++ {
					dp[t][i][j] = (dp[t][i][j] + c[j][k]*dp[t][i][j-k]%mod*sgn(k)%mod*(1<<k)) % mod
				}
				for k := 1; k <= i; k++ {
					for l := 1; l <= j; l++ {
						dp[t][i][j] = (dp[t][i][j] + 2*c[i][k]*c[j][l]%mod*dp[t][i-k][j-l]%mod*sgn(k+l)) % mod
					}
				}
			}
		}
	}

	var n, m int
	fmt.Fscan(in, &n, &m)
	row := [N]int{}
	col := [N]int{}
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(in, &s)
		for j := 0; j < m; j++ {
			if s[j] == '#' {
				row[i] |= 1 << j
				col[j] |= 1 << i
			}
		}
	}

	ans := 0
	for i := 1; i < (1 << n); i++ {
		for j := 1; j < (1 << m); j++ {
			fl := 0
			a := 0
			b := 0
			for k := 0; k < n; k++ {
				if (i>>k)&1 != 0 {
					a++
					tmp := 0
					if (row[k]&j) == j || (^row[k]&j) == j {
						tmp = 1
					}
					fl |= tmp
				}
			}
			for k := 0; k < m; k++ {
				if (j>>k)&1 != 0 {
					b++
					tmp := 0
					if (col[k]&i) == i || (^col[k]&i) == i {
						tmp = 1
					}
					fl |= tmp
				}
			}
			if fl == 0 {
				ans = (ans + dp[1][n-a][m-b]) % mod
			}
		}
	}

	ans = (ans + dp[0][n][m]) % mod
	fmt.Println(ans)
}
