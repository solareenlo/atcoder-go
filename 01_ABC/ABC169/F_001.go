package main

import "fmt"

func main() {
	var n, s int
	fmt.Scan(&n, &s)

	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	dp := make([]int, s+1)
	dp[0] = 1

	mod := 998244353
	for i := 0; i < n; i++ {
		tmp := make([]int, s+1)
		for j := 0; j < s+1; j++ {
			dp[j], tmp[j] = tmp[j], dp[j]
		}
		for j := 0; j < s+1; j++ {
			dp[j] += tmp[j] * 2
			dp[j] %= mod
			if j+a[i] <= s {
				dp[j+a[i]] += tmp[j]
				dp[j+a[i]] %= mod
			}
		}
	}

	fmt.Println(dp[s])
}
