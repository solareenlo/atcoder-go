package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	cost := make([][]int, 2020)
	for i := range cost {
		cost[i] = make([]int, 2020)
		for j := range cost[i] {
			cost[i][j] = -1
		}
	}
	for i := 0; i < m; i++ {
		var l, r, c int
		fmt.Fscan(in, &l, &r, &c)
		cost[l][r] = c
		cost[r][l] = c
	}
	uf := NewUnionFind(n)

	type edge struct {
		l, r, cost int
	}
	G := make([][]edge, 2020)
	for i := range G {
		G[i] = make([]edge, 0)
	}
	var solve func(int, int) int
	solve = func(x, y int) int {
		if uf.Same(x, y) {
			return 0
		}
		xid := make([]int, 0)
		yid := make([]int, 0)
		xp := uf.Root(x)
		yp := uf.Root(y)
		for i := 0; i < n; i++ {
			v := uf.Root(i)
			if v == xp {
				xid = append(xid, i)
			}
			if v == yp {
				yid = append(yid, i)
			}
		}
		ed := make([]edge, 0)
		for i := 0; i < len(xid); i++ {
			for j := 0; j < len(yid); j++ {
				if cost[xid[i]][yid[j]] != -1 {
					ed = append(ed, edge{xid[i], yid[j], cost[xid[i]][yid[j]]})
				}
			}
		}
		for i := 0; i < len(G[xp]); i++ {
			ed = append(ed, G[xp][i])
		}
		for i := 0; i < len(G[yp]); i++ {
			ed = append(ed, G[yp][i])
		}
		sort.Slice(ed, func(i, j int) bool {
			return ed[i].cost < ed[j].cost
		})
		check := NewUnionFind(n)
		cnt := 0
		SELECT := make([]edge, 0)
		res := 0
		for i := 0; i < len(ed); i++ {
			if !check.Same(ed[i].l, ed[i].r) {
				check.Merge(ed[i].l, ed[i].r)
				res += ed[i].cost
				cnt++
				SELECT = append(SELECT, ed[i])
			}
		}
		uf.Merge(x, y)
		np := uf.Root(x)
		G[np] = make([]edge, 0)
		for i := 0; i < len(SELECT); i++ {
			G[np] = append(G[np], SELECT[i])
		}
		if (cnt + 1) != (len(xid) + len(yid)) {
			return -1
		}
		return res
	}
	var q int
	fmt.Fscan(in, &q)
	for i := 0; i < q; i++ {
		var p, q int
		fmt.Fscan(in, &p, &q)
		res := solve(p, q)
		if res == -1 {
			fmt.Println("IMPOSSIBLE")
		} else {
			fmt.Println(res)
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
