package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const INT_MAX = 2147483647
const maxn = 5555
const maxm = 9999

type pair struct {
	x, y int
}

var ecnt, ccnt int
var e [maxm + 5]pair
var g [maxn + 5][]int
var ontree [maxm + 5]bool
var dep [maxn + 5]int
var col, bot [maxm + 5]int
var cl [maxn + 5][]int
var vis [maxn + 5]bool
var siz [maxn + 5]int

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 1; i <= m; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		ecnt++
		e[ecnt] = pair{u, v}
		g[u] = append(g[u], ecnt)
		g[v] = append(g[v], ecnt)
	}
	dfs0(1, 0)
	for i := 1; i <= m; i++ {
		if (ontree[i] && dep[e[i].x] > dep[e[i].y]) || (!ontree[i] && dep[e[i].y] > dep[e[i].x]) {
			e[i].x, e[i].y = e[i].y, e[i].x
		}
	}
	dfs1(1, 0)
	for i := 1; i <= m; i++ {
		if !ontree[i] {
			ccnt++
			col[i] = ccnt
		}
		fmt.Printf("%d ", col[i])
	}
	fmt.Printf("\n")
	for i := 1; i <= m; i++ {
		if ontree[i] {
			var u int
			if dep[e[bot[i]].x] > dep[e[bot[i]].y] {
				u = e[bot[i]].x
			} else {
				u = e[bot[i]].y
			}
			fmt.Printf("%d ", len(cl[u])+1)
			for _, c := range cl[u] {
				fmt.Printf("%d ", c)
			}
			fmt.Printf("%d", col[bot[i]])
		} else {
			var u int
			if dep[e[i].x] > dep[e[i].y] {
				u = e[i].x
			} else {
				u = e[i].y
			}
			fmt.Printf("%d ", len(cl[u]))
			for _, c := range cl[u] {
				fmt.Printf("%d ", c)
			}
		}
		fmt.Printf("\n")
	}
}

func dfs0(u, d int) {
	vis[u] = true
	dep[u] = d
	siz[u] = 1
	for _, eid := range g[u] {
		var v int
		if e[eid].x == u {
			v = e[eid].y
		} else {
			v = e[eid].x
		}
		if vis[v] {
			continue
		}
		ontree[eid] = true
		dfs0(v, d+1)
		siz[u] += siz[v]
	}
}

func dfs1(u, c int) pair {
	back := pair{INT_MAX, 0}
	son := make([]int, 0)
	for _, eid := range g[u] {
		v := e[eid].y
		if u == v {
			continue
		}
		if ontree[eid] {
			son = append(son, eid)
		} else {
			back = minPair(back, pair{dep[v], eid})
		}
	}
	sort.Slice(son, func(x, y int) bool {
		return siz[e[son[x]].y] > siz[e[son[y]].y]
	})
	pts := 0
	for _, eid := range son {
		v := e[eid].y
		cl[v] = make([]int, len(cl[u]))
		copy(cl[v], cl[u])
		if pts < len(cl[u]) && cl[u][pts] == c {
			pts++
		}
		if pts < len(cl[u]) {
			col[eid] = cl[u][pts]
			pts++
		} else {
			ccnt++
			cl[v] = append(cl[v], ccnt)
			col[eid] = ccnt
		}
		res := dfs1(v, col[eid])
		bot[eid] = res.y
		back = minPair(back, res)
	}
	return back
}

func minPair(a, b pair) pair {
	if a.x == b.x {
		if a.y < b.y {
			return a
		}
		return b
	}
	if a.x < b.x {
		return a
	}
	return b
}
