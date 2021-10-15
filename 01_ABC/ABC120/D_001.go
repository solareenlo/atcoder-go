package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	a := make([]int, m)
	b := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&a[i], &b[i])
		a[i]--
		b[i]--
	}

	uf := New(n)
	total := n * (n - 1) / 2
	res := make([]int, m)
	for i := m - 1; i >= 0; i-- {
		res[i] = total
		if !uf.Same(a[i], b[i]) {
			total -= uf.Size(a[i]) * uf.Size(b[i])
		}
		uf.Merge(a[i], b[i])
	}

	for i := range res {
		fmt.Println(res[i])
	}
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
