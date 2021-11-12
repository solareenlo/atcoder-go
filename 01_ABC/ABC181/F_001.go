package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	X := [120]float64{}
	Y := [120]float64{}
	for i := 0; i < n; i++ {
		fmt.Scan(&X[i], &Y[i])
	}

	l, r := 0.0, 100.0
	for i := 0; i < 50; i++ {
		m := l + r
		uf := New(n + 2)
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				if (X[j]-X[i])*(X[j]-X[i])+(Y[j]-Y[i])*(Y[j]-Y[i]) < m*m {
					uf.Merge(i, j)
				}
			}
			if Y[i]+m > 100.0 {
				uf.Merge(i, n)
			}
			if m-Y[i] > 100.0 {
				uf.Merge(i, n+1)
			}
		}
		if uf.Same(n, n+1) {
			r = m / 2.0
		} else {
			l = m / 2.0
		}
	}

	fmt.Println(l)
}

type UnionFind struct {
	root []int
	size []int
}

func New(size int) *UnionFind {
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
