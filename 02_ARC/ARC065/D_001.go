package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, k, l, x, y int
	fmt.Fscan(in, &n, &k, &l)
	d, e := New(n), New(n)
	for i := 0; i < k; i++ {
		fmt.Fscan(in, &x, &y)
		x--
		y--
		d.Merge(x, y)
	}
	for i := 0; i < l; i++ {
		fmt.Fscan(in, &x, &y)
		x--
		y--
		e.Merge(x, y)
	}

	m := make(map[pair]int)
	for i := 0; i < n; i++ {
		m[pair{d.Root(i), e.Root(i)}]++
	}

	for i := 0; i < n; i++ {
		fmt.Fprint(out, m[pair{d.Root(i), e.Root(i)}])
		if i != n-1 {
			fmt.Fprint(out, " ")
		} else {
			fmt.Fprintln(out)
		}
	}
}

type pair struct {
	a, b int
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

func (uf *UnionFind) Merge(p, q int) {
	q = uf.Root(q)
	p = uf.Root(p)

	if q == p {
		return
	}

	if uf.size[q] < uf.size[p] {
		uf.root[q] = uf.root[p]
		uf.size[p] += uf.size[q]
	} else {
		uf.root[p] = uf.root[q]
		uf.size[q] += uf.size[p]
	}
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
