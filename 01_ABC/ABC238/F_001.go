package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)

	p := make([]int, 305)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &p[i])
	}
	q := make([]int, 305)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &q[p[i]])
	}

	const mod = 998244353
	dp := [305][305][305]int{}
	dp[0][0][n+1] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			for k := 1; k <= n+1; k++ {
				if k > q[i] {
					dp[i][j][k] = dp[i][j][k] + dp[i-1][j-1][k]
					dp[i][j][k] %= mod
				}
				dp[i][j-1][min(q[i], k)] = dp[i][j-1][min(q[i], k)] + dp[i-1][j-1][k]
				dp[i][j-1][min(q[i], k)] %= mod
			}
		}
	}

	ans := 0
	for i := 1; i <= n+1; i++ {
		ans += dp[n][k][i]
		ans %= mod
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
