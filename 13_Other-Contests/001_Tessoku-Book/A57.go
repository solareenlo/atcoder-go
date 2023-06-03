package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, q int
	fmt.Fscan(in, &n, &q)
	dp := make([][]int, 31)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &dp[0][i])
		dp[0][i]--
	}
	for i := 1; i <= 30; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = dp[i-1][dp[i-1][j]]
		}
	}

	for q > 0 {
		q--
		var x, y int
		fmt.Fscan(in, &x, &y)
		x--
		for i := 0; i < 31; i++ {
			if (y & (1 << i)) != 0 {
				x = dp[i][x]
			}
		}
		fmt.Println(x + 1)
	}
}
