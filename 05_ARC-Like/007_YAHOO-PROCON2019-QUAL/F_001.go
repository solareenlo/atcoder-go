package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var s string
	fmt.Fscan(in, &s)
	n := len(s)

	red := 0
	blue := 0
	dp := [4005][4005]int{}
	dp[0][0] = 1

	for i := 1; i <= 2*n; i++ {
		if i <= n {
			blue += int(s[i-1] - '0')
			red = i*2 - blue
		}
		for j := max(0, i-blue); j <= red; j++ {
			if j != 0 {
				dp[i][j] = (dp[i-1][j-1] + (dp[i-1][j])) % 998244353
			} else {
				dp[i][j] = dp[i-1][j] % 998244353
			}
		}
	}
	fmt.Println(dp[2*n][red])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
