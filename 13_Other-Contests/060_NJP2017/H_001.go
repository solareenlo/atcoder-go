package main

import (
	"bufio"
	"fmt"
	"os"
)

const MAXN = 1 << 17

var num int
var in, sz, nxt, C [MAXN]int
var g [MAXN][]int
var B *Fenwick

func main() {
	IN := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(IN, &n)
	var par [MAXN]int
	for i := 1; i < n; i++ {
		fmt.Fscan(IN, &par[i])
		par[i]--
		g[par[i]] = append(g[par[i]], i)
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(IN, &C[i])
	}
	B = NewFenwick(MAXN)
	dfs_sz(0)
	dfs_hld(0)
	var q int
	fmt.Fscan(IN, &q)
	for ; q > 0; q-- {
		var t int
		fmt.Fscan(IN, &t)
		if t == 1 {
			var u int
			fmt.Fscan(IN, &u)
			u--
			B.Add(in[u], 1-2*C[u])
			C[u] ^= 1
		} else {
			var u, v int
			fmt.Fscan(IN, &u, &v)
			u--
			v--
			c := 0
			for {
				if in[u] > in[v] {
					u, v = v, u
				}
				if nxt[u] != nxt[v] {
					c += B.Sum(in[nxt[v]], in[v]+1)
					v = par[nxt[v]]
				} else {
					c += B.Sum(in[u]+1, in[v]+1)
					break
				}
			}
			if c != 0 {
				fmt.Println("NO")
			} else {
				fmt.Println("YES")
			}
		}
	}
}

func dfs_sz(u int) {
	sz[u] = 1
	for v := range g[u] {
		dfs_sz(g[u][v])
		sz[u] += sz[g[u][v]]
		if sz[g[u][v]] > sz[g[u][0]] {
			g[u][v], g[u][0] = g[u][0], g[u][v]
		}
	}
}

func dfs_hld(u int) {
	in[u] = num
	num++
	for _, v := range g[u] {
		if v == g[u][0] {
			nxt[v] = nxt[u]
		} else {
			nxt[v] = v
		}
		dfs_hld(v)
		C[v] ^= C[u] ^ 1
		B.Add(in[v], C[v])
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
