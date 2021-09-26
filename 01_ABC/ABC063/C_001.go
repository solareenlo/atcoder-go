package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	s := make([]int, n)
	for i := range s {
		fmt.Scan(&s[i])
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 10001)
	}
	dp[0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < 10001; j++ {
			if dp[i][j] != 0 {
				dp[i+1][j] = 1
				dp[i+1][j+s[i]] = 1
			}
		}
	}

	res := 0
	for i := 0; i < 10001; i++ {
		if i%10 != 0 && dp[n][i] == 1 {
			res = max(res, i)
		}
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
