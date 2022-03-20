package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	dp := make([]int, 1000002)
	dp[0] = 1
	sum := 1
	const mod = 1_000_000_007
	for i := 1; i < n+2; i++ {
		if i == 1 || i == n {
			dp[i] = 0
		} else {
			dp[i] = sum
		}
		sum += dp[i]
		sum %= mod
		if i-k >= 0 {
			sum += mod - dp[i-k]
			sum %= mod
		}
	}
	fmt.Println(dp[n+1])
}
