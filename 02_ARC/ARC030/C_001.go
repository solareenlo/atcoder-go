package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var N, M, K int
	fmt.Scan(&N, &M, &K)

	L := make([]string, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&L[i])
	}

	g := NewSccGraph(N)
	G := make([][]int, 300)
	for i := 0; i < M; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		a--
		b--
		G[a] = append(G[a], b)
		g.AddEdge(a, b)
	}

	D := make([]int, 300)
	V := g.Scc()
	S := len(V)
	for i := 0; i < S; i++ {
		for _, v := range V[i] {
			D[v] = i
		}
	}

	dp := [301][301]string{}
	ans := ""
	for i := 0; i < S; i++ {
		var s string
		for _, v := range V[i] {
			s += L[v]
		}
		tmp := strings.Split(s, "")
		sort.Strings(tmp)
		s = strings.Join(tmp, "")
		for j := K; j >= 0; j-- {
			if j != 0 && dp[i][j] == "" {
				continue
			}
			t := dp[i][j]
			for k := 1; j+k <= K && k <= len(s); k++ {
				t += string(s[k-1])
				if dp[i][j+k] == "" || dp[i][j+k] > t {
					dp[i][j+k] = t
				} else {
					break
				}

			}
		}
		for _, v := range V[i] {
			for _, t := range G[v] {
				if i < D[t] {
					for k := 0; k <= K; k++ {
						if dp[D[t]][k] == "" || dp[D[t]][k] > dp[i][k] {
							dp[D[t]][k] = dp[i][k]
						}
					}
				}
			}
		}
		if ans == "" || ans > dp[i][K] {
			ans = dp[i][K]
		}
	}
	if ans == "" {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}
}

type sccIdPair struct {
	First  int
	Second []int
}

type sccFromToPair struct {
	first, second int
}

type SccGraph struct {
	n     int
	edges []*sccFromToPair
}

type csr struct {
	start []int
	elist []int
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
