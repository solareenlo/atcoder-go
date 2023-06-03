package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var a, b, dp [100009]int

	var n int
	fmt.Fscan(in, &n)
	for i := 2; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 3; i <= n; i++ {
		fmt.Fscan(in, &b[i])
	}
	dp[2] = a[2]
	for i := 3; i <= n; i++ {
		dp[i] = min(dp[i-1]+a[i], dp[i-2]+b[i])
	}
	fmt.Println(dp[n])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
