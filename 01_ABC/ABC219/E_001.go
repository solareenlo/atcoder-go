package main

import "fmt"

func main() {
	n, m, a := 6, 4, 0

	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			ni := i*m + j
			var na int
			fmt.Scan(&na)
			a |= na << ni
		}
	}

	res := 0
	for s := 0; s < 1<<16; s++ {
		if s&a == a {
			x := make([][]int, n)
			for i := range x {
				x[i] = make([]int, n)
			}
			for i := 0; i < m; i++ {
				for j := 0; j < m; j++ {
					if s>>(i*m+j)&1 != 0 {
						x[i+1][j+1] = 1
					}
				}
			}
			uf := New(n * n)
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					ni := i*n + j
					if i+1 < n && x[i][j] == x[i+1][j] {
						uf.Merge(ni, ni+n)
					}
					if j+1 < n && x[i][j] == x[i][j+1] {
						uf.Merge(ni, ni+1)
					}
				}
			}
			cnt := 0
			for i := 0; i < n*n; i++ {
				if uf.Root(i) == i {
					cnt++
				}
			}
			if cnt == 2 {
				res++
			}
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
