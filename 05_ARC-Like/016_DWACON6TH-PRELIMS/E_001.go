package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const mod = 1_000_000_007

func upd(x, y int) int { return (x + y) % mod }

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([]int, n+1)
	a[0] = 1 << 60
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	const N = 120
	const M = 520
	dp := [N][N][M]int{}
	dp[1][1][a[1]] = 1
	for i := 2; i <= n; i++ {
		for j := 0; j <= i; j++ {
			for k := 0; k <= m; k++ {
				if dp[i-1][j][k] != 0 {
					dp[i][j][k] = upd(dp[i][j][k], dp[i-1][j][k]*(k-(a[i]-1)*j))
					if k+a[i] <= m {
						dp[i][j+1][k+a[i]] = upd(dp[i][j+1][k+a[i]], dp[i-1][j][k]*(j+1))
					}
					for l := 1; l < a[i] && k+l <= m; l++ {
						dp[i][j][k+l] = upd(dp[i][j][k+l], dp[i-1][j][k]*2*j)
					}
					if j > 1 {
						for l := 0; l < a[i] && k+l <= m; l++ {
							dp[i][j-1][k+l] = upd(dp[i][j-1][k+l], dp[i-1][j][k]*(j-1)*(a[i]-l-1))
						}
					}
				}
			}
		}
	}

	ans := 0
	for i := 1; i <= n; i++ {
		ans = upd(ans, dp[n][i][m])
	}
	fmt.Print(ans)
}
