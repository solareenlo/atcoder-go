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

	var N, M int
	fmt.Fscan(in, &N, &M)
	type tuple struct{ c, a, b int }
	Es := make([]tuple, M)
	A := make([]int, 2<<17)
	B := make([]int, 2<<17)
	C := make([]int, 2<<17)
	L := make([]int, 2<<17)
	R := make([]int, 2<<17)
	for i := 0; i < M; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		C[i] = c
		A[i] = a - 1
		B[i] = b - 1
		Es[i] = tuple{C[i], A[i], B[i]}
		L[i] = -1
		R[i] = M - 1
	}
	sort.Slice(Es, func(i, j int) bool {
		return Es[i].c < Es[j].c
	})

	var all int
	for tm := 0; tm < 20; tm++ {
		ids := make([][]int, M)
		for i := 0; i < M; i++ {
			ids[(L[i]+R[i])/2] = append(ids[(L[i]+R[i])/2], i)
		}
		uf := NewUnionFind(N)
		all = 0
		for i := 0; i < M; i++ {
			c := Es[i].c
			a := Es[i].a
			b := Es[i].b
			if !uf.Same(a, b) {
				all += c
			}
			uf.Merge(a, b)
			for _, id := range ids[i] {
				if uf.Same(A[id], B[id]) {
					R[id] = i
				} else {
					L[id] = i
				}
			}
		}
	}

	for i := 0; i < M; i++ {
		fmt.Fprintln(out, all+C[i]-Es[R[i]].c)
	}
}

type UnionFind struct {
	root []int
	size []int
	n    int
}

func NewUnionFind(size int) *UnionFind {
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
