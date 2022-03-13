package main

import "fmt"

var (
	a = [2][3]int{}
)

func solve(n, p int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		for j := 0; j < 3; j++ {
			tmp := 0
			if p == 0 {
				tmp = 1
			}
			if a[p][j] < a[tmp][j] && i-a[p][j] >= 0 {
				dp[i] = max(dp[i], a[tmp][j]-a[p][j]+dp[i-a[p][j]])
			}
		}
	}
	return dp[n] + n
}

func main() {
	var n int
	fmt.Scan(&n)

	for j := 0; j < 2; j++ {
		for i := 0; i < 3; i++ {
			fmt.Scan(&a[j][i])
		}
	}
	fmt.Println(solve(solve(n, 0), 1))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
