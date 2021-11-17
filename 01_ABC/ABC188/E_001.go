package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	G := make([][]int, n)
	for i := 0; i < m; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		x--
		y--
		G[x] = append(G[x], y)
	}

	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1 << 60
	}

	maxi := -1 << 60
	for i := 0; i < n; i++ {
		for _, j := range G[i] {
			dp[j] = min(dp[j], min(dp[i], a[i]))
		}
		if dp[i] != 1<<60 {
			maxi = max(maxi, a[i]-dp[i])
		}
	}

	fmt.Println(maxi)
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
