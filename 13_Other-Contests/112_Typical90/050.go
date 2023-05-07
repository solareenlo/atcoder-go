package main

import "fmt"

func main() {

	const MOD = 1000000007

	var n, l int
	fmt.Scan(&n, &l)

	var dp [100005]int
	dp[0] = 1
	for i := 0; i < n+1; i++ {
		if i+1 <= n {
			dp[i+1] += dp[i]
			dp[i+1] %= MOD
		}
		if i+l <= n {
			dp[i+l] += dp[i]
			dp[i+l] %= MOD
		}
	}
	fmt.Println(dp[n])
}
