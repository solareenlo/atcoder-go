package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, w int
	fmt.Fscan(in, &n, &w)

	dp := make([]int, w+1)
	for j := 0; j < n; j++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		for i := w; i >= 0; i-- {
			if i+a <= w {
				dp[i+a] = max(dp[i+a], dp[i]+b)
			}
		}
	}

	ans := 0
	for i := 0; i < w+1; i++ {
		ans = max(ans, dp[i])
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
