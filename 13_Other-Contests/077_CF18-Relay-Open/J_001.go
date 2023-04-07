package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = int(1e18)

	var n int
	fmt.Fscan(in, &n)
	var w, s, ss [2][3001]int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &w[0][i], &w[1][i])
		s[0][i] = s[0][i-1] + w[0][i]
		s[1][i] = s[1][i-1] + w[1][i]
		ss[0][i] = ss[0][i-1] + s[0][i]
		ss[1][i] = ss[1][i-1] + s[1][i]
	}
	var dp [2][3001]int
	for p := 0; p < 2; p++ {
		for i := range dp[p] {
			dp[p][i] = INF
		}
		for i := 1; i < n; i++ {
			dp[p][i] = ss[p^1][i]
		}
	}
	for i := 2; i <= n; i++ {
		for p := 0; p < 2; p++ {
			for j := 1; j < i; j++ {
				m := (i + j + 1) / 2
				dp[p][i] = min(dp[p][i], dp[p^1][j]+(m-j)*s[p^1][m]-ss[p^1][m-1]+ss[p^1][j-1]-(i-m)*s[p^1][m]+ss[p^1][i]-ss[p^1][m])
			}
		}
	}
	ans := INF
	for i := 1; i < n; i++ {
		for p := 0; p < 2; p++ {
			ans = min(ans, dp[p][i]+(n-i)*s[p][n]-ss[p][n-1]+ss[p][i-1])
		}
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
