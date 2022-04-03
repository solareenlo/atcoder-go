package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 200005

var (
	id   int
	G    = make([][]int, N)
	leaf = make([]int, N)
	R    = make([]int, N)
	L    = make([]int, N)
)

func dfs(v, p int) {
	if v != 1 && len(G[v]) == 1 {
		id++
		leaf[v] = id
		return
	}
	for _, u := range G[v] {
		if u != p {
			dfs(u, v)
		}
	}
}

func dfs2(v, p int) {
	if leaf[v] != 0 {
		R[v] = leaf[v]
		L[v] = leaf[v]
		return
	}
	L[v] = 1 << 60
	R[v] = -1 << 60
	for _, u := range G[v] {
		if u == p {
			continue
		}
		dfs2(u, v)
		L[v] = min(L[v], L[u])
		R[v] = max(R[v], R[u])
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	for i := 1; i <= n-1; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		G[u] = append(G[u], v)
		G[v] = append(G[v], u)
	}

	dfs(1, -1)
	dfs2(1, -1)

	for i := 1; i <= n; i++ {
		fmt.Fprintln(out, L[i], R[i])
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
