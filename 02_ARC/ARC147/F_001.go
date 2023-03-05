package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 30
const M = 63246

var n int
var DP [N + 1][2]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var T int
	fmt.Fscan(in, &T)
	pw2 := make([]int, N+1)
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, M+1)
	}
	for T > 0 {
		T--
		var X, Y, Z int
		fmt.Fscan(in, &n, &X, &Y, &Z)
		rst := X + Y + Z + 3
		pw2[0] = 1
		if rst*rst <= 2*n {
			pw2[0] = 1
			for i := 1; i <= N; i++ {
				pw2[i] = 2 * pw2[i-1] % rst
			}
			for i := 0; i <= N; i++ {
				for j := 0; j < rst; j++ {
					dp[i][j] = 0
				}
			}
			dp[0][0] = 1
			for i := 1; i <= N; i++ {
				for j := 0; j < rst; j++ {
					if dp[i-1][j] != 0 {
						//+1
						if ((n >> (i - 1)) & 1) != 0 {
							dp[i][(j+pw2[i-1])%rst] ^= 1
						}
						// 0
						dp[i][j] ^= 1
						//-1
						if (n>>(i-1))&1 != 0 {
							dp[i][(j+rst-pw2[i-1])%rst] ^= 1
						}
					}
				}
			}
			fmt.Fprintln(out, dp[N][X+1]^dp[N][Y+1]^dp[N][Z+1]^1)
		} else {
			ans := 1
			lst := 0
			if n-X-1 >= 0 {
				lst = (n-X-1)/rst*rst + X + 1
			} else {
				lst = X + 1 - rst
			}
			for i := lst; i >= -n; i -= rst {
				ans ^= F(i)
			}
			if n-Y-1 >= 0 {
				lst = (n-Y-1)/rst*rst + Y + 1
			} else {
				lst = Y + 1 - rst
			}
			for i := lst; i >= -n; i -= rst {
				ans ^= F(i)
			}
			if n-Z-1 >= 0 {
				lst = (n-Z-1)/rst*rst + Z + 1
			} else {
				lst = Z + 1 - rst
			}
			for i := lst; i >= -n; i -= rst {
				ans ^= F(i)
			}
			fmt.Fprintln(out, ans)
		}
	}
}

func F(x int) int {
	if x < 0 {
		x = -x
	}
	for i := 0; i <= N; i++ {
		for op := 0; op <= 1; op++ {
			DP[i][op] = 0
		}
	}
	DP[0][0] = 1
	for i := 1; i <= N; i++ {
		for op := 0; op <= 1; op++ {
			if DP[i-1][op] != 0 {
				//+1
				if ((n>>(i-1))&1) != 0 && ((op == 0 && (x>>(i-1))&1 != 0) || op != 0 && (x>>(i-1))&1 == 0) {
					DP[i][0] ^= 1
				}
				// 0
				if op == ((x >> (i - 1)) & 1) {
					DP[i][op] ^= 1
				}
				//-1
				if ((n>>(i-1))&1) != 0 && ((op == 0 && (x>>(i-1))&1 != 0) || (op != 0 && (x>>(i-1))&1 == 0)) {
					DP[i][1] ^= 1
				}
			}
		}
	}
	return DP[N][0]
}
