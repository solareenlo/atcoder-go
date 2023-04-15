package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)
	uf1 := NewUnionFind(n)
	uf2 := NewUnionFind(n)
	v := make([]tuple, 0)
	cnt := 0
	sum := 0
	for i := 0; i < m; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		a--
		b--
		if c == 0 {
			var w int
			fmt.Fscan(in, &w)
			v = append(v, tuple{a, b, w})
		} else {
			var s string
			fmt.Fscan(in, &s)
			if !uf1.Same(a, b) {
				uf1.Merge(a, b)
				cnt++
			}
		}
	}
	sort.Slice(v, func(a, b int) bool {
		return v[a].c < v[b].c
	})
	for _, s := range v {
		if !uf1.Same(s.a, s.b) {
			sum += s.c
			uf1.Merge(s.a, s.b)
			uf2.Merge(s.a, s.b)
		}
	}
	u := make([]tuple, 0)
	u = append(u, tuple{0, cnt, sum})
	for _, s := range v {
		if !uf2.Same(s.a, s.b) {
			sum += s.c
			cnt--
			uf2.Merge(s.a, s.b)
		}
		u = append(u, tuple{s.c, cnt, sum})
	}
	var q int
	fmt.Fscan(in, &q)
	for i := 0; i < q; i++ {
		var a int
		fmt.Fscan(in, &a)
		id := lowerBound(u, tuple{a, 0, 0})
		if id != 0 {
			id--
		}
		fmt.Fprintln(out, u[id].b*a+u[id].c)
	}
}

type tuple struct {
	a, b, c int
}

func lowerBound(a []tuple, x tuple) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i].a > x.a
	})
	return idx
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
