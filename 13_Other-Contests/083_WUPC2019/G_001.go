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
	table := make([][]int, N)
	for i := range table {
		table[i] = make([]int, N)
	}
	p := make([]int, N)
	for i := 0; i < M; i++ {
		var a int
		fmt.Fscan(in, &a)
		p[a]++
		for j := 0; j < N; j++ {
			if j == a {
				continue
			}
			table[j][a] += p[j]
		}
	}
	dp := make([]int, 1<<N)
	for i := range dp {
		dp[i] = 1e11
	}
	dp[0] = 0
	for i := 0; i < 1<<N; i++ {
		for j := 0; j < N; j++ {
			if ((1 << j) & i) != 0 {
				continue
			}
			tmp := dp[i]
			for k := 0; k < N; k++ {
				if k == j {
					continue
				}
				if ((1 << k) & i) != 0 {
					continue
				}
				tmp += table[k][j]
			}
			dp[i+(1<<j)] = min(tmp, dp[i+(1<<j)])
		}
	}
	fmt.Println(dp[(1<<N)-1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
