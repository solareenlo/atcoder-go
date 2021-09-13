package main

import "fmt"

func main() {
	var n, a, x int
	fmt.Scan(&n, &a)

	dp := [51][2501]int{}
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		fmt.Scan(&x)
		for j := i; j > 0; j-- {
			for k := a * n; k >= x; k-- {
				dp[j][k] += dp[j-1][k-x]
			}
		}
	}

	res := 0
	for i := 1; i <= n; i++ {
		res += dp[i][a*i]
	}
	fmt.Println(res)
}
