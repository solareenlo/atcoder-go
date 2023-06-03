package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(in, &N, &M)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
	}
	dp := make([]int, 1<<N)
	for i := range dp {
		dp[i] = int(1e9)
	}
	c := 0
	for i := 0; i < N; i++ {
		c |= (1 << i) * A[i]
	}
	dp[c] = 0
	for i := 0; i < M; i++ {
		var X, Y, Z int
		fmt.Fscan(in, &X, &Y, &Z)
		X--
		Y--
		Z--
		for j := 0; j < (1 << N); j++ {
			k := j
			k ^= 1 << X
			k ^= 1 << Y
			k ^= 1 << Z
			dp[k] = min(dp[k], dp[j]+1)
		}
	}
	ans := dp[(1<<N)-1]
	if ans == int(1e9) {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
