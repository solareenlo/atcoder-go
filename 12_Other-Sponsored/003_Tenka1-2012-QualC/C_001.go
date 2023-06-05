package main

import (
	"bufio"
	"fmt"
	"os"
)

var N, L int
var visited map[State]struct{}
var cntStates, nextStates []State
var emptyChair bool

func main() {
	in := bufio.NewReader(os.Stdin)

	var M int
	fmt.Fscan(in, &N, &M, &L)
	to := make([]int, N)
	for i := range to {
		to[i] = -1
	}
	for i := 0; i < L; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		to[a] = b
	}
	paths := make([][]int, 0)
	cycles := make([][]int, 0)
	decomposeDegOneOneGraph(to, &paths, &cycles)
	var initst State
	for i := range paths {
		initst.paths[len(paths[i])-1]++
	}
	for i := range cycles {
		initst.cycles[len(cycles[i])]++
	}
	emptyChair = L < N

	visited = make(map[State]struct{})
	visited[initst] = struct{}{}
	cntStates = make([]State, 0)
	nextStates = make([]State, 0)
	nextStates = append(nextStates, initst)
	time := 0
	for len(nextStates) != 0 {
		cntStates, nextStates = nextStates, cntStates
		for i := range cntStates {
			ok := true
			for j := 1; j <= L; j++ {
				if ok && cntStates[i].paths[j] == 0 {
					ok = true
				} else {
					ok = false
				}
			}
			for j := 2; j <= L; j++ {
				if ok && cntStates[i].cycles[j] == 0 {
					ok = true
				} else {
					ok = false
				}
			}
			if ok {
				fmt.Println(time)
				return
			}
		}
		for len(cntStates) != 0 {
			s := cntStates[len(cntStates)-1]
			cntStates = cntStates[:len(cntStates)-1]
			dfs(s, 0, M)
		}
		time += 10
	}
	fmt.Println(-1)
}

func decomposeDegOneOneGraph(g []int, paths, cycles *[][]int) {
	n := len(g)
	inDegree := make([]int, n)
	for i := 0; i < N; i++ {
		if g[i] >= 0 {
			inDegree[g[i]]++
		}
	}
	vis := make([]bool, n)
	*paths = make([][]int, 0)
	*cycles = make([][]int, 0)
	for i := 0; i < N; i++ {
		if inDegree[i] == 0 {
			v := i
			path := make([]int, 0)
			for v >= 0 {
				vis[v] = true
				path = append(path, v)
				v = g[v]
			}
			*paths = append(*paths, path)
		}
	}
	for i := 0; i < N; i++ {
		if !vis[i] {
			v := i
			cycle := make([]int, 0)
			for !vis[v] {
				vis[v] = true
				cycle = append(cycle, v)
				v = g[v]
			}
			*cycles = append(*cycles, cycle)
		}
	}
}

type State struct {
	paths, cycles [16]int
}

func dfs(s State, i, m int) {
	if i == (L+1)*2 {
		if _, ok := visited[s]; !ok {
			visited[s] = struct{}{}
			nextStates = append(nextStates, s)
		}
		return
	}
	dfs(s, i+1, m)
	if i < L+1 {
		Len := i
		if Len > 0 && s.paths[Len] > 0 {
			s.paths[Len]--
			for j := 1; j <= Len && j <= m; j++ {
				s.paths[Len-j]++
				dfs(s, i, m-j)
				s.paths[Len-j]--
			}
			s.paths[Len]++
		}
	} else {
		Len := i - (L + 1)
		if Len > 1 && s.cycles[Len] > 0 {
			s.cycles[Len]--
			for j := 2; j <= Len && j <= m; j++ {
				s.cycles[Len-j+1]++
				dfs(s, i, m-j)
				s.cycles[Len-j+1]--
			}
			if emptyChair {
				for j := 1; j < Len && j <= m; j++ {
					s.paths[Len-j+1]++
					dfs(s, i, m-j)
					s.paths[Len-j+1]--
				}
			}
			s.cycles[Len]++
		}
	}
}
