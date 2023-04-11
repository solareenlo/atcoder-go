package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)
	dp := make([]int, n+3)
	for i := range dp {
		dp[i] = 1e9
	}
	dp[0] = 0
	var h [100000]int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &h[i])
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n && j <= i+k; j++ {
			dp[j] = min(dp[j], dp[i]+abs(h[i]-h[j]))
		}
	}
	fmt.Println(dp[n-1])
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
