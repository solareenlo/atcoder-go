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

	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	dp := make([]int, n)
	dp[0] = 1000
	for i := 1; i < n; i++ {
		dp[i] = dp[i-1]
		for j := 0; j < i; j++ {
			v := dp[j] / a[j]
			w := dp[j] + (a[i]-a[j])*v
			dp[i] = max(dp[i], w)
		}
	}

	fmt.Println(dp[n-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
