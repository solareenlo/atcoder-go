package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 606

	var n int
	fmt.Fscan(in, &n)

	dp := [N][N >> 1][N >> 1]int{}
	dp[n*2][0][0] = 1
	vst := [N]int{}
	gg := 0
	for i := 1; i <= n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		if x > -1 && y > -1 {
			vst[y] = 1
			vst[x] = 1
		} else if x > -1 {
			vst[x] = 2
		} else if y > -1 {
			vst[y] = 2
		} else {
			gg++
		}
	}

	const mod = 1_000_000_007
	c1, c2 := 0, 0
	for i := 2 * n; i > 0; i-- {
		if vst[i] == 2 {
			c1++
			for j := 0; j < c1; j++ {
				for k := 0; k <= c2; k++ {
					dp[i-1][j+1][k] += dp[i][j][k]
					dp[i-1][j+1][k] %= mod
					if k != 0 {
						dp[i-1][j][k-1] += dp[i][j][k]
						dp[i-1][j][k-1] %= mod
					}
				}
			}
		}
		if vst[i] == 0 {
			c2 = min(c2+1, n)
			for j := 0; j <= c1; j++ {
				for k := 0; k <= c2; k++ {
					dp[i-1][j][k+1] += dp[i][j][k]
					dp[i-1][j][k+1] %= mod
					if j != 0 {
						dp[i-1][j-1][k] += dp[i][j][k] * j % mod
						dp[i-1][j-1][k] %= mod
					}
					if k != 0 {
						dp[i-1][j][k-1] += dp[i][j][k]
						dp[i-1][j][k-1] %= mod
					}
				}
			}
		}
		if vst[i] == 1 {
			for j := 0; j <= n; j++ {
				for k := 0; k <= n; k++ {
					dp[i-1][j][k] = dp[i][j][k]
				}
			}
		}
	}

	ans := dp[0][0][0]
	for gg != 0 {
		ans = ans * gg % mod
		gg--
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
