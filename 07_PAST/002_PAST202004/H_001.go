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

	type tuple struct{ x, y, v int }
	dp := make([][]tuple, 11)
	for i := 0; i < n; i++ {
		var c string
		fmt.Fscan(in, &c)
		for j := 0; j < m; j++ {
			tmp := int(c[j] - '0')
			if c[j] == 'S' {
				tmp = 0
			} else if c[j] == 'G' {
				tmp = 10
			}
			dp[tmp] = append(dp[tmp], tuple{i, j, 1 << 60})
		}
	}

	dp[0][0].v = 0
	for i := 1; i <= 10; i++ {
		if len(dp[i]) == 0 {
			fmt.Println(-1)
			return
		}
		for j := range dp[i] {
			for k := range dp[i-1] {
				px := dp[i-1][k].x
				py := dp[i-1][k].y
				pv := dp[i-1][k].v
				dp[i][j].v = min(dp[i][j].v, abs(dp[i][j].x-px)+abs(dp[i][j].y-py)+pv)
			}
		}
	}
	fmt.Println(dp[10][0].v)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
