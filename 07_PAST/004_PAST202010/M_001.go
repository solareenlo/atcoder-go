package main

import (
	"bufio"
	"fmt"
	"os"
)

type pair struct{ x, y int }

var (
	edges  = make([][]pair, 0)
	parent = make([]pair, 0)
	depth  = make([]int, 0)
)

func dfs(now, par int) {
	for _, p := range edges[now] {
		nxt := p.x
		e := p.y
		if nxt == par {
			continue
		}
		parent[nxt] = pair{now, e}
		depth[nxt] = depth[now] + 1
		dfs(nxt, now)
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N, Q int
	fmt.Fscan(in, &N, &Q)
	edges = make([][]pair, N+1)
	for i := 1; i <= N-1; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		edges[a] = append(edges[a], pair{b, i})
		edges[b] = append(edges[b], pair{a, i})
	}
	parent = make([]pair, N+1)
	depth = make([]int, N+1)
	for i := range depth {
		depth[i] = -1
	}
	depth[1] = 1
	parent[1] = pair{0, 0}
	dfs(1, 0)

	u := make([]int, Q)
	v := make([]int, Q)
	c := make([]int, Q)
	for i := 0; i < Q; i++ {
		fmt.Fscan(in, &u[i], &v[i], &c[i])
	}

	ans := make([]int, N)
	for z := Q - 1; z >= 0; z-- {
		color := c[z]
		s := u[z]
		t := v[z]
		reparent := make([]int, 0)
		recolor := make([]int, 0)
		for s != t {
			if depth[s] < depth[t] {
				s, t = t, s
			}
			ns := parent[s].x
			e := parent[s].y
			reparent = append(reparent, s)
			recolor = append(recolor, e)
			s = ns
		}
		for len(reparent) > 0 {
			parent[reparent[0]] = pair{s, 0}
			reparent = reparent[1:]
		}
		for len(recolor) > 0 {
			ans[recolor[0]] = color
			recolor = recolor[1:]
		}
	}
	for i := 1; i <= N-1; i++ {
		fmt.Fprintln(out, ans[i])
	}
}
