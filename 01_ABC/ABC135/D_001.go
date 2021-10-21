package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	n := len(s)
	dp := make([][13]int, n+1)
	dp[0][0] = 1

	mod := int(1e9 + 7)
	for i := 0; i < n; i++ {
		for j := 0; j < 13; j++ {
			if s[i] == '?' {
				for k := 0; k < 10; k++ {
					dp[i+1][(j*10+k)%13] += dp[i][j]
					dp[i+1][(j*10+k)%13] %= mod
				}
			} else {
				dp[i+1][(j*10+int(s[i]-'0'))%13] += dp[i][j]
				dp[i+1][(j*10+int(s[i]-'0'))%13] %= mod
			}
		}
	}
	fmt.Println(dp[n][5])
}
