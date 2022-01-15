package main

import "fmt"

func main() {
	var x, y, n int
	fmt.Scan(&x, &y, &n)

	dp := [302][602]int{}
	for i := 0; i < n; i++ {
		var t, h int
		fmt.Scan(&t, &h)
		for j := x; j >= 0; j-- {
			for k := x + y - t; k >= 0; k-- {
				dp[j+1][k+t] = max(dp[j+1][k+t], dp[j][k]+h)
			}
		}
	}

	fmt.Println(dp[x][x+y])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
