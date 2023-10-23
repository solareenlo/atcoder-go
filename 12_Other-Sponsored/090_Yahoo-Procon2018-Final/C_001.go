package main

import (
	"bufio"
	"fmt"
	"os"
)

type pair struct {
	x, y int
}

type Pair struct {
	x []int
	y []pair
}

type tuple struct {
	x, y, z int
}

var ans, sz [1 << 17]int
var dead [1 << 17]bool
var G [1 << 17][]int
var Q [1 << 17][]pair

func dfs_sz(u, p int) {
	sz[u] = 1
	for _, v := range G[u] {
		if v != p && !dead[v] {
			dfs_sz(v, u)
			sz[u] += sz[v]
		}
	}
}

func bfs_cd(u int) Pair {
	D := make([]int, 0)
	qs := make([]pair, 0)
	P := make([]tuple, 0)
	P = append(P, tuple{u, -1, 1})
	for len(P) > 0 {
		u, p, d := P[0].x, P[0].y, P[0].z
		P = P[1:]
		qs = append(qs, pair{u, d})
		if d >= len(D) {
			resize(&D, d+1)
		}
		D[d]++
		for _, v := range G[u] {
			if v != p && !dead[v] {
				P = append(P, tuple{v, u, d + 1})
			}
		}
	}
	return Pair{D, qs}
}

func calc(D []int, qs []pair, w int) {
	for _, tmp0 := range qs {
		u, d := tmp0.x, tmp0.y
		for _, tmp1 := range Q[u] {
			i, k := tmp1.x, tmp1.y
			if k >= d && k-d < len(D) {
				ans[i] += D[k-d] * w
			}
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
		for _, v := range G[centroid] {
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
	D := make([]int, n)
	D[0]++
	qs := make([]pair, 0)
	qs = append(qs, pair{centroid, 0})
	for _, v := range G[centroid] {
		if dead[v] {
			continue
		}
		dfs_cd(v)
		tmp := bfs_cd(v)
		nD, nqs := tmp.x, tmp.y
		calc(nD, nqs, -1)
		for i := 0; i < len(nD); i++ {
			D[i] += nD[i]
		}
		qs = append(qs, nqs...)
	}
	calc(D, qs, 1)
	dead[centroid] = false
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, q int
	fmt.Fscan(in, &n, &q)
	for i := 0; i < n-1; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		G[a] = append(G[a], b)
		G[b] = append(G[b], a)
	}
	for i := 0; i < q; i++ {
		var v, k int
		fmt.Fscan(in, &v, &k)
		v--
		Q[v] = append(Q[v], pair{i, k})
	}
	dfs_cd(0)
	for i := 0; i < q; i++ {
		fmt.Fprintln(out, ans[i])
	}
}

func resize(a *[]int, n int) {
	if len(*a) > n {
		*a = (*a)[:n]
	} else {
		n = n - len(*a)
		for i := 0; i < n; i++ {
			*a = append(*a, 0)
		}
	}
}
