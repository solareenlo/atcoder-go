package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)

	dp := make([]int, 50)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	fmt.Println(dp[n])
}
