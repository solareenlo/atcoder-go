package main

import "fmt"

func main() {
	var n, a, x int
	fmt.Scan(&n, &a)

	dp := [51][51][2551]int{}
	dp[0][0][0] = 1
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		for j := 0; j < n; j++ {
			for k := 0; k <= n*a; k++ {
				dp[i+1][j][k] += dp[i][j][k]
				dp[i+1][j+1][k+x] += dp[i][j][k]
			}
		}
	}

	res := 0
	for i := 1; i <= n; i++ {
		res += dp[n][i][i*a]
	}
	fmt.Println(res)
}
