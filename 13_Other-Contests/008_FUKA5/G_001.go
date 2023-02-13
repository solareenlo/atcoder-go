package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Pair struct {
	x, y int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for {
		var h, w int
		fmt.Fscan(in, &h, &w)
		var aa, bb float64
		fmt.Fscan(in, &aa, &bb)
		a := int(math.Round(aa * 1000.0))
		b := int(math.Round(bb * 1000.0))
		if h == 0 {
			break
		}
		s := make([]string, h)
		for i := 0; i < h; i++ {
			fmt.Fscan(in, &s[i])
		}

		G := NewMfgraph(int(h*w + 2))
		S := h * w
		T := h*w + 1
		E := make([][]Pair, h*w+2)
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				x := i*w + j
				if s[i][j] == '.' {
					E[S] = append(E[S], Pair{x, G.AddEdge(S, x, 0)})
					E[x] = append(E[x], Pair{T, G.AddEdge(x, T, a)})
				} else {
					E[S] = append(E[S], Pair{x, G.AddEdge(S, x, a)})
					E[x] = append(E[x], Pair{T, G.AddEdge(x, T, 0)})
				}
				if j != w-1 {
					y := i*w + j + 1
					E[x] = append(E[x], Pair{y, G.AddEdge(x, y, b)})
					E[y] = append(E[y], Pair{x, G.AddEdge(y, x, b)})
				}
				if i != h-1 {
					y := (i+1)*w + j
					E[x] = append(E[x], Pair{y, G.AddEdge(x, y, b)})
					E[y] = append(E[y], Pair{x, G.AddEdge(y, x, b)})
				}
			}
		}

		Ans := G.Flow(S, T)
		fmt.Fprintln(out, float64(Ans)/1000.0)

		ans := make([][]string, h)
		for i := range ans {
			ans[i] = make([]string, w)
			for j := range ans[i] {
				ans[i][j] = "."
			}
		}
		ret := G.MinCut(S)

		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				if ret[i*w+j] {
					ans[i][j] = "#"
				}
			}
		}

		sum := 0
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				if s[i][j] != ans[i][j][0] {
					sum += a
				}
				if i != h-1 {
					if ans[i][j] != ans[i+1][j] {
						sum += b
					}
				}
				if j != w-1 {
					if ans[i][j] != ans[i][j+1] {
						sum += b
					}
				}
			}
		}

		for i := 0; i < h; i++ {
			for j := range ans[i] {
				fmt.Fprint(out, ans[i][j])
			}
			fmt.Fprintln(out)
		}
	}
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
