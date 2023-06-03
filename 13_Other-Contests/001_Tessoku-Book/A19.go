package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, W int
	fmt.Fscan(in, &n, &W)
	dp := make([]int, W+1)
	for n > 0 {
		n--
		var w, v int
		fmt.Fscan(in, &w, &v)
		for i := W; i >= w; i-- {
			dp[i] = max(dp[i], dp[i-w]+v)
		}
	}
	fmt.Println(dp[W])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
