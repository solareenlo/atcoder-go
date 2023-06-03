package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m, q int
	fmt.Fscan(in, &n, &m, &q)
	var dp [2000][2000]int
	for q > 0 {
		q--
		var a, b, c, d int
		fmt.Fscan(in, &a, &b, &c, &d)
		a--
		b--
		dp[a][b]++
		dp[a][d]--
		dp[c][b]--
		dp[c][d]++
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			dp[i][j+1] += dp[i][j]
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			dp[i+1][j] += dp[i][j]
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%d ", dp[i][j])
		}
	}
}
