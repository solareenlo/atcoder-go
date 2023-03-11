package main

import (
	"bufio"
	"fmt"
	"os"
)

const INF = int(1e18)

var n int

var X, Y [100000]int

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		X[i] = x + y
		Y[i] = x - y
	}

	l, r := 1, INF
	for r-l > 1 {
		m := (l + r) / 2
		if judge(m) {
			l = m
		} else {
			r = m
		}
	}
	fmt.Println(l)
}

func judge(m int) bool {
	x_min := INF
	x_max := -INF
	x_min_i := -1
	x_max_i := -1
	y_min := INF
	y_max := -INF
	y_min_i := -1
	y_max_i := -1
	for i := 0; i < n; i++ {
		if chmin(&x_min, X[i]) {
			x_min_i = i
		}
		if chmax(&x_max, X[i]) {
			x_max_i = i
		}
		if chmin(&y_min, Y[i]) {
			y_min_i = i
		}
		if chmax(&y_max, Y[i]) {
			y_max_i = i
		}
	}
	uf := NewUnionFind(n)
	for i := 0; i < n; i++ {
		if x_min+m <= X[i] {
			uf.Merge(i, x_min_i)
		}
		if x_max-m >= X[i] {
			uf.Merge(i, x_max_i)
		}
		if y_min+m <= Y[i] {
			uf.Merge(i, y_min_i)
		}
		if y_max-m >= Y[i] {
			uf.Merge(i, y_max_i)
		}
	}
	return (uf.Size(0) == n)
}

func chmax(a *int, b int) bool {
	if *a < b {
		*a = b
		return true
	}
	return false
}

func chmin(a *int, b int) bool {
	if *a > b {
		*a = b
		return true
	}
	return false
}

type UnionFind struct {
	root []int
	size []int
}

func NewUnionFind(size int) *UnionFind {
	uf := new(UnionFind)
	uf.root = make([]int, size)
	uf.size = make([]int, size)

	for i := 0; i < size; i++ {
		uf.root[i] = i
		uf.size[i] = 1
	}

	return uf
}

func (uf *UnionFind) Merge(p, q int) bool {
	q = uf.Root(q)
	p = uf.Root(p)

	if q == p {
		return false
	}

	if uf.size[q] < uf.size[p] {
		uf.root[q] = uf.root[p]
		uf.size[p] += uf.size[q]
	} else {
		uf.root[p] = uf.root[q]
		uf.size[q] += uf.size[p]
	}
	return true
}

func (uf *UnionFind) Root(p int) int {
	if p > len(uf.root)-1 {
		return -1
	}

	for uf.root[p] != p {
		uf.root[p] = uf.root[uf.root[p]]
		p = uf.root[p]
	}

	return p
}

func (uf *UnionFind) Same(p, q int) bool {
	return uf.Root(p) == uf.Root(q)
}

func (uf *UnionFind) Size(x int) int {
	return uf.size[uf.Root(x)]
}
