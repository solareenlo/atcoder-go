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

	var n, t, l int
	fmt.Fscan(in, &n, &t, &l)

	const INF = 1001001001
	g := make([][]int, n)
	g2 := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
		g2[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = INF
			g2[i][j] = INF
		}
	}
	for i := 0; i < t; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		g[a][b] = i + 1
	}

	mul := func(a, b [][]int) [][]int {
		c := make([][]int, n)
		for i := range c {
			c[i] = make([]int, n)
			for j := range c[i] {
				c[i][j] = INF
			}
		}
		for i := 0; i < n; i++ {
			for k := 0; k < n; k++ {
				for j := 0; j < n; j++ {
					c[i][j] = min(c[i][j], max(a[i][k], b[k][j]))
				}
			}
		}
		return c
	}

	for i := 0; i < n; i++ {
		g2[i][i] = 0
	}
	for l > 0 {
		if l&1 != 0 {
			g2 = mul(g2, g)
		}
		g = mul(g, g)
		l >>= 1
	}

	for i := 0; i < n; i++ {
		ans := g2[0][i]
		if ans == INF {
			ans = -1
		}
		fmt.Fprint(out, ans, " ")
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
