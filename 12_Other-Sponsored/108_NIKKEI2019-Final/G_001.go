package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const SIZE = 100005

type Line struct {
	a, b int // y=ax+b
}

func (l *Line) lessThan(r Line) bool    { return l.a < r.a }
func (l *Line) greaterThan(r Line) bool { return l.a > r.a }
func (l *Line) get(x int) int           { return l.a*x + l.b }

func up(p, q, r Line) bool { // p.a<q.a<r.a
	vx := float64(q.b-p.b) / float64(p.a-q.a)
	vy := float64(r.b-p.b) / float64(p.a-r.a)
	return vx+(1e-10) < vy
}

type getmn struct {
	que [SIZE]Line
	sz  int
}

func (g *getmn) init() { g.sz = 0 }

func (g *getmn) add(l Line) { // l.a : increasing
	if g.sz >= 1 && g.que[g.sz-1].a == l.a {
		g.que[g.sz-1].b = max(g.que[g.sz-1].b, l.b)
	} else {
		for g.sz >= 2 && !up(g.que[g.sz-2], g.que[g.sz-1], l) {
			g.sz--
		}
		g.que[g.sz] = l
		g.sz++
	}
}

func (g *getmn) get(x int) int {
	if g.sz == 0 {
		return 0
	}
	l, r := 0, g.sz
	ret := 0
	for r-l > 1 {
		d := (l + r) / 2
		ret = max(ret, g.que[d].get(x))
		if d+1 == g.sz {
			r = d
		} else {
			ret = max(ret, g.que[d+1].get(x))
			if g.que[d+1].get(x) > g.que[d].get(x) {
				l = d
			} else {
				r = d
			}
		}
	}
	ret = max(ret, g.que[l].get(x))
	return ret
}

var que getmn

type edge struct {
	to, cost int
}

type P struct {
	x, y int
}

var vec [SIZE][]edge
var query [SIZE][]P
var nd [SIZE]int
var ans [SIZE]int
var use [SIZE]bool
var n, Q int

func dfs(v, p int) {
	nd[v] = 1
	for i := 0; i < len(vec[v]); i++ {
		to := vec[v][i].to
		if to != p && !use[to] {
			dfs(to, v)
			nd[v] += nd[to]
		}
	}
}

func center(v, p, all int) P {
	mx := all - nd[v]
	ret := P{all, v}
	for i := 0; i < len(vec[v]); i++ {
		to := vec[v][i].to
		if to != p && !use[to] {
			ret = minP(ret, center(to, v, all))
			mx = max(mx, nd[to])
		}
	}
	ret = minP(ret, P{mx, v})
	return ret
}

var all []Line

func Make(v, p, d, bef, sum int) {
	if bef != -1 && sum-bef*d <= 0 {
		all = append(all, Line{bef, sum - bef*d})
	}
	for i := 0; i < len(vec[v]); i++ {
		e := vec[v][i]
		if !use[e.to] && e.to != p {
			Make(e.to, v, d+1, e.cost, sum+e.cost)
		}
	}
}

func dfs2(v, p, d, sum int) {
	for i := 0; i < len(query[v]); i++ {
		L := query[v][i].x
		if L >= d {
			id := query[v][i].y
			ans[id] = max(ans[id], sum+que.get(L-d))
		}
	}
	for i := 0; i < len(vec[v]); i++ {
		e := vec[v][i]
		if !use[e.to] && e.to != p {
			dfs2(e.to, v, d+1, sum+e.cost)
		}
	}
}

func solve(v int) {
	dfs(v, -1)
	ct := center(v, -1, nd[v]).y
	que.init()
	all = make([]Line, 0)
	Make(ct, -1, 0, -1, 0)
	sort.Slice(all, func(i, j int) bool {
		return all[i].lessThan(all[j])
	})
	for i := 0; i < len(all); i++ {
		que.add(all[i])
	}
	dfs2(ct, -1, 0, 0)
	use[ct] = true
	for i := 0; i < len(vec[ct]); i++ {
		e := vec[ct][i]
		if !use[e.to] {
			solve(e.to)
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &n)
	for i := 0; i < n-1; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		a--
		b--
		vec[a] = append(vec[a], edge{b, c})
		vec[b] = append(vec[b], edge{a, c})
	}
	fmt.Fscan(in, &Q)
	for i := 0; i < Q; i++ {
		var v, l int
		fmt.Fscan(in, &v, &l)
		v--
		query[v] = append(query[v], P{l, i})
	}
	solve(0)
	for i := 0; i < Q; i++ {
		fmt.Fprintln(out, ans[i])
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minP(a, b P) P {
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
