package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const mod = 998244353
	const N = 200200

	var n, d int
	fmt.Fscan(in, &n, &d)
	var p [N]int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &p[i])
	}
	var q [N]int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &q[i])
	}

	var dp [1005][1005]int
	dp[0][0] = 1
	var a [1005][1005]int
	var b [1005][1005]int
	for i := 1; i <= n; i++ {
		s := abs(p[i] - q[i])
		for j := 0; j <= d; j++ {
			for k := 0; k <= d; k++ {
				a[j][k] = dp[j][k]
				b[j][k] = dp[j][k]
				if j != 0 && k != 0 {
					a[j][k] = (a[j][k] + a[j-1][k-1]) % mod
				}
				if j != 0 && k < d {
					b[j][k] = (b[j][k] + b[j-1][k+1]) % mod
				}
			}
		}
		for j := 0; j <= d; j++ {
			for k := 0; k <= d; k++ {
				dp[j][k] = 0
				if j+k >= s {
					l := max(0, j-s)
					r := j + k - s - max(0, k-s)
					if l == 0 {
						dp[j][k] = (dp[j][k] + b[r][j+k-s-r]) % mod
					} else {
						dp[j][k] = (dp[j][k] + b[r][j+k-s-r] - b[l-1][j+k-s-(l-1)] + mod) % mod
					}
				}
				if j != 0 && k > s {
					dp[j][k] = (dp[j][k] + a[j-1][k-1-s]) % mod
				}
				if k != 0 && j > s {
					dp[j][k] = (dp[j][k] + a[j-1-s][k-1]) % mod
				}
			}
		}
	}
	ans := 0
	for i := 0; i <= d; i++ {
		for j := 0; j <= d; j++ {
			ans = (ans + dp[i][j]) % mod
		}
	}
	fmt.Println(ans)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
