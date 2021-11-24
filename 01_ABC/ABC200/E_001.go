package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	dp := [4][3000005]int{}
	dp[0][0] = 1
	for i := 0; i < 3; i++ {
		for j := 0; j < i*n+1; j++ {
			dp[i+1][j+1] += dp[i][j]
			dp[i+1][j+n+1] -= dp[i][j]
		}
		for j := 1; j < (i+1)*n+1; j++ {
			dp[i+1][j] += dp[i+1][j-1]
		}
	}

	var x int
	for i := 3; i <= 3*n; i++ {
		if k <= dp[3][i] {
			x = i
			break
		} else {
			k -= dp[3][i]
		}
	}

	for i := 1; i <= n; i++ {
		jmi := max(1, x-i-n)
		jma := min(n, x-i-1)
		if jmi > jma {
			continue
		}
		if k > (jma - jmi + 1) {
			k -= (jma - jmi + 1)
			continue
		}
		y := jmi + k - 1
		z := x - i - y
		fmt.Println(i, y, z)
		return
	}
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
