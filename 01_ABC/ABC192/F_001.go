package main

import "fmt"

func main() {
	var n, s int
	fmt.Scan(&n, &s)

	a := [105]int{}
	for i := 1; i < n+1; i++ {
		fmt.Scan(&a[i])
	}

	dp := [105][105]int{}
	r := 1 << 60
	for i := 1; i < n+1; i++ {
		for j := 0; j < 105; j++ {
			for k := 0; k < 105; k++ {
				dp[j][k] = -1 << 60
			}
		}
		dp[0][0] = 0
		for j := 1; j < n+1; j++ {
			for k := i - 1; k >= 0; k-- {
				for l := 0; l < n; l++ {
					dp[k+1][(l+a[j])%i] =
						max(dp[k][l]+a[j], dp[k+1][(l+a[j])%i])
				}
			}
		}
		if dp[i][s%i] >= 0 {
			r = min(r, (s-dp[i][s%i])/i)
		}
	}

	fmt.Println(r)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
