package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type pair struct {
	x, y int
}

type tuple struct {
	x, y, z int
}

const MAXN = 70000

var sz [MAXN]int
var G, P [MAXN][]pair
var dead [MAXN]bool

func dfs_sz(u, p int) {
	sz[u] = 1
	for _, tmp := range G[u] {
		v := tmp.x
		if v != p && !dead[v] {
			dfs_sz(v, u)
			sz[u] += sz[v]
		}
	}
}

func dfs_cd(u int) {
	dfs_sz(u, -1)
	n := sz[u]
	p := -1
	centroid := u
	for {
		w := -1
		for _, tmp := range G[centroid] {
			v := tmp.x
			if v != p && !dead[v] && sz[v] > n/2 {
				w = v
			}
		}
		if w == -1 {
			break
		}
		p = centroid
		centroid = w
	}
	dead[centroid] = true
	Q := make([]tuple, 0)
	Q = append(Q, tuple{centroid, -1, 0})
	for len(Q) > 0 {
		u, p, d := Q[0].x, Q[0].y, Q[0].z
		Q = Q[1:]
		P[u] = append(P[u], pair{centroid, d})
		for _, tmp := range G[u] {
			v, c := tmp.x, tmp.y
			if v != p && !dead[v] && d+c < 1e9 {
				Q = append(Q, tuple{v, u, d + c})
			}
		}
	}
	for _, tmp := range G[centroid] {
		v := tmp.x
		if !dead[v] {
			dfs_cd(v)
		}
	}
	dead[centroid] = false
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var X, H, dp [MAXN]int

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 0; i < n-1; i++ {
		var a, b, d int
		fmt.Fscan(in, &a, &b, &d)
		a--
		b--
		G[a] = append(G[a], pair{b, d})
		G[b] = append(G[b], pair{a, d})
	}
	dfs_cd(0)
	W := make([]Tuple, 0)
	for i := 0; i < m; i++ {
		var s, e, c int
		fmt.Fscan(in, &s, &e, &c, &X[i])
		c--
		for _, tmp := range P[c] {
			v, d := tmp.x, tmp.y
			W = append(W, Tuple{e + d, 0, v, i})
			W = append(W, Tuple{s - d, 1, v, i})
		}
	}
	sortTuple(W)
	for _, tmp := range W {
		t, v, i := tmp.x, tmp.y, tmp.z
		if t == 0 {
			H[v] = max(H[v], dp[i])
		}
		if t == 1 {
			dp[i] = max(dp[i], H[v]+X[i])
		}
	}
	ans := 0
	for i := 0; i < m; i++ {
		ans = max(ans, dp[i])
	}
	fmt.Println(ans)
}

type Tuple struct {
	a, x, y, z int
}

func sortTuple(tup []Tuple) {
	sort.Slice(tup, func(i, j int) bool {
		return tup[i].a < tup[j].a
	})
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
