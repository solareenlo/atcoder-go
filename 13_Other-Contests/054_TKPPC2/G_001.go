package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 201000

var n, m int
var tot, ind, top, cnt, T int
var l, id, dep, hv, r, q, s, f, hs, bel [N]int
var low, dfn, st, bbel [N]int
var mark, vid [N]int
var ins, valid [N]bool
var e [N][]int
var p [N]P

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &n, &m)
	for i := 1; i < m+1; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		p[i] = P{u, v}
	}
	for i := m + 1; i < 2*m; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		e[i] = append(e[i], u)
		e[i] = append(e[i], v)
	}
	HLDoT(2*m - 1)
	lst := 2*m - 1
	for i := 2*m - 1; i >= 1; i-- {
		u := id[i]
		if u == bel[u] {
			if check(l[u], r[u]) {
				for j := i; j < lst+1; j++ {
					valid[id[j]] = true
				}
			} else {
				p := i
				q := lst + 1
				for p+1 < q {
					md := (p + q) >> 1
					if check(l[id[md]], r[id[md]]) {
						q = md
					} else {
						p = md
					}
				}
				for j := i; j < p+1; j++ {
					valid[id[j]] = false
				}
				for j := q; j < lst+1; j++ {
					valid[id[j]] = true
				}
			}
			lst = i - 1
		}
	}
	for i := m + 1; i < 2*m; i++ {
		if valid[i] {
			fmt.Fprintln(out, "Possible")
		} else {
			fmt.Fprintln(out, "Impossible")
		}
	}
}

func Dfs(u, f int) {
	tot++
	l[u] = tot
	id[l[u]] = u
	dep[u] = dep[f] + 1
	if hv[u] != 0 {
		Dfs(hv[u], u)
	}
	for j := 0; j < len(e[u]); j++ {
		if e[u][j] != f && e[u][j] != hv[u] {
			Dfs(e[u][j], u)
		}
	}
	r[u] = tot
}

func HLDoT(rt int) {
	t := 1
	q[0] = rt
	for i := 0; i < t; i++ {
		u := q[i]
		for j := 0; j < len(e[u]); j++ {
			if e[u][j] != f[u] {
				f[e[u][j]] = u
				q[t] = e[u][j]
				dep[q[t]] = dep[u] + 1
				t++
			}
		}
	}
	for i := t - 1; i >= 0; i-- {
		u := q[i]
		p := f[u]
		s[u]++
		s[p] += s[u]
		if l[u] == 0 {
			l[u] = 1
		}
		if hs[p] < s[u] {
			hs[p] = s[u]
			hv[p] = u
			l[p] = l[u] + 1
		}
	}
	for i := 0; i < t; i++ {
		u := q[i]
		if bel[u] == 0 {
			bel[u] = u
		}
		if hv[u] != 0 {
			bel[hv[u]] = bel[u]
		}
	}
	Dfs(rt, 0)
}

func tarjan(u int) {
	ind++
	low[u] = ind
	dfn[u] = low[u]
	ins[u] = true
	top++
	st[top] = u
	for i := 0; i < len(e[u]); i++ {
		v := e[u][i]
		if dfn[v] == 0 {
			tarjan(v)
			low[u] = min(low[u], low[v])
		} else if ins[v] {
			low[u] = min(low[u], low[v])
		}
	}
	if dfn[u] == low[u] {
		cnt++
		for {
			bbel[st[top]] = cnt
			ins[st[top]] = false
			if st[top] == u {
				top--
				break
			}
			top--
		}
	}
}

type P struct {
	first  int
	second int
}

func check(pl, pr int) bool {
	c := make([]P, 0)
	T++
	cand := make([]int, 0)
	for i := pl; i < pr+1; i++ {
		if id[i] <= m {
			u := p[id[i]].first
			v := p[id[i]].second
			c = append(c, p[id[i]])
			if mark[abs(u)] != T {
				mark[abs(u)] = T
				vid[abs(u)] = len(cand)
				cand = append(cand, abs(u))
			}
			if mark[abs(v)] != T {
				mark[abs(v)] = T
				vid[abs(v)] = len(cand)
				cand = append(cand, abs(v))
			}
		}
	}
	z := len(cand)
	for i := 0; i < 2*z; i++ {
		e[i] = make([]int, 0)
	}
	for _, p := range c {
		u := p.first
		v := p.second
		idu := vid[abs(u)]
		idv := vid[abs(v)]
		if u > 0 {
			idu = 2 * idu
		} else {
			idu = 2*idu + 1
		}
		if v > 0 {
			idv = 2 * idv
		} else {
			idv = 2*idv + 1
		}
		e[idu] = append(e[idu], idv^1)
		e[idv] = append(e[idv], idu^1)
	}
	ind = 0
	top = 0
	cnt = 0
	for i := 0; i < 2*z; i++ {
		dfn[i] = 0
	}
	for i := 0; i < 2*z; i++ {
		if dfn[i] == 0 {
			tarjan(i)
		}
	}
	for i := 0; i < z; i++ {
		if bbel[2*i] == bbel[2*i+1] {
			return false
		}
	}
	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
