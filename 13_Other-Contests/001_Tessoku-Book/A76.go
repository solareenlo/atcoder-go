package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 1000000007

	var n, w, L, R int
	fmt.Fscan(in, &n, &w, &L, &R)
	x := make([]int, n+2)
	dp := make([]int, n+3)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &x[i])
	}
	x[n+1] = w
	dp[0] = 1
	dp[1] = -1
	for i, l, r := 0, 0, 0; i < n+1; i++ {
		if dp[i] == 0 {
			continue
		}
		for l < n+2 && x[l] < x[i]+L {
			l++
		}
		r = max(r, l)
		for r < n+2 && x[r] <= x[i]+R {
			r++
		}
		dp[l] = (dp[l] + dp[i]) % MOD
		dp[r] = (dp[r] - dp[i] + MOD) % MOD
		dp[i+1] = (dp[i+1] + dp[i]) % MOD
	}
	fmt.Println(dp[n+1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
