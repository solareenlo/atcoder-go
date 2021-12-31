package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	dp := make([][]int, 4)
	for i := range dp {
		dp[i] = make([]int, 1<<5)
	}
	dp[0][0] = 1e9
	for l := 0; l < n; l++ {
		dp2 := make([][]int, 4)
		for i := range dp {
			dp2[i] = make([]int, 1<<5)
		}
		for i := range dp {
			copy(dp2[i], dp[i])
		}
		for i := 0; i < 5; i++ {
			var a int
			fmt.Scan(&a)
			for j := 0; j < 3; j++ {
				for k := 0; k < 1<<5; k++ {
					if k&(1<<i) != 0 {
						dp2[j][k] = max(dp2[j][k], min(dp2[j][k^(1<<i)], a))
					}
				}
			}
		}
		for j := 0; j < 3; j++ {
			for k := 0; k < 1<<5; k++ {
				dp[j+1][k] = max(dp[j+1][k], dp2[j][k])
			}
		}
	}

	fmt.Println(dp[3][(1<<5)-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
