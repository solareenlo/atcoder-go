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

	dp := make([]int, 5)
	for i := 1; i <= n; i++ {
		var a int
		fmt.Fscan(in, &a)
		dp[0] += a
		tmp := 2
		if a != 0 {
			tmp = a & 1
		}
		dp[1] = min(dp[0], dp[1]+tmp)
		dp[2] = min(dp[1], dp[2]+((a+1)&1))
		dp[3] = min(dp[2], dp[3]+tmp)
		dp[4] = min(dp[3], dp[4]+a)
	}
	fmt.Println(dp[4])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
