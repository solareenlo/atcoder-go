package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	type pair struct {
		x, y int
	}

	var n, C int
	fmt.Fscan(in, &n, &C)

	G := make([][]pair, 1<<17)
	for i := 0; i < n; i++ {
		var A, B, F int
		fmt.Fscan(in, &A, &B, &F)
		G[A] = append(G[A], pair{B, F})
	}

	ans := 0
	dp := make([]int, 1<<17)
	for i := int(1e5); i >= 0; i-- {
		dp[i] = max(dp[i+1]-C, dp[i])
		for _, p := range G[i] {
			dp[i] = max(dp[i], dp[p.x]+p.y-C*(p.x-i))
		}
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
