package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	var dp [110][110][110]int
	for i := 1; i <= n; i++ {
		var a, b, c, w int
		fmt.Fscan(in, &a, &b, &c, &w)
		dp[a][b][c] = max(dp[a][b][c], w)
	}
	for i := 0; i < 101; i++ {
		for j := 0; j < 101; j++ {
			for z := 0; z <= 100; z++ {
				if i != 0 {
					dp[i][j][z] = max(dp[i-1][j][z], dp[i][j][z])
				}
				if j != 0 {
					dp[i][j][z] = max(dp[i][j-1][z], dp[i][j][z])
				}
				if z != 0 {
					dp[i][j][z] = max(dp[i][j][z-1], dp[i][j][z])
				}
			}
		}
	}
	for i := 1; i <= m; i++ {
		var x, y, z int
		fmt.Fscan(in, &x, &y, &z)
		fmt.Println(dp[x][y][z])
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
