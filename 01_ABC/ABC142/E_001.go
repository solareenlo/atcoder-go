package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	a := make([]int, m+1)
	b := make([]int, m+1)
	c := make([][20]int, m+1)
	for i := 1; i < m+1; i++ {
		fmt.Scan(&a[i], &b[i])
		for j := 1; j < b[i]+1; j++ {
			fmt.Scan(&c[i][j])
		}
	}

	dp := [5050]int{}
	for i := 0; i < 5050; i++ {
		dp[i] = 1 << 60
	}
	dp[0] = 0

	for i := 1; i <= m; i++ {
		now := 0
		for j := 1; j <= b[i]; j++ {
			now |= (1 << (c[i][j] - 1))
		}

		for bit := 0; bit < 1<<n; bit++ {
			dp[bit|now] = min(dp[bit|now], dp[bit]+a[i])
		}
	}

	if dp[(1<<n)-1] == 1<<60 {
		fmt.Println(-1)
	} else {
		fmt.Println(dp[(1<<n)-1])
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
