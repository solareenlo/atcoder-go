package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const INF = 1001001001

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, q int
	fmt.Fscan(in, &n, &q)
	var g Graph
	g.init0(n)
	for i := 0; i < n-1; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		g.add(a, b)
	}
	g.init(0)
	for qi := 0; qi < q; qi++ {
		var v, d int
		fmt.Fscan(in, &v, &d)
		v--
		if d == 1 {
			fmt.Println(1)
			continue
		}
		if d == 2 {
			fmt.Println(max(2, g.corner(v, 2)))
			continue
		}
		ans := g.get(v, d)
		fmt.Println(max(ans, g.corner(v, d)))
		continue
	}
}

type Graph struct {
	n, bi       int
	dep         []int
	to, pa, Len [][]int
	vs          []int
	lst         []map[int]struct{}
}

func (G *Graph) init0(n int) {
	G.n = n
	G.dep = make([]int, n)
	G.to = make([][]int, n)
}

func (G *Graph) add(a, b int) {
	G.to[a] = append(G.to[a], b)
	G.to[b] = append(G.to[b], a)
}

func (G *Graph) dfs(v, ndep, p int) {
	G.dep[v] = ndep
	G.pa[0][v] = p
	for i := 0; i < len(G.to[v]); i++ {
		u := G.to[v][i]
		if u == p {
			continue
		}
		G.dfs(u, ndep+1, v)
	}
}

func (G *Graph) pfs(v, nd, p int) pair {
	res := pair{nd, v}
	for i := 0; i < len(G.to[v]); i++ {
		u := G.to[v][i]
		if u == p {
			continue
		}
		res = maxPair(res, G.pfs(u, nd+1, v))
	}
	return res
}

func (G *Graph) vfs(v, tv, p int) bool {
	if v == tv {
		G.vs = append(G.vs, v)
		return true
	}
	for i := 0; i < len(G.to[v]); i++ {
		u := G.to[v][i]
		if u == p {
			continue
		}
		if G.vfs(u, tv, v) {
			G.vs = append(G.vs, v)
			return true
		}
	}
	return false
}

func (G *Graph) lfs(v, p int) int {
	ls := []int{0, 0}
	if p != -1 {
		ls = append(ls, INF)
	}
	for i := 0; i < len(G.to[v]); i++ {
		u := G.to[v][i]
		if u == p {
			continue
		}
		ls = append(ls, G.lfs(u, v))
		G.lst[v][ls[len(ls)-1]] = struct{}{}
	}
	sort.Slice(ls, func(i, j int) bool {
		return ls[i] > ls[j]
	})
	G.Len[0][v] = ls[1] + G.dep[v]
	return ls[1] + 1
}

func (G *Graph) init(root int) {
	a := G.pfs(0, 0, -1).y
	b := G.pfs(a, 0, -1).y
	G.vfs(a, b, -1)
	root = G.vs[len(G.vs)/2]

	G.bi = 0
	for (1 << G.bi) <= G.n {
		G.bi++
	}
	G.Len = make([][]int, G.bi)
	G.pa = make([][]int, G.bi)
	for i := range G.pa {
		G.pa[i] = make([]int, G.n)
		G.Len[i] = make([]int, G.n)
		for j := range G.pa[i] {
			G.pa[i][j] = -1
			G.Len[i][j] = -1
		}
	}
	G.dfs(root, 0, -1)
	G.lst = make([]map[int]struct{}, G.n)
	for i := range G.lst {
		G.lst[i] = make(map[int]struct{})
	}
	G.lfs(root, -1)
	for i := 0; i < G.bi-1; i++ {
		for j := 0; j < G.n; j++ {
			if G.pa[i][j] == -1 {
				G.pa[i+1][j] = -1
			} else {
				G.pa[i+1][j] = G.pa[i][G.pa[i][j]]
			}
		}
	}
	for i := 0; i < G.bi-1; i++ {
		for j := 0; j < G.n; j++ {
			if G.pa[i][j] == -1 {
				G.Len[i+1][j] = max(G.Len[i][j], -1)
			} else {
				G.Len[i+1][j] = max(G.Len[i][j], G.Len[i][G.pa[i][j]])
			}
		}
	}
}

func (G Graph) get(v, d int) int {
	a := v
	x := (d - 1) / 2
	res := 0
	for i := G.bi - 1; a != -1 && i >= 0; i-- {
		if 1<<i <= x {
			res = max(res, G.Len[i][a])
			a = G.pa[i][a]
			x -= 1 << i
		}
	}
	return res - G.dep[v] + d
}

func (G *Graph) corner(v, d int) int {
	d -= 2
	if (d & 1) != 0 {
		return 0
	}
	d /= 2
	if G.dep[v] < d {
		return 0
	}
	a := v
	x := d
	for i := G.bi - 1; a != -1 && i >= 0; i-- {
		if 1<<i <= x {
			a = G.pa[i][a]
			x -= 1 << i
		}
	}
	if _, ok := G.lst[a][d+1]; ok {
		return d + d + 3
	}
	return 0
}

type pair struct {
	x, y int
}

func maxPair(a, b pair) pair {
	if a.x == b.x {
		if a.y > b.y {
			return a
		}
		return b
	}
	if a.x > b.x {
		return a
	}
	return b
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
