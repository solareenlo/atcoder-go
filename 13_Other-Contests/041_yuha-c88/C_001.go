package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	m := make(map[string]int)
	ss := make([]string, 0)
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(in, &s)
		ss = append(ss, s)
		m[s] = i
	}
	uf := NewUnionFind(2 * n)
	for i := 0; i < n; i++ {
		var a, b, c, d, e string
		fmt.Fscan(in, &a, &b, &c, &d, &e)
		if d == "good" {
			uf.Merge(i, m[a])
			uf.Merge(i+n, m[a]+n)
		} else {
			uf.Merge(i, m[a]+n)
			uf.Merge(i+n, m[a])
		}
	}
	for i := 0; i < n; i++ {
		if uf.Same(i, i+n) {
			fmt.Println("No answers")
			return
		}
	}
	ans := make([]string, 0)
	for i := 0; i < n; i++ {
		if uf.Root(i) != i {
			continue
		}
		a := make([]string, 0)
		b := make([]string, 0)
		for j := 0; j < n; j++ {
			if uf.Same(i, j) {
				a = append(a, ss[j])
			}
			if uf.Same(i+n, j) {
				b = append(b, ss[j])
			}
		}
		if len(a) < len(b) {
			a, b = b, a
		} else if len(a) == len(b) {
			sort.Strings(a)
			if lessThan(b, a) {
				a, b = b, a
			}
		}
		for _, k := range a {
			ans = append(ans, k)
		}
	}
	sort.Strings(ans)
	for i := range ans {
		fmt.Println(ans[i])
	}
}

func lessThan(a, b []string) bool {
	n := len(a)
	for i := 0; i < n; i++ {
		if a[i] == b[i] {
			continue
		}
		return a[i] < b[i]
	}
	return false
}

type UnionFind struct {
	root []int
	size []int
}

func NewUnionFind(size int) *UnionFind {
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
