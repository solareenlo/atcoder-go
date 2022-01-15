package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, l int
	fmt.Fscan(in, &n, &l)

	h := make([][][2]int, 100001)
	for i := 0; i < n; i++ {
		var l, r, c int
		fmt.Fscan(in, &l, &r, &c)
		h[r] = append(h[r], [2]int{l, c})
	}

	dp := make([]int, l+1)
	for i := 1; i <= l; i++ {
		dp[i] = 1 << 60
		for j := 0; j < len(h[i]); j++ {
			dp[i] = min(dp[i], dp[h[i][j][0]]+h[i][j][1])
		}
		for R := i - 1; dp[R] > dp[i]; {
			dp[R] = dp[i]
			R--
		}
	}
	fmt.Println(dp[l])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
