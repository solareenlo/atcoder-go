package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD = 998244353

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)

	l := make([]int, k)
	r := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Fscan(in, &l[i], &r[i])
	}

	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = -1
	for i := 0; i < n; i++ {
		dp[i] = ((dp[i] % MOD) + MOD) % MOD
		dp[i+1] = (dp[i+1] + dp[i]) % MOD
		for j := 0; j < k; j++ {
			if i+l[j] < n {
				dp[i+l[j]] += dp[i]
			}
			if i+r[j]+1 < n {
				dp[i+r[j]+1] -= dp[i]
			}
		}
	}
	fmt.Println(dp[n-1])
}
