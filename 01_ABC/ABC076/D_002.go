package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	T, v := make([]int, n), make([]int, n)
	sum := 0
	for i := range T {
		fmt.Scan(&T[i])
		sum += T[i]
	}
	for i := range v {
		fmt.Scan(&v[i])
	}

	dp := make([]int, 2*sum+1)
	for i := 1; i < 2*sum; i++ {
		dp[i] = 1000
	}

	t := 0
	for i := 0; i < n; i++ {
		for dt := 0; dt <= 2*T[i]; dt++ {
			dp[t+dt] = min(dp[t+dt], 2*v[i])
		}
		t += 2 * T[i]
	}

	for i := 0; i < 2*sum; i++ {
		dp[i+1] = min(dp[i+1], dp[i]+1)
	}

	for i := 2*sum - 1; i >= 0; i-- {
		dp[i] = min(dp[i], dp[i+1]+1)
	}

	dist := 0
	for i := 0; i < 2*sum; i++ {
		dist += dp[i] + dp[i+1]
	}
	fmt.Println(float64(dist) / 8.0)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
