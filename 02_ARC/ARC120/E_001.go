package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n+2)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	a[0] = a[1]
	a[n+1] = a[n]
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = a[i+1] - a[0]
		if i > 2 {
			dp[i] = min(dp[i], max(dp[i-2], a[i+1]-a[i-2]))
		}
		if i > 3 {
			dp[i] = min(dp[i], max(dp[i-3], a[i+1]-a[i-3]))
		}
	}
	fmt.Println(dp[n] / 2)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
