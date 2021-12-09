package main

import "fmt"

func main() {
	var n, x, y int
	fmt.Scan(&n, &x, &y)

	dp := make([][]int, 303)
	for i := range dp {
		dp[i] = make([]int, 303)
		for j := range dp[i] {
			dp[i][j] = 1 << 60
		}
	}
	dp[0][0] = 0

	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		for j := 302; j >= 0; j-- {
			for k := 302; k >= 0; k-- {
				dp[j][k] = min(dp[j][k], dp[max(j-a, 0)][max(k-b, 0)]+1)
			}
		}
	}

	if dp[x][y] == 1<<60 {
		fmt.Println(-1)
	} else {
		fmt.Println(dp[x][y])
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
