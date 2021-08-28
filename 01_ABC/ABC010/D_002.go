package main

import (
	"fmt"
)

func main() {
	var n, g, e, p int
	fmt.Scan(&n, &g, &e)
	mf = make([][]*Edge, n+1)
	for i := 0; i < g; i++ {
		fmt.Scan(&p)
		AddEdge(p, n, 1)
	}

	var a, b int
	for i := 0; i < e; i++ {
		fmt.Scan(&a, &b)
		AddEdge(a, b, 1)
		AddEdge(b, a, 1)
	}

	fmt.Println(MaxFlow(0, n))
}

const Inf = 1 << 60

type Edge struct {
	To, Cap, Rev int
}

var mf [][]*Edge

func MaxFlow(s, t int) int {
	res := 0
	for {
		f := dfs(s, t, Inf, make([]bool, len(mf)))
		if f == 0 {
			return res
		}
		res += f
	}
}

func AddEdge(from, to, cap int) {
	mf[from] = append(mf[from], &Edge{To: to, Cap: cap, Rev: len(mf[to])})
	mf[to] = append(mf[to], &Edge{To: from, Cap: 0, Rev: len(mf[from]) - 1})
}

func dfs(v, t, up int, used []bool) int {
	if v == t {
		return up
	}
	used[v] = true
	for _, edge := range mf[v] {
		if !used[edge.To] && edge.Cap > 0 {
			d := dfs(edge.To, t, min(up, edge.Cap), used)
			if d > 0 {
				edge.Cap -= d
				mf[edge.To][edge.Rev].Cap += d
				return d
			}
		}
	}
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
