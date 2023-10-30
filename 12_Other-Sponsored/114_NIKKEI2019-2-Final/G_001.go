package main

import (
	"bufio"
	"fmt"
	"os"
)

const MX = 202020

type P struct {
	x, y int
}

var root, dep, rootl, dp [MX]int
var graph [MX][]P

func dfs1(p, rt, d int) int {
	root[p] = rt
	dep[p] = d
	ans := 0
	for _, e := range graph[p] {
		q := e.x
		w := e.y
		if q == rt {
			continue
		}
		rootl[q] = w
		ans += max(0, w+dfs1(q, p, d+1))
	}
	dp[p] = ans
	return dp[p]
}

var dublin [18][202020]int

func upupup(p, l int) int {
	if l < 0 {
		os.Exit(1)
	}
	for i := 18 - 1; i >= 0; i-- {
		if (l & (1 << i)) != 0 {
			p = dublin[i][p]
		}
	}
	return p
}

func lca(p, q int) int {
	if dep[p] > dep[q] {
		return lca(q, p)
	}
	q = upupup(q, dep[q]-dep[p])
	if p == q {
		return p
	}
	for i := 18 - 1; i >= 0; i-- {
		if dublin[i][p] != dublin[i][q] {
			p = dublin[i][p]
			q = dublin[i][q]
		}
	}
	return dublin[0][p]
}

func mofmof(g, v int) int {
	return upupup(v, dep[v]-dep[g]-1)
}

var super [202020]int

func dfs2(p, sup, supe int) {
	sup += supe
	sup = max(sup, 0)
	super[p] = sup
	for _, e := range graph[p] {
		q := e.x
		w := e.y
		if q == root[p] {
			continue
		}
		dfs2(q, sup+dp[p]-max(0, w+dp[q]), w)
	}
}

var relka [202020]int

func dfs3(p, supe int) {
	if p == 0 {
		relka[p] = dp[p]
	} else {
		r := root[p]
		relka[p] = relka[r] - max(0, supe+dp[p]) + supe + dp[p]
	}
	for _, e := range graph[p] {
		q := e.x
		w := e.y
		if q == root[p] {
			continue
		}
		dfs3(q, w)
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N int
	fmt.Fscan(in, &N)
	for i := 0; i < N-1; i++ {
		var u, v, w int
		fmt.Fscan(in, &u, &v, &w)
		u--
		v--
		graph[u] = append(graph[u], P{v, w})
		graph[v] = append(graph[v], P{u, w})
	}

	dfs1(0, -1, 0)
	dfs2(0, 0, 0)
	dfs3(0, 0)

	for i := 0; i < N; i++ {
		dublin[0][i] = root[i]
	}
	for i := 1; i < 18; i++ {
		for j := 0; j < N; j++ {
			if dublin[i-1][j] == -1 {
				dublin[i][j] = -1
			} else {
				dublin[i][j] = dublin[i-1][dublin[i-1][j]]
			}
		}
	}

	var Q int
	fmt.Fscan(in, &Q)
	for Q > 0 {
		Q--
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		g := lca(u, v)
		ans := 0
		if g == u || g == v {
			if g == v {
				u, v = v, u
			}
			vv := mofmof(g, v)
			ans += relka[v] - relka[g] + max(0, rootl[vv]+dp[vv])
			ans += super[g]
			ans += dp[g] - max(0, dp[vv]+rootl[vv])
		} else {
			uu := mofmof(g, u)
			vv := mofmof(g, v)
			ans += relka[u] - relka[g] + max(0, rootl[uu]+dp[uu])
			ans += relka[v] - relka[g] + max(0, rootl[vv]+dp[vv])
			ans += super[g]
			ans += dp[g] - max(0, dp[uu]+rootl[uu]) - max(0, dp[vv]+rootl[vv])
		}
		fmt.Fprintln(out, ans)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
