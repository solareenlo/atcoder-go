package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)

	a := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		fmt.Fscan(in, &a[i])
	}
	sort.Ints(a[1:])

	dp := [200001][31]int{}
	p := 1
	for i := 1; i < n+1; i++ {
		for ; a[p]*2 <= a[i]; p++ {
		}
		dp[i][0] = i
		for j := 1; j < 31; j++ {
			dp[i][j] = min(dp[i-1][j]+1, dp[p-1][j-1])
		}
	}

	for p = 0; dp[n][p] > k; p++ {
	}

	fmt.Println(p, dp[n][p])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
