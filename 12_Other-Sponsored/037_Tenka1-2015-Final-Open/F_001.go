package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const MAX = 1 << 17

var id, n int
var sz, in [MAX]int
var par [17][MAX]int
var D [MAX]int
var V [MAX]int
var G [MAX][]int

func dfs(u int) {
	in[u] = id
	id++
	for i := 0; i < len(G[u]); i++ {
		v := G[u][i]
		if v != par[0][u] {
			par[0][v] = u
			D[v] = D[u] + 1
			dfs(v)
		}
	}
}

func lca(u int, v int) int {
	if D[u] > D[v] {
		u, v = v, u
	}
	d := D[v] - D[u]
	for i := 0; i < 17; i++ {
		if d>>i&1 == 1 {
			v = par[i][v]
		}
	}
	if u == v {
		return u
	}
	for i := 16; i >= 0; i-- {
		if par[i][u] != par[i][v] {
			u = par[i][u]
			v = par[i][v]
		}
	}
	return par[0][u]
}

func main() {
	IN := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(IN, &n)
	for i := 0; i < n-1; i++ {
		var a, b int
		fmt.Fscan(IN, &a, &b)
		a--
		b--
		G[a] = append(G[a], b)
		G[b] = append(G[b], a)
	}
	dfs(0)
	for i := 0; i < 16; i++ {
		for u := 0; u < n; u++ {
			par[i+1][u] = par[i][par[i][u]]
		}
	}
	var q int
	fmt.Fscan(IN, &q)
	for i := 0; i < q; i++ {
		var m, k int
		fmt.Fscan(IN, &m, &k)
		for j := 0; j < m; j++ {
			fmt.Fscan(IN, &V[j])
			V[j]--
		}
		sort.Slice(V[:m], func(a, b int) bool {
			return in[V[a]] < in[V[b]]
		})
		ans := 0
		for j := 0; j+k-1 < m; j++ {
			ans = max(ans, D[lca(V[j], V[j+k-1])])
		}
		fmt.Fprintln(out, ans)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
