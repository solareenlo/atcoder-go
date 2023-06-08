package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = int(1e18)

	var n, m int
	fmt.Fscan(in, &n, &m)
	P := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &P[i])
	}
	cnt := make([]bool, 1<<n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if P[j][i] != P[k][i] && P[j][i] != '*' && P[k][i] != '*' {
					cnt[(1<<j)|(1<<k)] = true
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < (1 << n); j++ {
			if ((j >> i) & 1) != 0 {
				if !cnt[j] && !cnt[j-(1<<i)] {
					cnt[j] = false
				} else {
					cnt[j] = true
				}
			}
		}
	}
	dp := make([]int, 1<<n)
	for i := range dp {
		dp[i] = INF
	}
	dp[0] = 0
	for i := 1; i < (1 << n); i++ {
		for j := i; j > 0; j = (j - 1) & i {
			if !cnt[j] {
				dp[i] = min(dp[i], dp[i-j]+1)
			}
		}
	}
	fmt.Println(dp[(1<<n)-1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
