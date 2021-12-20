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

	var n, m int
	fmt.Fscan(in, &n, &m)
	v := make([][]int, n+1)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		v[a] = append(v[a], b)
	}

	d := New(n + 1)
	res := make([]int, 1)
	ans := 0
	for i := n; i >= 2; i-- {
		ans++
		for _, j := range v[i] {
			if !d.Same(i, j) {
				d.Merge(i, j)
				ans--
			}
		}
		res = append(res, ans)
	}
	res = reverseOrderInt(res)
	for i := range res {
		fmt.Fprintln(out, res[i])
	}
}

func reverseOrderInt(a []int) []int {
	n := len(a)
	res := make([]int, n)
	n = copy(res, a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
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
