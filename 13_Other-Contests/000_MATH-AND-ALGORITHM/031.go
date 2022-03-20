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
	dp[0] = a[0]
	dp[1] = max(a[1], dp[0])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-2]+a[i], dp[i-1])
	}
	fmt.Println(dp[n-1])

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
