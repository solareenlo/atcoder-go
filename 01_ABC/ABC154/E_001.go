package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	n := len(s)

	var K int
	fmt.Scan(&K)

	dp := make([][2][]int, n+1)
	for i := range dp {
		dp[i][0] = make([]int, K+1)
		dp[i][1] = make([]int, K+1)
	}
	dp[0][0][0] = 1

	for i := 0; i < n; i++ {
		for j := 0; j < 2; j++ {
			r := 9
			if j == 0 {
				r = int(s[i] - '0')
			}
			for k := 0; k < K+1; k++ {
				for l := 0; l < r+1; l++ {
					c := k
					if l != 0 {
						c += 1
					}
					if c <= K {
						if l < r {
							dp[i+1][j|1][c] += dp[i][j][k]
						} else {
							dp[i+1][j|0][c] += dp[i][j][k]
						}
					}
				}
			}
		}
	}

	fmt.Println(dp[n][0][K] + dp[n][1][K])
}
