package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 1000000007

	var n int
	fmt.Fscan(in, &n)
	var s string
	fmt.Fscan(in, &s)
	m := len(s)
	t := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &t[i])
	}

	dp := make([]int, m+1)
	dp[0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if s[i:min(len(s), i+len(t[j]))] == t[j] {
				dp[i+len(t[j])] += dp[i]
				dp[i+len(t[j])] %= MOD
			}
		}
	}
	fmt.Println(dp[m])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
