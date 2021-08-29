package main

import (
	"bufio"
	"fmt"
	"os"
)

type Graph [][]int

type LCA struct {
	parent [][]int
	depth  []int
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	g := make(Graph, n)
	var x, y int
	for i := 0; i < n-1; i++ {
		fmt.Fscan(in, &x, &y)
		x--
		y--
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	lca := initLCA(&g, 0)

	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x, &y)
		x--
		y--
		fmt.Println(lca.dist(x, y) + 1)
	}
}

func initLCA(g *Graph, r int) *LCA {
	lca := &LCA{}
	v := len(*g)
	h := 1
	for 1<<h < v {
		h++
	}
	lca.parent = make([][]int, h)
	for i := range lca.parent {
		lca.parent[i] = make([]int, v)
		for j := range lca.parent[i] {
			lca.parent[i][j] = -1
		}
	}
	lca.depth = make([]int, v)
	for i := range lca.depth {
		lca.depth[i] = -1
	}
	lca.dfs(g, r, -1, 0)
	for i := 0; i+1 < len(lca.parent); i++ {
		for u := 0; u < v; u++ {
			if lca.parent[i][u] != -1 {
				lca.parent[i+1][u] = lca.parent[i][lca.parent[i][u]]
			}
		}
	}
	return lca
}

func (lca *LCA) dfs(g *Graph, v, p, d int) {
	lca.parent[0][v] = p
	lca.depth[v] = d
	for _, e := range (*g)[v] {
		if e != p {
			lca.dfs(g, e, v, d+1)
		}
	}
}

func (lca *LCA) get(u, v int) int {
	if lca.depth[u] > lca.depth[v] {
		u, v = v, u
	}
	for i := 0; i < len(lca.parent); i++ {
		if (lca.depth[v]-lca.depth[u])&(1<<i) != 0 {
			v = lca.parent[i][v]
		}
	}
	if u == v {
		return u
	}
	for i := len(lca.parent) - 1; i >= 0; i-- {
		if lca.parent[i][u] != lca.parent[i][v] {
			u = lca.parent[i][u]
			v = lca.parent[i][v]
		}
	}
	return lca.parent[0][u]
}

func (lca *LCA) dist(u, v int) int {
	return lca.depth[u] + lca.depth[v] - 2*lca.depth[lca.get(u, v)]
}
