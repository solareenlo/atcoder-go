package main

import (
	"bufio"
	"fmt"
	"os"
)

type edge struct {
	to, cap, rev int
}

var n, m int
var c [505]int
var S, T int
var used [1005]bool
var G [1005][]edge

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = int(1e18)

	fmt.Fscan(in, &n, &m)
	for i := 2; i <= n-1; i++ {
		fmt.Fscan(in, &c[i])
	}

	S = 2*n + 1
	T = 2*n + 2
	for i := 1; i <= m; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		if c[u] == -1 && c[v] == -1 {
			fmt.Println(-1)
			return
		}
		if c[u] == -1 {
			add_edge(S, v, INF)
		} else if c[v] == -1 {
			add_edge(u+n, T, INF)
		} else {
			add_edge(u+n, v, INF)
		}
	}
	for i := 2; i <= n-1; i++ {
		if c[i] >= 0 {
			add_edge(i, i+n, c[i])
		}
	}

	ans := 0
	for {
		for i := 1; i <= T; i++ {
			used[i] = false
		}
		flow := dfs(S, INF)
		if flow <= 0 {
			break
		}
		ans += flow
	}
	fmt.Println(ans)
}

func add_edge(s, t, Cap int) {
	G[s] = append(G[s], edge{t, Cap, len(G[t])})
	G[t] = append(G[t], edge{s, 0, len(G[s]) - 1})
}

func dfs(v, f int) int {
	used[v] = true
	if v == T {
		return f
	}
	for i := 0; i < len(G[v]); i++ {
		if used[G[v][i].to] || G[v][i].cap <= 0 {
			continue
		}
		ret := dfs(G[v][i].to, min(f, G[v][i].cap))
		if ret > 0 {
			G[v][i].cap -= ret
			G[G[v][i].to][G[v][i].rev].cap += ret
			return ret
		}
	}
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
