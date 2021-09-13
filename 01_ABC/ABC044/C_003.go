package main

import "fmt"

func main() {
	var n, a, x int
	fmt.Scan(&n, &a)

	dp := [51][5050]int{}
	dp[0][n*50] = 1
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		for j := 0; j < 2*n*50; j++ {
			dp[i+1][j] += dp[i][j]
			if j+x-a > 0 {
				dp[i+1][j+x-a] += dp[i][j]
			}
		}
	}

	fmt.Println(dp[n][n*50] - 1)
}
