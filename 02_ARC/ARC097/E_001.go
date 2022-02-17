package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	const N = 2002
	w := make([]int, N)
	b := make([]int, N)
	dw := [2 * N][N]int{}
	db := [2 * N][N]int{}
	for i := 1; i <= 2*n; i++ {
		var x int
		var c string
		fmt.Fscan(in, &c, &x)
		if c == "W" {
			w[x] = i
		} else {
			b[x] = i
		}
		for j := 0; j < n; j++ {
			dw[i][j] = dw[i-1][j]
			db[i][j] = db[i-1][j]
			if c == "W" {
				if x > j {
					dw[i][j]++
				}
			} else {
				if x > j {
					db[i][j]++
				}
			}
		}
	}

	dp := [N][N]int{}
	for j := 1; j <= n; j++ {
		dp[0][j] = 1 << 60
	}

	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			if i != 0 {
				dp[i][j] = dp[i-1][j] + dw[w[i]][i] + db[w[i]][j]
			}
			if j != 0 {
				dp[i][j] = min(dp[i][j], dp[i][j-1]+dw[b[j]][i]+db[b[j]][j])
			}
		}
	}
	fmt.Println(dp[n][n])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
