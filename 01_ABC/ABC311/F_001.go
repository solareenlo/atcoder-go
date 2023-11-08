package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 998244353
	const N = 2002

	var s [N][N]int
	var h [N]int

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 1; i <= n; i++ {
		var c string
		fmt.Fscan(in, &c)
		c = " " + c
		for j := 1; j <= m; j++ {
			if c[j] == '#' {
				h[j] = max(h[j], n-i+1)
			}
		}
	}
	for i := 0; i < n+1; i++ {
		s[0][i] = 1
	}
	for i := 1; i <= m; i++ {
		for j := h[i]; j <= n; j++ {
			dp := s[i-1][min(j+1, n)]
			if j != 0 {
				s[i][j] = (s[i][j-1] + dp) % MOD
			} else {
				s[i][j] = dp
			}
		}
	}
	fmt.Println(s[m][n])
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
