package main

import "fmt"

func main() {
	var s int
	fmt.Scan(&s)

	mod := int(1e9 + 7)
	dp := make([]int, 2001)
	dp[3] = 1
	for i := 4; i < s+1; i++ {
		dp[i] = (dp[i-1] + dp[i-3]) % mod
	}

	fmt.Println(dp[s])
}
