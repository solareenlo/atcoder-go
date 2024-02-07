package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 55

	var n, m int
	fmt.Fscan(in, &n, &m)

	G := NewMfgraph(4*n + m + 3)
	var c [N]int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &c[i])
		for j := 1; j <= 4; j++ {
			G.AddEdge(4*n+m+1, 4*(i-1)+j, c[i])
		}
	}
	var a [N]int
	ans := 0
	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &a[i])
		ans += a[i]
		G.AddEdge(4*n+i, 4*n+m+2, a[i])
	}
	var l [N][N]int
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			fmt.Fscan(in, &l[i][j])
			for k := 1; k < l[i][j]; k++ {
				G.AddEdge(4*(j-1)+k, 4*n+i, 1e9)
			}
		}
	}
	ans -= G.Flow(4*n+m+1, 4*n+m+2)
	fmt.Println(ans)
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
