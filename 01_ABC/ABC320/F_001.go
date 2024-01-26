package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 305
	var dp [N][N][N]int
	var a, f, p [N]int

	var n, m int
	fmt.Fscan(in, &n, &m)

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			for k := 0; k < N; k++ {
				dp[i][j][k] = 1061109567
			}
		}
	}
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 1; i < n; i++ {
		fmt.Fscan(in, &f[i], &p[i])
	}
	dp[0][m][0] = 0
	for i := 1; i <= n; i++ {
		for j := 0; j <= m; j++ {
			for k := 0; k <= m; k++ {
				d := a[i] - a[i-1]
				x := j - d
				y := k + d
				if x < 0 || y > m {
					break
				}
				dp[i][x][y] = min(dp[i][x][y], dp[i-1][j][k])
				dp[i][min(m, x+p[i])][y] = min(dp[i][min(m, x+p[i])][y], dp[i-1][j][k]+f[i])
				dp[i][x][max(0, y-p[i])] = min(dp[i][x][max(0, y-p[i])], dp[i-1][j][k]+f[i])
			}
		}
	}
	ans := int(1e9)
	for i := 0; i <= m; i++ {
		for j := 0; j <= m; j++ {
			if i >= j {
				ans = min(ans, dp[n][i][j])
			}
		}
	}
	if ans != int(1e9) {
		fmt.Println(ans)
	} else {
		fmt.Println(-1)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
