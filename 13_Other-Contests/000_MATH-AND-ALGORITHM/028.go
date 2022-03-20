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

	h := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		fmt.Fscan(in, &h[i])
	}

	dp := make([]int, n+1)
	dp[1] = 0
	dp[2] = abs(h[1] - h[2])

	for i := 3; i <= n; i++ {
		dp[i] = min(dp[i-1]+abs(h[i]-h[i-1]), dp[i-2]+abs(h[i]-h[i-2]))
	}

	fmt.Println(dp[n])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
