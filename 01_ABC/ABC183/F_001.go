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

	var n, Q int
	fmt.Fscan(in, &n, &Q)

	G := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		var c int
		fmt.Fscan(in, &c)
		G[i] = make(map[int]int)
		G[i][c] = 1
	}

	uf := New(n)
	for i := 0; i < Q; i++ {
		var q, x, y int
		fmt.Fscan(in, &q, &x, &y)
		if q == 2 {
			fmt.Fprintln(out, G[uf.Root(x-1)][y])
		} else if !uf.Same(uf.Root(x-1), uf.Root(y-1)) {
			x = uf.Root(x - 1)
			y = uf.Root(y - 1)
			if uf.Size(x) < uf.Size(y) {
				x, y = y, x
			}
			for k, v := range G[y] {
				G[x][k] += v
			}
			uf.Merge(x, y)
		}
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

	if uf.Size(p) < uf.Size(q) {
		p, q = q, p
	}
	uf.root[q] = p
	uf.size[p] += uf.size[q]
	return true
}

func (uf *UnionFind) Root(p int) int {
	if uf.root[p] == p {
		return p
	}
	uf.root[p] = uf.Root(uf.root[p])
	return uf.root[p]
}

func (uf *UnionFind) Same(p, q int) bool {
	return uf.Root(p) == uf.Root(q)
}

func (uf *UnionFind) Size(x int) int {
	return uf.size[uf.Root(x)]
}
