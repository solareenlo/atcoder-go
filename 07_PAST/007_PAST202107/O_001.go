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

	A := make([]int, N+1)
	B := make([]int, N+1)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i], &B[i+1])
	}

	cum := make([]int, N+2)
	for i := 0; i <= N; i++ {
		cum[i+1] = cum[i] + A[i]
	}

	dp := make([]int, N+1)
	s := 0
	m := 0
	dp[0] = A[0]
	for i := 1; i <= N; i++ {
		if B[i] <= m {
			dp[i] = dp[i-1] + A[i]
			continue
		}
		for s < i && dp[s] < B[i] {
			s++
		}
		if s == i {
			fmt.Println(-1)
			return
		}
		dp[i] = dp[s] - B[i] + cum[i+1] - cum[s+1]
		m = B[i]
	}
	fmt.Println(dp[N])
}
