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
	a := make([]int, n-1)
	b := make([]int, n-1)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	for i := range b {
		fmt.Fscan(in, &b[i])
	}
	dp := make([]int, n)
	for i := range dp {
		dp[i] = -1000000000
	}
	dp[0] = 0
	for i := 0; i < n-1; i++ {
		dp[a[i]-1] = max(dp[a[i]-1], dp[i]+100)
		dp[b[i]-1] = max(dp[b[i]-1], dp[i]+150)
	}
	fmt.Println(dp[n-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
