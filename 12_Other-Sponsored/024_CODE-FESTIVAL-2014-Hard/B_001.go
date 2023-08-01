package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const B = 1000
	const MOD = 1000000007

	var n int
	fmt.Fscan(in, &n)
	var p, l [1009]int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &p[i], &l[i])
	}
	var dp [2009][2009]int
	for i := B - l[1]; i <= B+l[1]; i++ {
		dp[1][i] = 1
	}
	for i := 1; i <= 2000; i++ {
		dp[1][i] += dp[1][i-1]
	}
	for i := 2; i <= n; i++ {
		for j := B - l[i]; j <= B+l[i]; j++ {
			R := p[i] + (j - B) - 1
			R -= p[i-1]
			if R+B >= 0 {
				dp[i][j] = dp[i-1][min(2000, R+B)]
			}
		}
		for j := 1; j <= 2000; j++ {
			dp[i][j] += dp[i][j-1]
			dp[i][j] %= MOD
		}
	}
	fmt.Println(dp[n][2000])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
