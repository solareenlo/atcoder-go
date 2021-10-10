package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	dp := make([]int, 4)
	dp[0] = 1
	mod := int(1e9 + 7)
	abc := "ABC"

	for i := range s {
		for j := 3; j >= 0; j-- {
			if s[i] == '?' {
				dp[j] *= 3
				dp[j] %= mod
			}
			if j > 0 && (s[i] == '?' || s[i] == abc[j-1]) {
				dp[j] += dp[j-1]
				dp[j] %= mod
			}
		}
	}
	fmt.Println(dp[3])
}
