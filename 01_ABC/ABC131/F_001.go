package main

import "fmt"

const max = 100000

func main() {
	var n int
	fmt.Scan(&n)

	uf := New(2 * max)
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		x--
		y--
		uf.Merge(x, y+max)
	}

	mx := map[int]int{}
	my := map[int]int{}
	for i := 0; i < max; i++ {
		mx[uf.Root(i)]++
		my[uf.Root(i+max)]++
	}

	res := 0
	for i := 0; i < 2*max; i++ {
		res += mx[i] * my[i]
	}

	fmt.Println(res - n)
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
