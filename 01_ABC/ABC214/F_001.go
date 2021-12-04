package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	n := len(s)

	dp := make([]int, n+2)
	dp[0] = 1
	mod := int(1e9 + 7)
	for i := 0; i < n; i++ {
		for j := i - 1; ; j-- {
			dp[i+2] = (dp[i+2] + dp[j+1]) % mod
			if j == -1 || s[j] == s[i] {
				break
			}
		}
	}

	res := 0
	for i := 2; i < n+2; i++ {
		res += dp[i]
		res %= mod
	}
	fmt.Println(res)
}
