package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var a, dp [202020]int

	var n, k int
	fmt.Fscan(in, &n, &k)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	dp[0] = 0
	for i := 1; i < k; i++ {
		dp[i] = dp[i-1] + a[i]
	}
	x := dp[0]
	for i := k; i <= n; i++ {
		x = max(dp[i-k], x)
		dp[i] = max(x, dp[i-1]+a[i])
	}
	fmt.Println(dp[n])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
