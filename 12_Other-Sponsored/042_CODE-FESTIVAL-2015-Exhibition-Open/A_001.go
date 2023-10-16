package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 100000

var v, w [N][]int
var z []int
var a, b, f, c [N]int

func dfs(x int) {
	f[x] = 1
	for i := 0; i < len(v[x]); i++ {
		if f[v[x][i]] == 0 {
			dfs(v[x][i])
		}
	}
	z = append(z, x)
}

func dfs2(x, k int) {
	f[x] = 1
	c[x] = k
	for i := 0; i < len(w[x]); i++ {
		if f[w[x][i]] == 0 {
			dfs2(w[x][i], k)
		}
	}
}

func scc(n int) int {
	k := 0
	for i := 0; i < n; i++ {
		f[i] = 0
	}
	z = make([]int, 0)
	for i := 0; i < n; i++ {
		if f[i] == 0 {
			dfs(i)
		}
	}
	for i := 0; i < n; i++ {
		f[i] = 0
	}
	for i := len(z) - 1; i >= 0; i-- {
		if f[z[i]] == 0 {
			dfs2(z[i], k)
			k++
		}
	}
	return k
}

func dfs3(x int) {
	a[x] = 1
	for i := 0; i < len(v[x]); i++ {
		if a[v[x][i]] == 0 {
			dfs3(v[x][i])
		}
	}
}

func dfs4(x int) {
	b[x] = 1
	for i := 0; i < len(w[x]); i++ {
		if b[w[x][i]] == 0 {
			dfs4(w[x][i])
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	mf := NewMfgraph(n)
	for i := 0; i < m; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		mf.AddEdge(x-1, y-1, 1)
	}
	mf.Flow(0, n-1)
	edge := mf.GetEdges()
	for i := 0; i < len(edge); i++ {
		v[edge[i].a] = append(v[edge[i].a], edge[i].b)
		w[edge[i].b] = append(w[edge[i].b], edge[i].a)
	}
	scc(n)
	dfs3(0)
	dfs4(n - 1)

	var q int
	fmt.Fscan(in, &q)
	for i := 0; i < q; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		x--
		y--
		if a[x] == 1 && b[y] == 1 {
			fmt.Fprintln(out, "NO YES")
		} else if a[y] == 1 && b[x] == 1 {
			fmt.Fprintln(out, "NO YES")
		} else if a[x] == 1 && a[y] == 1 {
			fmt.Fprintln(out, "YES NO")
		} else if b[x] == 1 && b[y] == 1 {
			fmt.Fprintln(out, "YES NO")
		} else if a[x] == 0 && a[y] == 0 && b[x] == 0 && b[y] == 0 && c[x] == c[y] {
			fmt.Fprintln(out, "YES NO")
		} else {
			fmt.Fprintln(out, "YES YES")
		}
	}
	return
}

type mfPreEdge struct{ to, rev, cap int }
type mfEdge struct{ from, to, cap, flow int }
type mfPos struct{ x, y int }
type Mfgraph struct {
	n   int
	pos []mfPos
	g   [][]mfPreEdge
}

func NewMfgraph(n int) *Mfgraph {
	g := make([][]mfPreEdge, n)
	pos := make([]mfPos, 0)
	return &Mfgraph{n, pos, g}
}

func (q *Mfgraph) AddEdge(from, to, cap int) int {
	m := len(q.pos)
	formId := len(q.g[from])
	toId := len(q.g[to])
	q.pos = append(q.pos, mfPos{from, formId})
	if from == to {
		toId++
	}
	q.g[from] = append(q.g[from], mfPreEdge{to, toId, cap})
	q.g[to] = append(q.g[to], mfPreEdge{from, formId, 0})
	return m
}

func (q *Mfgraph) GetEdge(i int) mfEdge {
	e := q.g[q.pos[i].x][q.pos[i].y]
	re := q.g[e.to][e.rev]
	return mfEdge{q.pos[i].x, e.to, e.cap + re.cap, re.cap}
}

type P struct {
	a, b int
}

func (q *Mfgraph) GetEdges() []P {
	v := make([]P, 0)
	for i := 0; i < len(q.g); i++ {
		for j := 0; j < len(q.g[i]); j++ {
			if q.g[i][j].cap > 0 {
				v = append(v, P{i, q.g[i][j].to})
			}
		}
	}
	return v
}

func (q *Mfgraph) Edges() []mfEdge {
	m := len(q.pos)
	res := make([]mfEdge, 0)
	for i := 0; i < m; i++ {
		res = append(res, q.GetEdge(i))
	}
	return res
}

func (q *Mfgraph) ChangeEdge(i, newcap, newflow int) {
	e := &(q.g[q.pos[i].x][q.pos[i].y])
	re := &(q.g[e.to][e.rev])
	e.cap = newcap - newflow
	re.cap = newflow
}

func (q *Mfgraph) Flow(s, t int) int {
	return q.FlowCapped(s, t, 1<<62)
}

func (q *Mfgraph) FlowCapped(s, t, flowlimit int) int {
	level := make([]int, q.n)
	iter := make([]int, q.n)
	bfs := func() {
		for i := 0; i < q.n; i++ {
			level[i] = -1
		}
		level[s] = 0
		que := make([]int, 0, q.n)
		que = append(que, s)
		for len(que) > 0 {
			v := que[0]
			que = que[1:]
			for _, e := range q.g[v] {
				if e.cap == 0 || level[e.to] >= 0 {
					continue
				}
				level[e.to] = level[v] + 1
				if e.to == t {
					return
				}
				que = append(que, e.to)
			}
		}
	}
	var dfs func(int, int) int
	dfs = func(v, up int) int {
		if v == s {
			return up
		}
		res := 0
		level_v := level[v]
		for i := iter[v]; i < len(q.g[v]); i++ {
			e := q.g[v][i]
			cap := q.g[e.to][e.rev].cap
			if level_v <= level[e.to] || cap == 0 {
				continue
			}
			newup := up - res
			if cap < up-res {
				newup = cap
			}
			d := dfs(e.to, newup)
			if d <= 0 {
				continue
			}
			q.g[v][i].cap += d
			q.g[e.to][e.rev].cap -= d
			res += d
			if res == up {
				return res
			}
		}
		level[v] = q.n
		return res
	}
	flow := 0
	for flow < flowlimit {
		bfs()
		if level[t] == -1 {
			break
		}
		for i := 0; i < q.n; i++ {
			iter[i] = 0
		}
		f := dfs(t, flowlimit-flow)
		if f == 0 {
			break
		}
		flow += f
	}
	return flow
}

func (q *Mfgraph) MinCut(s int) []bool {
	visited := make([]bool, q.n)
	que := make([]int, 0, q.n)
	que = append(que, s)
	for len(que) > 0 {
		p := que[0]
		que = que[1:]
		visited[p] = true
		for _, e := range q.g[p] {
			if e.cap > 0 && !visited[e.to] {
				visited[e.to] = true
				que = append(que, e.to)
			}
		}
	}
	return visited
}
