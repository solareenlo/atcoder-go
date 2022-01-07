package main

import "fmt"

const MAX_V = 6000

var (
	V     int
	G     = make([][]int, MAX_V)
	match = [MAX_V]int{}
	used  = [MAX_V]bool{}
	g     = [1000][1000]bool{}
)

func add_edge(u, v int) {
	G[u] = append(G[u], v)
	G[v] = append(G[v], u)
}

func dfs(v int) bool {
	used[v] = true
	for i := 0; i < len(G[v]); i++ {
		u := G[v][i]
		w := match[u]
		if w < 0 || !used[w] && dfs(w) {
			match[v] = u
			match[u] = v
			return true
		}
	}
	return false
}

func bipartite_matching() int {
	res := 0
	for i := range match {
		match[i] = -1
	}
	for v := 0; v < V; v++ {
		if match[v] < 0 {
			for i := range used {
				used[i] = false
			}
			if dfs(v) {
				res++
			}
		}
	}
	return res
}

func main() {
	var n int
	fmt.Scan(&n)

	c := [100][3]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < 3; j++ {
			fmt.Scan(&c[i][j])
		}
	}
	es := make([][2]int, 0)
	idx := map[int]int{}
	for i := 0; i < n; i++ {
		d := c[i][0] * c[i][1] * c[i][2]
		for j := 0; j < 3; j++ {
			for x := 1; x < c[i][j]; x++ {
				a := d / c[i][j] * x
				b := d / c[i][j] * (c[i][j] - x)
				idx[a] = 0
				idx[b] = 0
				es = append(es, [2]int{a, b})
			}
		}
	}
	V = 0
	for i := range idx {
		idx[i] = V
		V++
	}
	for i := range es {
		g[idx[es[i][0]]][idx[es[i][1]]] = true
	}
	for i := 0; i < V; i++ {
		for j := 0; j < V; j++ {
			if g[i][j] {
				add_edge(i, j+V)
			}
		}
	}
	fmt.Println(2*V - bipartite_matching())
}
