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

	var n, m, q int
	fmt.Fscan(in, &n, &m, &q)
	q += m

	type pair struct{ id, c int }
	a := make([]int, q)
	b := make([]int, q)
	c := make([]int, q)
	I := make([]int, q)
	tmp := make([]pair, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(in, &a[i], &b[i], &c[i])
		tmp[i].id = i
		tmp[i].c = c[i]
	}
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].c < tmp[j].c
	})
	for i := 0; i < q; i++ {
		I[i] = tmp[i].id
	}

	u := NewUnionFind(n + 1)
	for _, i := range I {
		if i < m {
			u.Merge(a[i], b[i])
		} else {
			if u.Same(a[i], b[i]) {
				c[i] = 1
			} else {
				c[i] = 0
			}
		}
	}

	for i := m; i < q; i++ {
		if c[i] != 0 {
			fmt.Fprintln(out, "No")
		} else {
			fmt.Fprintln(out, "Yes")
		}
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
