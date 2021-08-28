package main

import "fmt"

func main() {
	var n, g, e int
	fmt.Scan(&n, &g, &e)

	p := make([]int, n)
	for i := 0; i < g; i++ {
		fmt.Scan(&p[i])
	}

	graph := newMaxFlow(n + 1)
	var a, b int
	for i := 0; i < e; i++ {
		fmt.Scan(&a, &b)
		graph.AddEdge(b, a, 1)
		graph.AddEdge(a, b, 1)
	}
	for i := 0; i < g; i++ {
		graph.AddEdge(p[i], n, 1)
		graph.AddEdge(n, p[i], 1)
	}

	fmt.Println(graph.Flow(0, n))
}

type _Edge struct {
	to   int
	rev  int
	capa int
}

type MaxFlow struct {
	n   int
	pos [][2]int
	g   [][]_Edge
}

func newMaxFlow(n int) *MaxFlow {
	return &MaxFlow{
		n: n,
		g: make([][]_Edge, n),
	}
}

func (mf *MaxFlow) min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (mf *MaxFlow) AddEdge(from, to, capa int) int {
	m := len(mf.pos)
	mf.pos = append(mf.pos, [2]int{from, len(mf.g[from])})
	mf.g[from] = append(mf.g[from], _Edge{to, len(mf.g[to]), capa})
	mf.g[to] = append(mf.g[to], _Edge{from, len(mf.g[from]) - 1, 0})
	return m
}

func (mf *MaxFlow) Flow(s, t int) int {
	return mf.FlowL(s, t, int(1e+18))
}

func (mf *MaxFlow) FlowL(s, t, flowLim int) int {
	level := make([]int, mf.n)
	iter := make([]int, mf.n)
	bfs := func() {
		for i, _ := range level {
			level[i] = -1
		}
		level[s] = 0
		q := make([]int, 0, mf.n)
		q = append(q, s)
		for len(q) != 0 {
			v := q[0]
			q = q[1:]
			for _, e := range mf.g[v] {
				if e.capa == 0 || level[e.to] >= 0 {
					continue
				}
				level[e.to] = level[v] + 1
				if e.to == t {
					return
				}
				q = append(q, e.to)
			}
		}
	}
	var dfs func(v, up int) int
	dfs = func(v, up int) int {
		if v == s {
			return up
		}
		res := 0
		lv := level[v]
		for ; iter[v] < len(mf.g[v]); iter[v]++ {
			e := &mf.g[v][iter[v]]
			if lv <= level[e.to] || mf.g[e.to][e.rev].capa == 0 {
				continue
			}
			d := dfs(e.to, mf.min(up-res, mf.g[e.to][e.rev].capa))
			if d <= 0 {
				continue
			}
			mf.g[v][iter[v]].capa += d
			mf.g[e.to][e.rev].capa -= d
			res += d
			if res == up {
				break
			}
		}
		return res
	}
	flow := 0
	for flow < flowLim {
		bfs()
		if level[t] == -1 {
			break
		}
		for i, _ := range iter {
			iter[i] = 0
		}
		for flow < flowLim {
			f := dfs(t, flowLim-flow)
			if f == 0 {
				break
			}
			flow += f
		}
	}
	return flow
}
