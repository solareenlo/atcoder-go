package main

import (
	"fmt"
)

func main() {
	var s, t string
	fmt.Scan(&s, &t)

	n, nt := len(s), len(t)
	s = s + s
	for len(s) < 2*nt {
		s += s
	}

	u := t + s
	z := ZAlgorithm(u)

	ok := make([]bool, n)
	for i := 0; i < n; i++ {
		if z[nt+i] >= nt {
			ok[i] = true
		}
	}

	uf := New(n)
	for i := 0; i < n; i++ {
		if ok[i] && ok[(i+nt)%n] {
			if uf.Same(i, (i+nt)%n) {
				fmt.Println(-1)
				return
			}
			uf.Merge(i, (i+nt)%n)
		}
	}

	maxi := 0
	for i := 0; i < n; i++ {
		if ok[i] {
			maxi = max(maxi, uf.size[i])
		}
	}
	fmt.Println(maxi)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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

func ZAlgorithm(s string) []int {
	n := len(s)
	z := make([]int, n)
	z[0] = n
	for i, j := 1, 0; i < n; {
		for i+j < n && s[j:j+1] == s[i+j:i+j+1] {
			j++
		}
		z[i] = j
		if j == 0 {
			i++
			continue
		}
		k := 1
		for ; i+k < n && k+z[k] < j; k++ {
			z[i+k] = z[k]
		}
		i, j = i+k, j-k
	}
	return z
}
