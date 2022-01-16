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

	var n int
	fmt.Fscan(in, &n)

	p := make([]int, n)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(in, &p[i+1])
	}

	m := make([]int, n)
	dp := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		m[i]++
		if i != 0 {
			m[p[i]] += m[i]
		}
		if i != 0 {
			dp[p[i]] = max(dp[p[i]], m[i])
		}
		dp[i] = max(dp[i], n-m[i])
	}

	for i := 0; i < n; i++ {
		fmt.Fprintln(out, dp[i])
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
