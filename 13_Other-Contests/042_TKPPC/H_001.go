package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type DSU struct {
	p []int
}

func NewDSU(n int) *DSU {
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}
	return &DSU{p}
}

func (dsu *DSU) findSet(x int) int {
	if dsu.p[x] == x {
		return x
	}
	dsu.p[x] = dsu.findSet(dsu.p[x])
	return dsu.p[x]
}

func (dsu *DSU) merge(x, y int) bool {
	x, y = dsu.findSet(x), dsu.findSet(y)
	if x == y {
		return false
	}
	dsu.p[x] = y
	return true
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m, q, p int
	fmt.Fscan(in, &n, &m, &q, &p)

	a := make([]int, m)
	b := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &a[i], &b[i])
		a[i]--
		b[i]--
	}
	del := make([]bool, m)
	d := make([]int, q)
	g := make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(in, &d[i], &g[i])
		d[i]--
		del[d[i]] = true
	}

	dsu := NewDSU(n)
	for i := 0; i < m; i++ {
		if !del[i] {
			dsu.merge(a[i], b[i])
		}
	}
	var ans int64 = 0
	var bad = make([]bool, q)
	for i := q - 1; i >= 0; i-- {
		if dsu.merge(a[d[i]], b[d[i]]) {
			bad[i] = true
		}
	}

	deg := make([]int, n)
	adj := make([][]int, n)
	for i := 0; i < m; i++ {
		deg[a[i]]++
		deg[b[i]]++
		adj[a[i]] = append(adj[a[i]], b[i])
		adj[b[i]] = append(adj[b[i]], a[i])
	}

	for i := 0; i < n; i++ {
		sort.Ints(adj[i])
	}

	x := minIndex(deg)
	dsu = NewDSU(n)
	for i := 0; i < m; i++ {
		if !del[i] {
			dsu.merge(a[i], b[i])
		}
	}

	for i := 0; i < n; i++ {
		if !binarySearch(adj[x], i) {
			dsu.merge(x, i)
		}
	}

	for _, y := range adj[x] {
		for i := 0; i < n; i++ {
			if !binarySearch(adj[y], i) {
				dsu.merge(y, i)
			}
		}
	}

	var cand []int
	var order []int
	for i := 0; i < q; i++ {
		if bad[i] {
			order = append(order, i)
		}
	}
	sort.Slice(order, func(i, j int) bool {
		return g[order[i]] < g[order[j]]
	})
	for _, i := range order {
		if dsu.merge(a[d[i]], b[d[i]]) {
			ans += int64(g[i])
		} else {
			cand = append(cand, g[i])
		}
	}
	sort.Ints(cand)
	for i := 0; i+p < len(cand); i++ {
		ans += int64(cand[i])
	}
	fmt.Println(ans)
}

func minIndex(a []int) int {
	idx := 0
	for i := 1; i < len(a); i++ {
		if a[i] < a[idx] {
			idx = i
		}
	}
	return idx
}

func binarySearch(a []int, x int) bool {
	i := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return i < len(a) && a[i] == x
}
