package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	dp := make([]int, 100)
	for i := 1; i < 100; i++ {
		dp[i] = -1 << 60
	}

	tmp := make([]int, 100)
	for i := 0; i < N; i++ {
		for j := 0; j < 100; j++ {
			tmp[j] = dp[j]
		}
		var P, U int
		fmt.Scan(&P, &U)
		for j := 0; j < 100; j++ {
			dp[(j+P)%100] = max(dp[(j+P)%100], tmp[j]+U-P+(j+P)/100*20)
		}
	}

	ans := 0
	for i := 0; i < 100; i++ {
		ans = max(ans, dp[i])
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
