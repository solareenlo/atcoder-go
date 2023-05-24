package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M, K int
	fmt.Fscan(in, &N, &M, &K)
	A := make([]int, M+1)
	for i := range A {
		A[i] = N + 1
	}
	B := make([]int, M)
	for i := 0; i < M; i++ {
		fmt.Fscan(in, &A[i], &B[i])
	}
	now := 0
	ans := 0
	for i := 0; i < M; i++ {
		now += B[i]
		ans += max(0, min(A[i+1]-A[i], now-K))
		now -= A[i+1] - A[i]
		now = max(0, now)
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
