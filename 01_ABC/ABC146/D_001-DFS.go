package main

import "fmt"

type Edge struct{ to, id int }

var g [][]Edge
var res []int

func dfs(v, color, p int) {
	k := 1
	for i := 0; i < len(g[v]); i++ {
		u := g[v][i].to
		id := g[v][i].id
		if u == p {
			continue
		}
		if k == color {
			k++
		}
		res[id] = k
		k++
		dfs(u, res[id], v)
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	g = make([][]Edge, n)
	res = make([]int, n-1)
	for i := 0; i < n-1; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		a--
		b--
		g[a] = append(g[a], Edge{b, i})
		g[b] = append(g[b], Edge{a, i})
	}

	dfs(0, -1, -1)

	maxi := 0
	for i := 0; i < n; i++ {
		maxi = max(maxi, len(g[i]))
	}
	fmt.Println(maxi)

	for i := 0; i < n-1; i++ {
		fmt.Println(res[i])
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
