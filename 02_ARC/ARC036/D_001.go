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

	var N, Q int
	fmt.Fscan(in, &N, &Q)

	D := New(3 * N)
	for i := 0; i < Q; i++ {
		var w, x, y, z int
		fmt.Fscan(in, &w, &x, &y, &z)
		if w-2 != 0 {
			D.Merge(x, y+z%2*N)
			D.Merge(x+N, y-^z%2*N)
		} else if D.Same(x, y) {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}

type UnionFind struct {
	root []int
	size []int
	n    int
}

func New(size int) *UnionFind {
	uf := new(UnionFind)
	uf.root = make([]int, size)
	uf.size = make([]int, size)
	uf.n = size

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

func (uf UnionFind) Groups() [][]int {
	rootBuf, groupSize := make([]int, uf.n), make([]int, uf.n)
	for i := 0; i < uf.n; i++ {
		rootBuf[i] = uf.Root(i)
		groupSize[rootBuf[i]]++
	}
	res := make([][]int, uf.n)
	for i := 0; i < uf.n; i++ {
		res[i] = make([]int, 0, groupSize[i])
	}
	for i := 0; i < uf.n; i++ {
		res[rootBuf[i]] = append(res[rootBuf[i]], i)
	}
	result := make([][]int, 0, uf.n)
	for i := 0; i < uf.n; i++ {
		if len(res[i]) != 0 {
			r := make([]int, len(res[i]))
			copy(r, res[i])
			result = append(result, r)
		}
	}
	return result
}
