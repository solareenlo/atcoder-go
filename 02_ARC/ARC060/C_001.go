package main

import "fmt"

func main() {
	var n, a int
	fmt.Scan(&n, &a)

	x := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x[i])
	}

	dp := [51][6000]int{}
	dp[0][3000] = 1
	for i := 0; i < n; i++ {
		for j := -2500; j <= 2500; j++ {
			dp[i+1][j+3000] += dp[i][j+3000]
			dp[i+1][j+3000+x[i]-a] += dp[i][j+3000]
		}
	}
	fmt.Println(dp[n][3000] - 1)
}
