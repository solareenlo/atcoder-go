package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	const INF = 1 << 61
	dp := [1 << 10]int{}
	for i := 0; i < 1<<n; i++ {
		dp[i] = INF
	}
	dp[0] = 0
	for k := 0; k < m; k++ {
		var s string
		var cost int
		fmt.Fscan(in, &s, &cost)
		b := 0
		for i := 0; i < n; i++ {
			if s[i] == 'Y' {
				b |= 1 << i
			}
		}
		for i := 0; i < 1<<n; i++ {
			dp[i|b] = min(dp[i|b], dp[i]+cost)
		}
	}

	if dp[(1<<n)-1] < INF {
		fmt.Println(dp[(1<<n)-1])
	} else {
		fmt.Println(-1)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
