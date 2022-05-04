package main

import (
	"bufio"
	"fmt"
	"os"
)

type pair struct{ x, y int }

var (
	G = make([][]int, 1<<17)
	E = make([][]pair, 1<<17)
	L = make([]int, 1<<17)
	P = make([]int, 1<<17)
	D = make([]int, 1<<17)
	V = make([]int, 0)
)

func root(u int) int {
	if P[u] == u {
		return u
	}
	P[u] = root(P[u])
	return P[u]
}

func dfs(u, p int) {
	for _, v := range G[u] {
		if v == p {
			continue
		}
		if D[v] != -1 {
			V = append(V, u)
			continue
		}
		D[v] = D[u] + 1
		dfs(v, u)
		P[v] = u
	}
	for j := range E[u] {
		i := E[u][j].x
		v := E[u][j].y
		L[i] = D[u] + D[v] - 2*D[root(v)]
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		G[a] = append(G[a], b)
		G[b] = append(G[b], a)
	}
	var q int
	fmt.Fscan(in, &q)
	for i := 0; i < q; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		E[u] = append(E[u], pair{i, v})
		E[v] = append(E[v], pair{i, u})
	}

	for i := 0; i < n; i++ {
		P[i] = i
	}
	for i := 0; i < n; i++ {
		D[i] = -1
	}
	D[0] = 0
	dfs(0, -1)
	for _, s := range V {
		for i := 0; i < n; i++ {
			D[i] = -1
		}
		D[s] = 0
		Q := make([]int, 0)
		Q = append(Q, s)
		for len(Q) > 0 {
			u := Q[0]
			Q = Q[1:]
			for _, v := range G[u] {
				if D[v] == -1 {
					D[v] = D[u] + 1
					Q = append(Q, v)
				}
			}
		}
		for u := 0; u < n; u++ {
			for j := range E[u] {
				i := E[u][j].x
				v := E[u][j].y
				L[i] = min(L[i], D[u]+D[v])
			}
		}
	}
	for i := 0; i < q; i++ {
		fmt.Fprintln(out, L[i])
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
