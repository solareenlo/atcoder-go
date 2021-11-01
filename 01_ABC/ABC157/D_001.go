package main

import "fmt"

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)

	res := make([]int, n)
	uf := New(n)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		a--
		b--
		uf.Merge(a, b)
		// res[a]--
		// res[b]--
		res[a]++
		res[b]++
	}

	for i := 0; i < k; i++ {
		var c, d int
		fmt.Scan(&c, &d)
		c--
		d--
		if uf.Same(c, d) {
			// res[c]--
			// res[d]--
			res[c]++
			res[d]++
		}
	}

	for i := 0; i < n; i++ {
		fmt.Print(uf.size[i])
	}
	fmt.Println()

	for i := 0; i < n; i++ {
		fmt.Print(res[i] + uf.size[i] - 1)
		if i != n-1 {
			fmt.Printf(" ")
		}
	}
	fmt.Println()
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
