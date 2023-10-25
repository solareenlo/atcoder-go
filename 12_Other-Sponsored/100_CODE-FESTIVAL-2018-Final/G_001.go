package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = int(1e18)

	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([]int, n)
	sum := make([]int, n+10)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	sort.Ints(a)
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + a[i-1]
	}
	dp := make([][]int, m+10)
	for i := range dp {
		dp[i] = make([]int, n+10)
		for j := range dp[i] {
			dp[i][j] = INF
		}
	}
	dp[0][0] = 0
	for i := 1; i <= m; i++ {
		for j := i; j <= n; j++ {
			for k := 1; k <= (j+10)/i+1; k++ {
				if j-k < 0 {
					continue
				}
				dp[i][j] = min(dp[i][j], dp[i-1][j-k]+(sum[j]-sum[j-k])*k)
			}
		}
	}
	fmt.Println(dp[m][n])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
