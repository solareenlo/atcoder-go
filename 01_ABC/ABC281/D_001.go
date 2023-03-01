package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k, d int
	fmt.Fscan(in, &n, &k, &d)

	dp := make([][]int, 108)
	for i := range dp {
		dp[i] = make([]int, 108)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	dp[0][0] = 0

	for i := 1; i <= n; i++ {
		var x int
		fmt.Fscan(in, &x)
		for j := k; j >= 1; j-- {
			for k := 0; k < d; k++ {
				if dp[j-1][k] != -1 {
					dp[j][(x+k)%d] = max(dp[j][(x+k)%d], dp[j-1][k]+x)
				}
			}
		}
	}
	fmt.Println(dp[k][0])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
