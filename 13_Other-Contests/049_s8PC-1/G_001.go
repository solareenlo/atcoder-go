package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MAXN = 16
	const INF = int(1e18)

	var n, m int
	fmt.Fscan(in, &n, &m)
	var dp [(1 << MAXN)][MAXN]int
	for i := 0; i < (1 << n); i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = INF
		}
	}
	dp[0][0] = 0
	var cou [(1 << MAXN)][MAXN]int
	cou[0][0] = 1
	var d [MAXN][MAXN]int
	var T [MAXN][MAXN]int
	for i := 0; i < m; i++ {
		var a, b, c, t int
		fmt.Fscan(in, &a, &b, &c, &t)
		a--
		b--
		d[a][b] = c
		d[b][a] = c
		T[a][b] = t
		T[b][a] = t
	}
	for i := 0; i < (1 << n); i++ {
		for k := 0; k < n; k++ {
			if cou[i][k] == 0 {
				continue
			}
			for j := 0; j < n; j++ {
				a := dp[i][k] + d[k][j]
				b := 1 << j
				if ((i>>j)&1) != 0 || d[k][j] == 0 || a > T[k][j] {
					continue
				}
				if dp[i|b][j] == a {
					cou[i|b][j] += cou[i][k]
				} else if dp[i|b][j] > a {
					dp[i|b][j] = a
					cou[i|b][j] = cou[i][k]
				}
			}
		}
	}
	if dp[(1<<n)-1][0] == INF {
		fmt.Println("IMPOSSIBLE")
	} else {
		fmt.Println(dp[(1<<n)-1][0], cou[(1<<n)-1][0])
	}
}
