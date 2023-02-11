package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)

	dp := make([]int, 100009)
	dp[0] = 1
	sum := 0
	for i := 0; i < N; i++ {
		var t int
		fmt.Fscan(in, &t)
		sum += t
		for j := 100008; j-t >= 0; j-- {
			dp[j] = dp[j] | dp[j-t]
		}
	}

	const mod = 1_000_000_007
	ans := mod
	for i := 0; i < 100008; i++ {
		if dp[i] == 0 {
			continue
		}
		ans = min(ans, max(sum-i, i))
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
