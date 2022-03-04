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

	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	const mod = 998244353
	pre := make([]int, n+1)
	dp := make([]int, n+1)
	pre[1] = 1
	dp[1] = 1
	pos1, pos2 := 0, 0
	for i := 2; i <= n; i++ {
		dp[i] += dp[i-1]
		if a[i] == a[i-2] && a[i] != a[i-1] && i >= 3 {
			dp[i] = (dp[i]%mod + dp[i-2]%mod) % mod
		}
		if a[i] == a[i-1] {
			pos2 = i - 1
		}
		if a[i] != a[i-2] && i >= 3 {
			pos1 = i - 2
		}
		if pos1 > pos2 {
			dp[i] = (dp[i]%mod + (pre[pos1]%mod-pre[pos2]%mod+mod)%mod) % mod
		}
		pre[i] = (pre[i-1]%mod + dp[i]%mod) % mod
	}
	fmt.Println(dp[n])
}
