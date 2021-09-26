package main

import (
	"fmt"
	"sort"
)

type town struct{ i, x, y int }
type edge struct{ s, t, c int }

func main() {
	var n int
	fmt.Scan(&n)

	t := make([]town, n)
	var x, y int
	for i := 0; i < n; i++ {
		fmt.Scan(&x, &y)
		t[i] = town{i, x, y}
	}
	sort.Slice(t, func(i, j int) bool {
		return t[i].x < t[j].x
	})

	e := make([]edge, 0)
	for i := 0; i < n-1; i++ {
		e = append(e, edge{t[i].i, t[i+1].i, t[i+1].x - t[i].x})
	}
	sort.Slice(t, func(i, j int) bool {
		return t[i].y < t[j].y
	})
	for i := 0; i < n-1; i++ {
		e = append(e, edge{t[i].i, t[i+1].i, t[i+1].y - t[i].y})
	}
	sort.Slice(e, func(i, j int) bool {
		return e[i].c < e[j].c
	})

	uf, res := New(n), 0
	for _, t := range e {
		if uf.Merge(t.s, t.t) {
			res += t.c
		}
	}
	fmt.Println(res)
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
