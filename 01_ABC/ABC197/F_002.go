package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	g := make([][][]int, 1010)
	for i := range g {
		g[i] = make([][]int, 26)
	}
	adj := [1010][1010]bool{}
	for i := 0; i < m; i++ {
		var a, b int
		var c string
		fmt.Scan(&a, &b, &c)
		a--
		b--
		g[a][c[0]-'a'] = append(g[a][c[0]-'a'], b)
		g[b][c[0]-'a'] = append(g[b][c[0]-'a'], a)
		adj[a][b] = true
		adj[b][a] = true
	}

	dist := [1010][1010]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dist[i][j] = -1
		}
	}
	dist[0][n-1] = 0

	res := -1
	type pair struct{ u, v int }
	q := make([]pair, 0)
	q = append(q, pair{0, n - 1})
	for len(q) > 0 {
		u := q[0].u
		v := q[0].v
		q = q[1:]
		if res != -1 && res <= dist[u][v]*2 {
			break
		}
		if u == v {
			res = dist[u][v] * 2
			break
		}
		if adj[u][v] {
			res = dist[u][v]*2 + 1
		}
		for k := 0; k < 26; k++ {
			for _, i := range g[u][k] {
				for _, j := range g[v][k] {
					if dist[i][j] == -1 {
						dist[i][j] = dist[u][v] + 1
						q = append(q, pair{i, j})
					}
				}
			}
		}
	}

	fmt.Println(res)
}
