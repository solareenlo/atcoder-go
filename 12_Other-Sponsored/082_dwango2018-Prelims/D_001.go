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
	var x [20]int
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &x[N-1-i])
	}
	var G [20]int
	for i := 1; i < N; i++ {
		var a int
		fmt.Fscan(in, &a)
		G[N-a] |= 1 << (N - 1 - i)
	}
	var sum [1 << 20]int
	for i := 0; i < 1<<N; i++ {
		for j := 0; j < N; j++ {
			if ((i >> j) & 1) != 0 {
				sum[i] += x[j]
			}
		}
	}
	var dp [1 << 20]int
	for i := 1; i < 1<<N; i++ {
		dp[i] = int(1e9)
	}
	for i := 0; i < 1<<N; i++ {
		for j := 0; j < N; j++ {
			if ((i >> j) & 1) != 0 {
				continue
			}
			if ^i&G[j] != 0 {
				continue
			}
			next := i & ^G[j]
			next |= 1 << j
			dp[next] = min(dp[next], max(dp[i], sum[i|1<<j]))
		}
	}
	fmt.Println(dp[1<<(N-1)])
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
