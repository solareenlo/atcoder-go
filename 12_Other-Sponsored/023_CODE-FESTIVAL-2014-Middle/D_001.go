package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 1000000007

	var n, pp, pl int
	fmt.Fscan(in, &n, &pp, &pl)
	dp := make([]int, 2*pl+2)
	for i := range dp {
		dp[i] = 1
	}
	for i := 0; i < n-1; i++ {
		var p, l int
		fmt.Fscan(in, &p, &l)
		dp2 := make([]int, 2*l+2)
		for j := 0; j < 2*pl+1; j++ {
			x := min(max(pp-pl+j+1, p-l), p+l+1) - p + l
			dp2[x] += dp[j]
		}
		dp2[0] %= MOD
		for i := 0; i < 2*l+1; i++ {
			dp2[i+1] = (dp2[i+1] + dp2[i]) % MOD
		}
		pp = p
		pl = l
		dp = dp2
	}
	ans := 0
	for i := 0; i < 2*pl+1; i++ {
		ans += dp[i]
	}
	fmt.Println(ans % MOD)
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
