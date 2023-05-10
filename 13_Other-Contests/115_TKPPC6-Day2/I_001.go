package main

import (
	"bufio"
	"fmt"
	"os"
)

type pair struct {
	x, y int
}

var G [1 << 17][]pair
var D, X [1 << 17]int
var P, W [17][1 << 17]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m, q int
	fmt.Fscan(in, &n, &m, &q)
	for i := 0; i < n-1; i++ {
		var u, v, w int
		fmt.Fscan(in, &u, &v, &w)
		u--
		v--
		G[u] = append(G[u], pair{v, w})
		G[v] = append(G[v], pair{u, w})
	}
	f(0)
	for k := 0; k < 16; k++ {
		for u := 0; u < n; u++ {
			P[k+1][u] = P[k][P[k][u]]
			W[k+1][u] = W[k][u] + W[k][P[k][u]]
		}
	}
	for i := 0; i < m+1; i++ {
		fmt.Fscan(in, &X[i])
		X[i]--
	}
	B := NewFenwick(m)
	for i := 0; i < m; i++ {
		B.Add(i, dist(X[i], X[i+1], 0))
	}
	for q > 0 {
		q--
		var t int
		fmt.Fscan(in, &t)
		if t == 1 {
			var a, b int
			fmt.Fscan(in, &a, &b)
			b--
			if a != 0 {
				B.Add(a-1, dist(X[a-1], b, 0)-dist(X[a-1], X[a], 0))
			}
			if a < m {
				B.Add(a, dist(b, X[a+1], 0)-dist(X[a], X[a+1], 0))
			}
			X[a] = b
		} else {
			var r int
			fmt.Fscan(in, &r)
			ans := (B.Sum(0, r)+dist(X[0], X[r], 0))/2 - dist(X[0], X[r], 1)
			fmt.Fprintln(out, ans)
		}
	}
}

func f(u int) {
	for _, tmp := range G[u] {
		v := tmp.x
		w := tmp.y
		if v != P[0][u] {
			P[0][v] = u
			W[0][v] = w
			D[v] = D[u] + 1
			f(v)
		}
	}
}

func dist(u, v, t int) int {
	res := 0
	if D[u] > D[v] {
		u, v = v, u
	}
	for k := 0; k < 17; k++ {
		if (((D[v] - D[u]) >> k) & 1) != 0 {
			res += 1 << k
			if t != 0 {
				res -= W[k][v]
			}
			v = P[k][v]
		}
	}
	if u == v {
		return res
	}
	for k := 16; k >= 0; k-- {
		if P[k][u] != P[k][v] {
			res += 2 << k
			if t != 0 {
				res -= W[k][u] + W[k][v]
			}
			u = P[k][u]
			v = P[k][v]
		}
	}
	res += 2
	if t != 0 {
		res -= W[0][u] + W[0][v]
	}
	return res
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
