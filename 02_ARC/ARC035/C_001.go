package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	dp := [500][500]int{}
	for i := 0; i < 500; i++ {
		for j := 0; j < 500; j++ {
			if i-j != 0 {
				dp[i][j] = 1e9
			} else {
				dp[i][j] = 0
			}
		}
	}

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 0; i < m; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		dp[a][b] = c
		dp[b][a] = c
	}

	for k := 1; k < n+1; k++ {
		for i := 1; i < n+1; i++ {
			for j := 1; j < n+1; j++ {
				dp[i][j] = min(dp[i][j], dp[i][k]+dp[k][j])
			}
		}
	}

	L := [3]int{}
	fmt.Fscan(in, &m)
	for k := 0; k < m; k++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		L[1] = a
		L[2] = b
		dp[a][b] = min(dp[a][b], c)
		dp[b][a] = min(dp[a][b], c)
		for x := 1; x < 3; x++ {
			for i := 1; i < n+1; i++ {
				for j := 1; j < n+1; j++ {
					dp[i][j] = min(dp[i][j], dp[i][L[x]]+dp[L[x]][j])
				}
			}
		}
		res := 0
		for i := 1; i < n+1; i++ {
			for j := 1; j < n+1; j++ {
				res += dp[i][j]
			}
		}
		fmt.Fprintln(out, res/2)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
