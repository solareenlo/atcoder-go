package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a, dp := [2][101]int{}, [2][101]int{}
	for i := 0; i < 2; i++ {
		for j := 0; j < n; j++ {
			fmt.Scan(&a[i][j])
		}
	}
	dp[0][0] = a[0][0]
	for i := 0; i < 2; i++ {
		for j := 0; j < n; j++ {
			if i != 0 {
				dp[i][j] = max(dp[i][j], dp[i-1][j]+a[i][j])
			}
			if j != 0 {
				dp[i][j] = max(dp[i][j], dp[i][j-1]+a[i][j])
			}
		}
	}
	fmt.Println(dp[1][n-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
