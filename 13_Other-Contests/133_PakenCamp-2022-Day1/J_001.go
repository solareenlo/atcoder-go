package main

import (
	"bufio"
	"fmt"
	"os"
)

type pair struct {
	x, y int
}

var G, Q [1 << 17][]pair
var B *Fenwick
var par [17][1 << 17]int
var dep [1 << 17]int
var dist [1 << 16]int
var X, ok, ng, mid [2 << 17]int

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	for i := 0; i < n-1; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		a--
		b--
		G[a] = append(G[a], pair{b, c})
		G[b] = append(G[b], pair{a, c})
	}
	dfs(0, 0, 0)
	for k := 0; k < 16; k++ {
		for u := 0; u < n; u++ {
			par[k+1][u] = par[k][par[k][u]]
		}
	}
	var q int
	fmt.Fscan(in, &q)
	for i := 0; i < q; i++ {
		var s, t int
		fmt.Fscan(in, &s, &t)
		s--
		t--
		l := lca(s, t)
		Q[s] = append(Q[s], pair{i, 1})
		Q[t] = append(Q[t], pair{i, 1})
		Q[l] = append(Q[l], pair{i, -2})
		dist[i] = dep[s] + dep[t] - 2*dep[l]
	}
	for i := 0; i < 2*q; i++ {
		ng[i] = 2 << 17
	}
	for j := 0; j < 18; j++ {
		for i := 0; i < 2*q; i++ {
			mid[i] = (ok[i] + ng[i]) >> 1
		}
		for i := 0; i < 2*q; i++ {
			X[i] = 0
		}
		B = NewFenwick(2 << 17)
		efs(0, -1)
		for i := 0; i < 2*q; i++ {
			if X[i] <= (dist[i/2]-i%2)/2 {
				ok[i] = mid[i]
			} else {
				ng[i] = mid[i]
			}
		}
	}
	for i := 0; i < q; i++ {
		fmt.Println((ok[2*i] + ok[2*i+1]) / 2)
	}
}

func dfs(u, p, d int) {
	par[0][u] = p
	dep[u] = d
	for _, tmp := range G[u] {
		v := tmp.x
		if v != p {
			dfs(v, u, d+1)
		}
	}
}

func lca(u, v int) int {
	if dep[u] > dep[v] {
		u, v = v, u
	}
	d := dep[v] - dep[u]
	for k := 0; k < 17; k++ {
		if ((d >> k) & 1) != 0 {
			v = par[k][v]
		}
	}
	if u == v {
		return u
	}
	for k := 16; k >= 0; k-- {
		if par[k][u] != par[k][v] {
			u = par[k][u]
			v = par[k][v]
		}
	}
	return par[0][u]
}

func efs(u, p int) {
	for _, tmp := range Q[u] {
		i := tmp.x
		k := tmp.y
		for _, j := range [2]int{2 * i, 2*i + 1} {
			X[j] += B.Sum(0, mid[j]) * k
		}
	}
	for _, tmp := range G[u] {
		v := tmp.x
		c := tmp.y
		if v != p {
			B.Add(c, 1)
			efs(v, u)
			B.Add(c, -1)
		}
	}
}

type Fenwick struct {
	n    int
	data []uint
}

func NewFenwick(n int) *Fenwick {
	fen := &Fenwick{
		n:    n,
		data: make([]uint, n),
	}
	for idx := range fen.data {
		fen.data[idx] = 0
	}
	return fen
}

func (fen *Fenwick) Add(pos, x int) {
	if !(0 <= pos && pos < fen.n) {
		panic("")
	}
	pos++
	for pos <= fen.n {
		fen.data[pos-1] += uint(x)
		pos += pos & -pos
	}
}

func (fen *Fenwick) Sum(l, r int) int {
	if !(0 <= l && l <= r && r <= fen.n) {
		panic("")
	}
	return int(fen.sum(r) - fen.sum(l))
}

func (fen *Fenwick) sum(r int) uint {
	s := uint(0)
	for r > 0 {
		s += fen.data[r-1]
		r -= r & -r
	}
	return s
}
