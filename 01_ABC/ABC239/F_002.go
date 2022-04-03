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

	d := make([]int, n)
	for i := range d {
		fmt.Fscan(in, &d[i])
	}

	dsu := NewUnionFind(n)
	for i := 0; i < m; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		x--
		y--
		d[x]--
		d[y]--
		dsu.Merge(x, y)
	}

	tmp := make([][]int, n)
	for i := 0; i < n; i++ {
		if d[i] < 0 {
			fmt.Fprintln(out, -1)
			return
		}
		for j := 0; j < d[i]; j++ {
			tmp[dsu.Root(i)] = append(tmp[dsu.Root(i)], i)
		}
	}

	c2 := make([][]int, 0)
	c1 := make([]int, 0)
	for i := 0; i < n; i++ {
		if len(tmp[i]) == 1 {
			c1 = append(c1, tmp[i][0])
		} else if len(tmp[i]) > 1 {
			c2 = append(c2, tmp[i])
		}
	}

	type pair struct{ x, y int }
	ans := make([]pair, 0)
	for _, v := range c2 {
		for i := 0; i < len(v)-1; i++ {
			if len(c1) == 0 {
				fmt.Fprintln(out, -1)
				return
			}
			dsu.Merge(v[i], c1[len(c1)-1])
			ans = append(ans, pair{v[i], c1[len(c1)-1]})
			c1 = c1[:len(c1)-1]
		}
		c1 = append(c1, v[len(v)-1])
	}

	if len(c1) != 2 {
		fmt.Fprintln(out, -1)
		return
	}
	dsu.Merge(c1[0], c1[1])
	ans = append(ans, pair{c1[0], c1[1]})

	if dsu.Size(0) != n {
		fmt.Fprintln(out, -1)
		return
	}

	for i := range ans {
		fmt.Fprintln(out, ans[i].x+1, ans[i].y+1)
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
