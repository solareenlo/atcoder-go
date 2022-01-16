package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m, q int
	fmt.Fscan(in, &n, &m, &q)

	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	dp := make([]int, m+1)
	for i := range dp {
		dp[i] = 1
	}

	mod := int(1e9 + 7)
	for i := 0; i < n; i++ {
		for j := m; j > a[i]; j-- {
			dp[j] += mod - dp[j-a[i]-1]
			dp[j] %= mod
		}
		if i != n-1 {
			for j := 0; j < m; j++ {
				dp[j+1] += dp[j]
				dp[j+1] %= mod
			}
		}
	}

	res := [2002][2002]int{}
	for i := 0; i < n; i++ {
		res[i][0] = 1
		for j := 0; j < m; j++ {
			res[i][j+1] = dp[j+1] - dp[j] + mod
			res[i][j+1] %= mod
		}
		for j := a[i]; j < m; j++ {
			res[i][j+1] += res[i][j-a[i]]
			res[i][j+1] %= mod
		}
	}

	for i := 0; i < q; i++ {
		var k, x int
		fmt.Fscan(in, &k, &x)
		k--
		fmt.Fprintln(out, res[k][m-x])
	}
}
