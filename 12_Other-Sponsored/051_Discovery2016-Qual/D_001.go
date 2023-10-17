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
	var dp [2444444]int
	dp[0] = 1
	s := 0
	for i := 0; i < N; i++ {
		var a int
		fmt.Fscan(in, &a)
		s += a
		for j := s; j >= a; j-- {
			dp[j] |= dp[j-a] << 1
		}
	}
	for i := 0; i < s+1; i++ {
		for j := 0; j < N+1; j++ {
			if j < N-K || (K == 1 && j != N-1) || (K >= 2 && j == N-K && j%2 != (N-K%2)%2) {
				dp[i] &= ^(1 << j)
			}
		}
	}
	for i := 0; i < M; i++ {
		var a int
		fmt.Fscan(in, &a)
		s += a
		for j := s; j >= a; j-- {
			dp[j] |= dp[j-a] << 1
			if ((dp[j] >> (N + 1)) & 1) != 0 {
				dp[j] -= 1 << (N + 1)
			}
		}
	}
	ret := 0
	for i := 0; i < s+1; i++ {
		if ((dp[i] >> N) & 1) != 0 {
			ret = max(ret, i*(s-i))
		}
	}
	fmt.Println(ret)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
