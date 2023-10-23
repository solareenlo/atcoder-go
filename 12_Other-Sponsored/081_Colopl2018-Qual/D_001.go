package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, x int
	fmt.Fscan(in, &n, &x)

	var t [5050]int
	t[0] = -x
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &t[i+1])
	}

	var dp [5050][5050]int
	l := 0
	for i := 1; i <= n; i++ {
		for l < i-1 && t[l]+x < t[i] {
			l++
		}
		for j := 0; j < n; j++ {
			if l == 0 {
				dp[i][j+1] = dp[l][j] + min(x, t[i]-t[l])
			} else {
				dp[i][j+1] = max(dp[l][j]+min(x, t[i]-t[l]), dp[l-1][j]+min(x, t[i]-t[l-1]))
			}
		}
	}
	for i := 1; i <= n; i++ {
		fmt.Fprintln(out, dp[n][i])
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
