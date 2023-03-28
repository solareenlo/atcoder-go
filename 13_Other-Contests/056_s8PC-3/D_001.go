package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var H, W int
	fmt.Scan(&H, &W)
	A := make([]int, W)
	dp := make([][]int, W)
	for i := 0; i < W; i++ {
		dp[i] = make([]int, W)
		for j := 0; j < W; j++ {
			dp[i][j] = -1e9
		}
	}
	dp[0][1] = 0
	for ii := 0; ii < H; ii++ {
		for j := 0; j < W; j++ {
			fmt.Scan(&A[j])
		}
		for i := 0; i < W; i++ {
			for j := i + 1; j < W; j++ {
				dp[i][j] += A[i]
				if i+1 < j {
					dp[i+1][j] = max(dp[i+1][j], dp[i][j])
				}
			}
		}
		for i := 0; i < W; i++ {
			for j := i + 1; j < W; j++ {
				dp[i][j] += A[j]
				if j+1 < W {
					dp[i][j+1] = max(dp[i][j+1], dp[i][j])
				}
			}
		}
	}
	ans := 0
	for i := 0; i+1 < W; i++ {
		ans = max(ans, dp[i][W-1])
	}
	fmt.Println(ans)
}
