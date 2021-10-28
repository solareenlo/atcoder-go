package main

import "fmt"

var (
	n, u, v int
	uu      []int
	vv      []int
	g       [][]int
)

func dfs(dist []int, x, fa int) {
	for i := 0; i < len(g[x]); i++ {
		v := g[x][i]
		if v == fa {
			continue
		}
		dist[v] = dist[x] + 1
		dfs(dist, v, x)
	}
}

func main() {
	fmt.Scan(&n, &u, &v)

	g = make([][]int, n+1)
	for i := 1; i < n; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	uu = make([]int, n+1)
	vv = make([]int, n+1)
	dfs(uu, u, 0)
	dfs(vv, v, 0)

	res := 0
	for i := 1; i < n+1; i++ {
		if uu[i] < vv[i] {
			res = max(res, vv[i])
		}
	}
	fmt.Println(res - 1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
