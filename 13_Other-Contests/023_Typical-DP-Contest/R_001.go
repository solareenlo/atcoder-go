package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 309

	var n int
	fmt.Fscan(in, &n)
	scc := NewSccGraph(n + 1)
	for i := 1; i <= n; i++ {
		scc.AddEdge(0, i)
	}
	var g [N][N]bool
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Fscan(in, &g[i][j])
			if g[i][j] {
				scc.AddEdge(i, j)
			}
		}
	}
	v := scc.Scc()
	m := len(v)
	var bel [N]int
	for i := 1; i <= m; i++ {
		for _, j := range v[i-1] {
			bel[j] = i
		}
	}
	var e [N][N]bool
	for i := 1; i <= n; i++ {
		e[bel[0]][bel[i]] = true
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if g[i][j] {
				e[bel[i]][bel[j]] = true
			}
		}
	}

	var cmax func(*int, int)
	cmax = func(x *int, y int) {
		if *x < y {
			*x = y
		}
	}

	var dp [N][N]int
	for i := 1; i <= m; i++ {
		s := len(v[i-1])
		for j := 1; j < i; j++ {
			for k := 1; k < i; k++ {
				fj := e[j][i]
				fk := e[k][i]
				nw := s + dp[j][k]
				if fj {
					cmax(&dp[i][k], nw)
				}
				if fk {
					cmax(&dp[i][j], nw)
				}
				if fj && fk {
					cmax(&dp[i][i], nw)
				}
			}
		}
	}
	ans := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= m; j++ {
			cmax(&ans, dp[i][j])
		}
	}
	fmt.Println(ans)
}

type TwoSat struct {
	n      int
	answer []bool
	scc    *SccGraph
}

func NewTwosat(n int) *TwoSat {
	ts := &TwoSat{
		n:      n,
		answer: make([]bool, n),
		scc:    NewSccGraph(n * 2),
	}
	return ts
}

func (ts *TwoSat) internalJudge(f bool, a int, b int) int {
	if f {
		return a
	}
	return b
}

func (ts *TwoSat) AddClause(i int, f bool, j int, g bool) {
	if !(0 <= i && i < ts.n) {
		panic("")
	}
	if !(0 <= j && j < ts.n) {
		panic("")
	}
	ts.scc.AddEdge(2*i+ts.internalJudge(f, 0, 1), 2*j+ts.internalJudge(g, 1, 0))
	ts.scc.AddEdge(2*j+ts.internalJudge(g, 0, 1), 2*i+ts.internalJudge(f, 1, 0))
}

func (ts *TwoSat) Satisfiable() bool {
	id := ts.scc.SccIds().Second

	for i := 0; i < ts.n; i++ {
		if id[2*i] == id[2*i+1] {
			return false
		}
		ts.answer[i] = id[2*i] < id[2*i+1]
	}
	return true
}

func (ts *TwoSat) Answer() []bool {
	return ts.answer
}

type sccFromToPair struct {
	first, second int
}

type sccIdPair struct {
	First  int
	Second []int
}

type csr struct {
	start []int
	elist []int
}

func initCsr(n int, edges []*sccFromToPair) *csr {
	var ret csr
	ret.start = make([]int, n+1)
	ret.elist = make([]int, len(edges))
	for _, e := range edges {
		ret.start[e.first+1]++
	}
	for i := 1; i <= n; i++ {
		ret.start[i] += ret.start[i-1]
	}
	counter := make([]int, len(ret.start))
	copy(counter, ret.start)
	for _, e := range edges {
		ret.elist[counter[e.first]] = e.second
		counter[e.first]++
	}
	return &ret
}

type SccGraph struct {
	n     int
	edges []*sccFromToPair
}

func NewSccGraph(n int) *SccGraph {
	var s SccGraph
	s.n = n
	return &s
}

func (s *SccGraph) AddEdge(from, to int) {
	s.edges = append(s.edges, &sccFromToPair{from, to})
}

func (s *SccGraph) min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func (s *SccGraph) SccIds() sccIdPair {
	g := initCsr(s.n, s.edges)
	nowOrd, groupNum := 0, 0
	visited := make([]int, 0, s.n)
	low := make([]int, s.n)
	ord := make([]int, s.n)
	ids := make([]int, s.n)
	for i := 0; i < s.n; i++ {
		ord[i] = -1
	}
	var dfs func(v int)
	dfs = func(v int) {
		low[v] = nowOrd
		ord[v] = nowOrd
		nowOrd++
		visited = append(visited, v)
		for i := g.start[v]; i < g.start[v+1]; i++ {
			to := g.elist[i]
			if ord[to] == -1 {
				dfs(to)
				low[v] = s.min(low[v], low[to])
			} else {
				low[v] = s.min(low[v], ord[to])
			}
		}
		if low[v] == ord[v] {
			for {
				u := visited[len(visited)-1]
				visited = visited[:len(visited)-1]
				ord[u] = s.n
				ids[u] = groupNum
				if u == v {
					break
				}
			}
			groupNum++
		}
	}
	for i := 0; i < s.n; i++ {
		if ord[i] == -1 {
			dfs(i)
		}
	}
	for i := 0; i < len(ids); i++ {
		ids[i] = groupNum - 1 - ids[i]
	}
	return sccIdPair{groupNum, ids}
}

func (s *SccGraph) Scc() [][]int {
	ids := s.SccIds()
	groupNum := ids.First
	counts := make([]int, groupNum)
	for _, x := range ids.Second {
		counts[x]++
	}
	groups := make([][]int, ids.First)
	for i := 0; i < groupNum; i++ {
		groups[i] = make([]int, 0, counts[i])
	}
	for i := 0; i < s.n; i++ {
		groups[ids.Second[i]] = append(groups[ids.Second[i]], i)
	}
	return groups
}
