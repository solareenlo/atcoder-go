package main

import "fmt"

func main() {
	var n, m, l int
	fmt.Scan(&n, &m, &l)

	g := [301][301]int{}
	for i := range g {
		for j := range g[i] {
			g[i][j] = 1 << 60
		}
	}
	for i := range g {
		g[i][i] = 0
	}

	for i := 0; i < m; i++ {
		var a, b, c int
		fmt.Scan(&a, &b, &c)
		a--
		b--
		g[a][b] = c
		g[b][a] = c
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				g[i][j] = min(g[i][j], g[i][k]+g[k][j])
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if g[i][j] <= l {
				g[i][j] = 1
			} else {
				g[i][j] = 1 << 60
			}
		}
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				g[i][j] = min(g[i][j], g[i][k]+g[k][j])
			}
		}
	}

	var q int
	fmt.Scan(&q)
	for i := 0; i < q; i++ {
		var s, t int
		fmt.Scan(&s, &t)
		s--
		t--
		if g[s][t] < 1<<60 {
			fmt.Println(g[s][t] - 1)
		} else {
			fmt.Println(-1)
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
