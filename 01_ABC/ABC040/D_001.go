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
	fmt.Scan(&n, &m)

	var a, b, y int
	road := make([][3]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &a, &b, &y)
		road[i] = [3]int{a - 1, b - 1, y}
	}
	sort.Slice(road, func(i, j int) bool {
		return road[i][2] > road[j][2]
	})

	var Q, v, w int
	fmt.Fscan(in, &Q)
	query := make([][3]int, Q)
	for i := 0; i < Q; i++ {
		fmt.Fscan(in, &v, &w)
		query[i] = [3]int{v - 1, w, i}
	}
	sort.Slice(query, func(i, j int) bool {
		return query[i][1] > query[j][1]
	})

	uf := New(n)
	res := make([]int, Q)
	now := 0
	for i := 0; i < Q; i++ {
		for now < m && road[now][2] > query[i][1] {
			uf.Merge(road[now][0], road[now][1])
			now++
		}
		res[query[i][2]] = uf.Size(query[i][0])
	}

	for i := 0; i < Q; i++ {
		fmt.Fprintln(out, res[i])
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
