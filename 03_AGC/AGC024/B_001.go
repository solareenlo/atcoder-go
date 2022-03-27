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

	dp := [200005]int{}
	maxi := 0
	for i := 0; i < n; i++ {
		var p int
		fmt.Fscan(in, &p)
		dp[p] = dp[p-1] + 1
		maxi = max(maxi, dp[p])
	}
	fmt.Println(n - maxi)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
