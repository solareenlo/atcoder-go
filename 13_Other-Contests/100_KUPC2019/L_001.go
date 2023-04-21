package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const nmax = 200010
const inf = 3074457345618258602

type pi struct {
	a, b int
}

var x, y, par, sub, head, idx [nmax]int
var gg [nmax][]int
var xybuf [nmax][]pi
var seg [nmax]segtree
var uf unionfind

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var q int
	fmt.Fscan(in, &q, &x[0], &y[0])
	for i := range par {
		par[i] = -1
	}
	uf.init(nmax)
	rebuild(0)
	cur, n := 0, 1
	for j := 0; j < q; j++ {
		var t int
		fmt.Fscan(in, &t)
		if t == 1 {
			var v, a, b int
			fmt.Fscan(in, &v, &a, &b)
			i := n
			n++
			par[i] = v
			x[i] = a - b*cur
			y[i] = b
			gg[v] = append(gg[v], i)
			rebuild(i)
		} else if t == 2 {
			var a int
			fmt.Fscan(in, &a)
			cur += a
		} else if t == 3 {
			var v int
			fmt.Fscan(in, &v)
			vs := make([]int, 0)
			for v != -1 {
				vs = append(vs, v)
				v = par[head[v]]
			}
			vs = reverseOrderInt(vs)
			z := -inf
			ans := 0
			for _, w := range vs {
				seg[head[w]].ask(0, idx[w]+1, cur, &z, &ans)
			}
			if ans >= 31 {
				fmt.Fprintln(out, "many")
				out.Flush()
			} else {
				fmt.Fprintln(out, ans)
				out.Flush()
			}
		}
	}
}

func rebuild(v int) {
	vs := make([]int, 0)
	dfs1(v, &vs)
	if par[v] != -1 {
		p := uf.find(par[v])
		if sub[p] <= 2*sub[v] {
			uf.unite(p, v)
			rebuild(p)
			return
		}
	}
	dfs2(v, -1)
	for _, z := range vs {
		if len(xybuf[z]) != 0 {
			seg[z].init(xybuf[z])
			xybuf[z] = make([]pi, 0)
		}
	}
}

func dfs1(v int, dst *[]int) int {
	sub[v] = 1
	*dst = append(*dst, v)
	for e := range gg[v] {
		sub[v] += dfs1(gg[v][e], dst)
		if sub[gg[v][0]] < sub[gg[v][e]] {
			gg[v][0], gg[v][e] = gg[v][e], gg[v][0]
		}
	}
	return sub[v]
}

func dfs2(v, h int) {
	if h == -1 {
		h = v
	}
	idx[v] = len(xybuf[h])
	head[v] = h
	xybuf[h] = append(xybuf[h], pi{y[v], x[v]})
	for _, to := range gg[v] {
		if to == gg[v][0] {
			dfs2(to, h)
		} else {
			dfs2(to, -1)
		}
	}
}

type Line struct {
	a, b int
}

func (l Line) Eval(x int) int {
	return l.a*x + l.b
}

type ConvexHull struct {
	ls   []Line
	head int
}

func (ch *ConvexHull) push_back(z Line) {
	if len(ch.ls) > 0 && ch.ls[len(ch.ls)-1].a == z.a {
		z.b = max(z.b, ch.ls[len(ch.ls)-1].b)
		ch.ls = ch.ls[:len(ch.ls)-1]
	}
	for len(ch.ls) >= 2 {
		s := len(ch.ls)
		if cmpline(ch.ls[s-2], ch.ls[s-1], z) {
			break
		}
		ch.ls = ch.ls[:len(ch.ls)-1]
	}
	ch.head = min(ch.head, len(ch.ls))
	ch.ls = append(ch.ls, z)
}

func (ch ConvexHull) get(x int) int {
	lw := 0
	up := len(ch.ls)
	for up-lw > 1 {
		mid := (lw + up) / 2
		if ch.ls[mid-1].Eval(x) < ch.ls[mid].Eval(x) {
			lw = mid
		} else {
			up = mid
		}
	}
	return ch.ls[lw].Eval(x)
}

func div(a, b int) int {
	if (a^b) < 0 && a%b != 0 {
		return a/b - 1
	}
	return a / b
}

func cmpline(a, b, c Line) bool {
	ay := a.b - b.b
	ax := b.a - a.a
	by := b.b - c.b
	bx := c.a - b.a
	return div(ay, ax) < div(by, bx)
}

type segtree struct {
	s   int
	buf []ConvexHull
}

func (s *segtree) initdfs(i, l, r int, raw []pi) {
	n := len(raw)
	if i >= s.s {
		if l < n {
			s.buf[i].push_back(Line{raw[l].a, raw[l].b})
		}
	} else {
		m := (l + r) / 2
		s.initdfs(i*2, l, m, raw)
		s.initdfs(i*2+1, m, r, raw)
		if r <= n {
			sortPair(raw[l:r])
		} else if l <= n {
			sortPair(raw[l:n])
		}
		for j := min(l, n); j < min(r, n); j++ {
			s.buf[i].push_back(Line{raw[j].a, raw[j].b})
		}
	}
}

func (s *segtree) init(raw []pi) {
	n := len(raw)
	s.s = 1
	for s.s < n {
		s.s *= 2
	}
	s.buf = make([]ConvexHull, s.s*2)
	s.initdfs(1, 0, s.s, raw)
}

func (s segtree) ask1(i, l, r, b, e, x int, cur, dst *int) {
	if *dst >= 31 {
		return
	}
	if e <= l || r <= b {
		return
	}
	w := s.buf[i].get(x)
	if w <= *cur {
		return
	}
	if l+1 == r {
		(*dst)++
		*cur = w
	} else {
		m := (l + r) / 2
		s.ask1(i*2, l, m, b, e, x, cur, dst)
		s.ask1(i*2+1, m, r, b, e, x, cur, dst)
	}
}

func (s segtree) ask(b, e, x int, cur, dst *int) {
	if *dst >= 31 {
		return
	}
	s.ask1(1, 0, s.s, b, e, x, cur, dst)
}

type unionfind struct {
	p, s [nmax]int
	c    int
}

func (u *unionfind) init(n int) {
	for i := 0; i < n; i++ {
		u.p[i] = -1
		u.s[i] = 1
	}
	u.c = n
}

func (u *unionfind) find(a int) int {
	if u.p[a] == -1 {
		return a
	}
	u.p[a] = u.find(u.p[a])
	return u.p[a]
}

func (u *unionfind) unite(a, b int) bool {
	a = u.find(a)
	b = u.find(b)
	if a == b {
		return false
	}
	u.p[b] = a
	u.s[a] += u.s[b]
	u.c--
	return true
}

func (u unionfind) same(a, b int) bool {
	return u.find(a) == u.find(b)
}

func (u unionfind) sz(a int) int {
	return u.s[u.find(a)]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
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

func sortPair(tmp []pi) {
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].a == tmp[j].a {
			return tmp[i].b < tmp[j].b
		}
		return tmp[i].a < tmp[j].a
	})
}
