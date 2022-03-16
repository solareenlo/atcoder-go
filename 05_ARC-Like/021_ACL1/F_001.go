package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	n    int
	TYPE = [35]int{}
	pos  = make([][]int, 35)
	posa = make([][]int, 35)
)

func getType(p, l, r int) string {
	if p <= l {
		return "L"
	} else if p >= r {
		return "R"
	}
	return "M"
}

func solve(l, r int) bool {
	type pair struct{ x, y int }
	match := make([]pair, 0)
	for i := 1; i <= n; i++ {
		TYPE[i] = 0
		Type := getType(pos[i][0], l, r) + getType(pos[i][1], l, r) + getType(pos[i][2], l, r)
		if Type == "LLL" || Type == "RRR" {
			return false
		}
		if Type == "LLR" {
			TYPE[i] = 1
		}
		if Type == "LRR" {
			TYPE[i] = 2
		}
		if Type == "LMR" {
			TYPE[i] = 3
		}
		if Type == "MMR" {
			match = append(match, pair{posa[i][0], pos[i][0]})
			match = append(match, pair{posa[i][2], pos[i][1]})
		}
		if Type == "LMM" {
			match = append(match, pair{posa[i][0], pos[i][1]})
			match = append(match, pair{posa[i][2], pos[i][2]})
		}
		if Type == "MMM" {
			for j := 0; j < 3; j++ {
				match = append(match, pair{posa[i][j], pos[i][j]})
			}
		}
		if Type == "LLM" {
			match = append(match, pair{posa[i][2], pos[i][2]})
		}
		if Type == "MRR" {
			match = append(match, pair{posa[i][0], pos[i][0]})
		}
	}
	check := func(a, b int) bool {
		return pos[a][0] < pos[b][0] || pos[a][2] < pos[b][2]
	}
	cross := func(a, b, c, d int) int {
		tmp1, tmp2 := 0, 0
		if a < c {
			tmp1 = 1
		}
		if b < d {
			tmp2 = 1
		}
		return tmp1 ^ tmp2
	}
	ts := NewTwosat(n + 1)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if TYPE[i] == 1 && TYPE[j] == 2 && !check(i, j) {
				return false
			}
		}
	}
	sort.Slice(match, func(i, j int) bool {
		return match[i].x < match[j].x
	})
	for i := 1; i < len(match); i++ {
		if match[i].y < match[i-1].y {
			return false
		}
	}
	for i := 1; i <= n; i++ {
		if TYPE[i] != 3 {
			continue
		}
		must := 0
		for _, s := range match {
			if cross(s.x, s.y, posa[i][0], pos[i][1]) != 0 {
				must |= 2
			}
			if cross(s.x, s.y, posa[i][2], pos[i][1]) != 0 {
				must |= 1
			}
		}
		for j := 1; j <= n; j++ {
			if TYPE[j] == 1 {
				if !check(j, i) {
					must |= 1
				}
			} else if TYPE[j] == 2 {
				if !check(i, j) {
					must |= 2
				}
			} else if i < j && TYPE[j] == 3 {
				if !check(i, j) || cross(posa[i][0], pos[i][1], posa[j][2], pos[j][1]) != 0 {
					ts.AddClause(i, false, j, true)
				}
				if !check(j, i) || cross(posa[i][2], pos[i][1], posa[j][0], pos[j][1]) != 0 {
					ts.AddClause(i, true, j, false)
				}
				if cross(posa[i][0], pos[i][1], posa[j][0], pos[j][1]) != 0 {
					ts.AddClause(i, false, j, false)
				}
				if cross(posa[i][2], pos[i][1], posa[j][2], pos[j][1]) != 0 {
					ts.AddClause(i, true, j, true)
				}
			}
		}
		if must == 3 {
			return false
		}
		if must&1 != 0 {
			ts.AddClause(i, true, i, true)
		}
		if must&2 != 0 {
			ts.AddClause(i, false, i, false)
		}
	}
	return ts.Satisfiable()
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)

	A := make([]int, 111)
	for i := 1; i <= 3*n; i++ {
		fmt.Fscan(in, &A[i])
		posa[A[i]] = append(posa[A[i]], i)
	}

	B := make([]int, 111)
	for i := 1; i <= 3*n; i++ {
		fmt.Fscan(in, &B[i])
		pos[B[i]] = append(pos[B[i]], i)
	}

	ans := -1
	for l := 0; l <= 3*n; l++ {
		for r := l + 1; r <= 3*n+1; r++ {
			if r-l-1 > ans && solve(l, r) {
				ans = r - l - 1
			}
		}
	}
	if ans > -1 {
		fmt.Println(3*n - ans)
	} else {
		fmt.Println(-1)
	}
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
