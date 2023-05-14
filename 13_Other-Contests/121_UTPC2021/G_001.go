package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 3010

	var n, m int
	fmt.Fscan(in, &n, &m)
	var row, col [N]int
	for i := 1; i < n+1; i++ {
		var s string
		fmt.Fscan(in, &s)
		s = " " + s
		for j := 1; j < m+1; j++ {
			row[i] = row[i] + int(s[j]-'0')
			col[j] = col[j] + int(s[j]-'0')
		}
	}
	fmt.Println(solve(n, row) + solve(m, col))
}

var DP [3010][3010]int

func solve(n int, a [3010]int) int {
	for i := 1; i <= n; i++ {
		DP[i][i] = 0
		a[i] += a[i-1]
	}
	for i := n; i >= 1; i-- {
		for j := i + 1; j < n+1; j++ {
			DP[i][j] = max(DP[i+1][j], DP[i][j-1]) + a[j] - a[i-1]
		}
	}
	return DP[1][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
