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

	g := make([][]int, n)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	var k int
	fmt.Fscan(in, &k)
	c := make([]int, k)
	for i := range c {
		fmt.Fscan(in, &c[i])
		c[i]--
	}

	dist := make([][]int, k)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = -1
		}
	}

	q := make([]int, 0)
	q = append(q, c[0])
	for i := 0; i < k; i++ {
		dist[i][c[i]] = 0
		q = append(q, c[i])
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			for _, nv := range g[v] {
				if dist[i][nv] != -1 {
					continue
				}
				dist[i][nv] = dist[i][v] + 1
				q = append(q, nv)
			}
		}
	}

	dp := make([][]int, 1<<k)
	for i := range dp {
		dp[i] = make([]int, k)
		for j := range dp[i] {
			dp[i][j] = int(1e9)
		}
	}
	for i := 0; i < k; i++ {
		dp[0][i] = 0
	}

	for bit := 1; bit < (1 << k); bit++ {
		for v := 0; v < k; v++ {
			if bit&(1<<v) != 0 {
				for p := 0; p < k; p++ {
					if dist[p][c[v]] != -1 {
						dp[bit][v] = min(dp[bit][v], dp[bit-(1<<v)][p]+dist[p][c[v]])
					}
				}
			}
		}
	}

	res := int(1e9)
	for i := 0; i < k; i++ {
		res = min(res, dp[(1<<k)-1][i])
	}
	if res == int(1e9) {
		fmt.Println(-1)
	} else {
		fmt.Println(res + 1)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
