package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 100000

var (
	z   int = 17
	id  int
	ord = make([]int, N)
	low = make([]int, N)
	D   = make([]int, N)
	B   = make([]int, N)
	G   = make([][]int, N)
	P   = make([][]int, 17)
)

func initSlice(s []int) {
	for i := range s {
		s[i] = -1
	}
}

func dfs(u, p int) {
	low[u] = id
	ord[u] = id
	id++
	for _, v := range G[u] {
		if v == p {
			continue
		}
		if ord[v] == -1 {
			dfs(v, u)
			low[u] = min(low[u], low[v])
		} else {
			low[u] = min(low[u], ord[v])
		}
	}
}

func bcc(u int) {
	for _, v := range G[u] {
		if D[v] != -1 {
			continue
		}
		P[0][v] = u
		D[v] = D[u] + 1
		B[v] = B[u]
		if ord[u] < low[v] {
			B[v]++
		}
		bcc(v)
	}
}

func dist(u, v int) int {
	if D[u] > D[v] {
		u, v = v, u
	}
	d := D[v] - D[u]
	res := 0
	for k := z - 1; k >= 0; k-- {
		if (d>>k)&1 != 0 {
			p := P[k][v]
			res += B[v] - B[p]
			v = p
		}
	}
	if u == v {
		return res
	}
	for k := z - 1; k >= 0; k-- {
		pu := P[k][u]
		pv := P[k][v]
		if pu != pv {
			res += (B[u] - B[pu]) + (B[v] - B[pv])
			u = pu
			v = pv
		}
	}
	res += (B[u] - B[P[0][u]]) + (B[v] - B[P[0][v]])
	return res
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	initSlice(ord)
	initSlice(low)
	initSlice(D)
	initSlice(B)
	for i := range P {
		P[i] = make([]int, N)
		initSlice(P[i])
	}

	var n, m int
	fmt.Fscan(in, &n, &m)

	for i := 0; i < m; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		x--
		y--
		G[x] = append(G[x], y)
		G[y] = append(G[y], x)
	}

	dfs(0, -1)
	D[0] = 0
	B[0] = 0
	bcc(0)

	for k := 0; k < z-1; k++ {
		for i := 0; i < n; i++ {
			if P[k][i] != -1 {
				P[k+1][i] = P[k][P[k][i]]
			}
		}
	}

	var q int
	fmt.Fscan(in, &q)
	for i := 0; i < q; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		a--
		b--
		c--
		if dist(a, b)+dist(b, c) == dist(a, c) {
			fmt.Fprintln(out, "OK")
		} else {
			fmt.Fprintln(out, "NG")
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
