package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const mod = 1000000007

	var x, s, t string
	fmt.Fscan(in, &x, &s, &t)

	dp := make([]int, 1001)
	dp[0] = 1
	for i := 0; i < len(x); i++ {
		if x[i:min(len(x), i+len(s))] == s {
			dp[i+len(s)] = (dp[i+len(s)] + dp[i]) % mod
		}
		if x[i:min(len(x), i+len(t))] == t {
			dp[i+len(t)] = (dp[i+len(t)] + dp[i]) % mod
		}
	}
	fmt.Println(dp[len(x)])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
