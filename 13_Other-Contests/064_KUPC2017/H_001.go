package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = int(1e18)

	var n, m int
	fmt.Fscan(in, &n, &m)
	v := make([]int, n)
	h := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &v[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &h[i])
	}
	vec := make([][]int, n)
	a := make([]int, m)
	x := make([]int, m)
	b := make([]int, m)
	y := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &a[i], &x[i], &b[i], &y[i])
		a[i]--
		b[i]--
		vec[a[i]] = append(vec[a[i]], x[i])
		vec[b[i]] = append(vec[b[i]], y[i])
	}
	r := make([]int, n+1)
	for i := 0; i < n; i++ {
		if h[i] > 0 {
			t := make([]int, 0)
			for _, j := range vec[i] {
				t = append(t, j-1)
				t = append(t, j)
			}
			vec[i] = t
		}
		vec[i] = append(vec[i], v[i])
		sort.Ints(vec[i])
		vec[i] = unique(vec[i])
		r[i+1] = r[i] + len(vec[i])
	}
	s := r[n]
	t := s + 1
	G := NewMfgraph(r[n] + 2)
	ans := 0
	for i := 0; i < n; i++ {
		if h[i] > 0 {
			ans += v[i] * h[i]
			G.AddEdge(s, r[i], v[i]*h[i])
			for j := 0; j < len(vec[i])-1; j++ {
				G.AddEdge(r[i]+j, r[i]+j+1, (v[i]-vec[i][j])*h[i])
			}
			G.AddEdge(r[i+1]-1, t, 0)
			for j := len(vec[i]) - 1; j > 0; j-- {
				G.AddEdge(r[i]+j, r[i]+j-1, INF)
			}
		} else {
			h[i] *= -1
			G.AddEdge(s, r[i], 0)
			for j := 0; j < len(vec[i])-1; j++ {
				G.AddEdge(r[i]+j, r[i]+j+1, vec[i][j]*h[i])
			}
			G.AddEdge(r[i+1]-1, t, v[i]*h[i])
			for j := len(vec[i]) - 1; j > 0; j-- {
				G.AddEdge(r[i]+j, r[i]+j-1, INF)
			}
		}
	}
	for i := 0; i < m; i++ {
		ita := s
		if x[i] != 0 {
			ita = r[a[i]] + lowerBound(vec[a[i]], x[i])
		}
		itb := s
		if y[i] != 0 {
			itb = r[b[i]] + lowerBound(vec[b[i]], y[i])
		}
		G.AddEdge(ita, itb, INF)
	}
	fmt.Println(ans - G.Flow(s, t))
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

func unique(a []int) []int {
	occurred := map[int]bool{}
	result := []int{}
	for i := range a {
		if occurred[a[i]] != true {
			occurred[a[i]] = true
			result = append(result, a[i])
		}
	}
	sort.Ints(result)
	n := len(result)
	for i := 0; i < n; i++ {
		a[i] = result[i]
	}
	return result
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
