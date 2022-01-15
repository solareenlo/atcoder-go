package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	N int
	M int
	U = [10001]int{}
	V = [10001]int{}
	C = [10001]int{}
	T = [10001]int{}
)

func check(R float64) bool {
	type pair struct {
		x float64
		y int
	}
	id := make([]pair, M)
	for i := 0; i < M; i++ {
		id[i].x = float64(C[i]) - R*float64(T[i])
		id[i].y = i
	}
	sort.Slice(id, func(i, j int) bool {
		return id[i].x < id[j].x
	})

	sum := 0.0
	P := New(N)
	for i := 0; i < M; i++ {
		jd := id[i].y
		u := U[jd]
		v := V[jd]
		if !P.Same(u, v) {
			P.Merge(u, v)
			sum += id[i].x
		} else if id[i].x < 0.0 {
			sum += id[i].x
		}
	}
	if sum <= 0.0 {
		return true
	}
	return false
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &N, &M)
	for i := 0; i < M; i++ {
		fmt.Fscan(in, &U[i], &V[i], &C[i], &T[i])
	}
	L := 0.0
	R := 1e9
	for c := 0; c < 100; c++ {
		mid := (L + R) / 2
		if check(mid) {
			R = mid
		} else {
			L = mid
		}
	}
	fmt.Println(R)
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
