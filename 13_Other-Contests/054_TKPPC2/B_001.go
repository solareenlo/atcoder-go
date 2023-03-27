package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	v := make([]int, n+1)
	w := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&v[i], &w[i])
	}

	dp := make([]int, m+1)
	for i := 1; i <= n; i++ {
		for j := m; j >= w[i]; j-- {
			dp[j] = max(dp[j], dp[j-w[i]]+v[i])
		}
	}

	fmt.Println(dp[m])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
