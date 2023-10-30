package main

import (
	"bufio"
	"fmt"
	"os"
)

type HeavyLightDecomposition struct {
	g                           [][]int
	sz, in, out, head, rev, par []int
}

func (hld *HeavyLightDecomposition) init(g [][]int) {
	hld.g = g
	hld.sz = make([]int, len(g))
	hld.in = make([]int, len(g))
	hld.out = make([]int, len(g))
	hld.head = make([]int, len(g))
	hld.rev = make([]int, len(g))
	hld.par = make([]int, len(g))
}

func (hld *HeavyLightDecomposition) dfs_sz(idx, p int) {
	hld.par[idx] = p
	hld.sz[idx] = 1
	if len(hld.g[idx]) != 0 && hld.g[idx][0] == p {
		hld.g[idx][0], hld.g[idx][len(hld.g[idx])-1] = hld.g[idx][len(hld.g[idx])-1], hld.g[idx][0]
	}
	for to := range hld.g[idx] {
		if hld.g[idx][to] == p {
			continue
		}
		hld.dfs_sz(hld.g[idx][to], idx)
		hld.sz[idx] += hld.sz[hld.g[idx][to]]
		if hld.sz[hld.g[idx][0]] < hld.sz[hld.g[idx][to]] {
			hld.g[idx][0], hld.g[idx][to] = hld.g[idx][to], hld.g[idx][0]
		}
	}
}

func (hld *HeavyLightDecomposition) dfs_hld(idx, par int, times *int) {
	hld.in[idx] = *times
	(*times)++
	hld.rev[hld.in[idx]] = idx
	for _, to := range hld.g[idx] {
		if to == par {
			continue
		}
		if hld.g[idx][0] == to {
			hld.head[to] = hld.head[idx]
		} else {
			hld.head[to] = to
		}
		hld.dfs_hld(to, idx, times)
	}
	hld.out[idx] = *times
}

func (hld *HeavyLightDecomposition) build() {
	hld.dfs_sz(0, -1)
	t := 0
	hld.dfs_hld(0, -1, &t)
}

func (hld *HeavyLightDecomposition) la(v, k int) int {
	for {
		u := hld.head[v]
		if hld.in[v]-k >= hld.in[u] {
			return hld.rev[hld.in[v]-k]
		}
		k -= hld.in[v] - hld.in[u] + 1
		v = hld.par[u]
	}
}

func (hld *HeavyLightDecomposition) lca(u, v int) int {
	for {
		if hld.in[u] > hld.in[v] {
			u, v = v, u
		}
		if hld.head[u] == hld.head[v] {
			return u
		}
		v = hld.par[hld.head[v]]
	}
}

type P struct {
	x, y int
}

type PP struct {
	x, y P
}

type QRY func(int, int) PP
type PP3 func(PP, PP) PP

func (hld *HeavyLightDecomposition) query(u, v int, ti PP, q QRY, f PP3) PP {
	l, r := ti, ti
	for {
		if hld.in[u] > hld.in[v] {
			u, v = v, u
			l, r = r, l
		}
		if hld.head[u] == hld.head[v] {
			break
		}
		l = f(q(hld.in[hld.head[v]], hld.in[v]+1), l)
		v = hld.par[hld.head[v]]
	}
	return f(f(q(hld.in[u], hld.in[v]+1), l), r)
}

func (hld *HeavyLightDecomposition) add(u, v int, q func(int, int)) {
	for {
		if hld.in[u] > hld.in[v] {
			u, v = v, u
		}
		if hld.head[u] == hld.head[v] {
			break
		}
		q(hld.in[hld.head[v]], hld.in[v]+1)
		v = hld.par[hld.head[v]]
	}
	q(hld.in[u], hld.in[v]+1)
}

type LazySegmentTree struct {
	sz   int
	data []PP
	lazy []int
	f    F
	g    G
	h    H
	e1   PP
	e0   int
}

type F func(PP, PP) PP
type G func(PP, int) PP
type H func(int, int) int

func (lst *LazySegmentTree) init(n int, f F, g G, h H, e1 PP, e0 int) {
	lst.f = f
	lst.g = g
	lst.h = h
	lst.e1 = e1
	lst.e0 = e0
	lst.sz = 1
	for lst.sz < n {
		lst.sz <<= 1
	}
	lst.data = make([]PP, 2*lst.sz-1)
	lst.lazy = make([]int, 2*lst.sz-1)
	for i := 0; i < 2*lst.sz-1; i++ {
		lst.data[i] = e1
		lst.lazy[i] = e0
	}
}

func (lst *LazySegmentTree) build(v []PP) {
	for i := 0; i < len(v); i++ {
		lst.data[i+lst.sz-1] = v[i]
	}
	for i := lst.sz - 2; i >= 0; i-- {
		lst.data[i] = lst.f(lst.data[2*i+1], lst.data[2*i+2])
	}
}

func (lst *LazySegmentTree) eval(k, l, r int) {
	if lst.lazy[k] != lst.e0 {
		lst.data[k] = lst.g(lst.data[k], lst.lazy[k])
		if k < lst.sz-1 {
			lst.lazy[2*k+1] = lst.h(lst.lazy[2*k+1], lst.lazy[k])
			lst.lazy[2*k+2] = lst.h(lst.lazy[2*k+2], lst.lazy[k])
		}
	}
	lst.lazy[k] = lst.e0
}

func (lst *LazySegmentTree) update(a, b, x, k, l, r int) {
	lst.eval(k, l, r)
	if r <= a || b <= l {
		return
	}
	if a <= l && r <= b {
		lst.lazy[k] = lst.h(lst.lazy[k], x)
		lst.eval(k, l, r)
	} else {
		lst.update(a, b, x, 2*k+1, l, (l+r)/2)
		lst.update(a, b, x, 2*k+2, (l+r)/2, r)
		lst.data[k] = lst.f(lst.data[2*k+1], lst.data[2*k+2])
	}
}
func (lst *LazySegmentTree) Update(a, b, x int) {
	lst.update(a, b, x, 0, 0, lst.sz)
}

func (lst *LazySegmentTree) find(a, b, k, l, r int) PP {
	lst.eval(k, l, r)
	if b <= l || r <= a {
		return lst.e1
	}
	if a <= l && r <= b {
		return lst.data[k]
	} else {
		return lst.f(lst.find(a, b, 2*k+1, l, (l+r)/2), lst.find(a, b, 2*k+2, (l+r)/2, r))
	}
}

func (lst *LazySegmentTree) Find(a, b int) PP {
	return lst.find(a, b, 0, 0, lst.sz)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, q int
	fmt.Fscan(in, &n, &q)
	g := make([][]int, n+1)
	var a, b [200020]int
	for i := 0; i < n-1; i++ {
		fmt.Fscan(in, &a[i], &b[i])
		g[a[i]] = append(g[a[i]], b[i])
		g[b[i]] = append(g[b[i]], a[i])
	}
	var hld HeavyLightDecomposition
	hld.init(g)
	hld.build()
	var used [200020]bool
	var f func(PP, PP) PP
	f = func(a, b PP) PP { return PP{minPair(a.x, b.x), maxPair(a.y, b.y)} }
	var g1 func(PP, int) PP
	g1 = func(a PP, x int) PP { return PP{P{a.x.x + x, a.x.y}, P{a.y.x + x, a.y.y}} }
	var h func(int, int) int
	h = func(x, y int) int { return x + y }
	const INF = 1e9 + 7
	ep := PP{P{INF, INF}, P{-1, -1}}
	var sz LazySegmentTree
	sz.init(n, f, g1, h, ep, 0)
	Init := make([]PP, n)
	for i := 0; i < n; i++ {
		Init[i] = PP{P{0, i}, P{0, i}}
	}
	sz.build(Init)
	var qry func(int, int) PP
	qry = func(i, j int) PP { return sz.Find(i, j) }
	var findlca func() int
	findlca = func() int {
		p := qry(0, n).y
		if p.x == 0 {
			return -1
		} else {
			return hld.rev[p.y]
		}
	}
	va := make([]P, q)
	for i := range va {
		va[i] = P{-1, -1}
	}
	ve := make([]P, q)
	for i := range ve {
		ve[i] = P{-1, -1}
	}
	for t := 0; t < q; t++ {
		var x int
		fmt.Fscan(in, &x)
		if !used[x] {
			used[x] = true
			lca := findlca()
			if lca == -1 {
				va[t] = P{x, x}
			} else {
				lcanew := hld.lca(lca, x)
				if lcanew != lca {
					va[t] = P{lca, x}
				} else {
					p := hld.query(0, x, ep, qry, f).x
					if p.x == 0 {
						va[t] = P{hld.par[hld.rev[p.y]], x}
					}
				}
			}
			var qrya func(int, int)
			qrya = func(i, j int) { sz.Update(i, j, 1) }
			hld.add(0, x, qrya)
		} else {
			used[x] = false
			var qrya func(int, int)
			qrya = func(i, j int) { sz.Update(i, j, -1) }
			hld.add(0, x, qrya)
			lca := findlca()
			if lca == -1 {
				ve[t] = P{x, x}
			} else {
				lcanew := hld.lca(lca, x)
				if lcanew != lca {
					ve[t] = P{lca, x}
				} else {
					p := hld.query(0, x, ep, qry, f).x
					if p.x == 0 {
						ve[t] = P{hld.par[hld.rev[p.y]], x}
					}
				}
			}
		}
	}
	var s [200020]int
	for i := 0; i < q; i++ {
		if va[i].x != -1 {
			u := va[i].x
			v := va[i].y
			for {
				if hld.in[u] > hld.in[v] {
					u, v = v, u
				}
				if hld.head[u] == hld.head[v] {
					break
				}
				l := hld.in[hld.head[v]]
				r := hld.in[v] + 1
				s[l] += q - i
				s[r] -= q - i
				v = hld.par[hld.head[v]]
			}
			l := hld.in[u] + 1
			r := hld.in[v] + 1
			s[l] += q - i
			s[r] -= q - i
		}
		if ve[i].x != -1 {
			u := ve[i].x
			v := ve[i].y
			for {
				if hld.in[u] > hld.in[v] {
					u, v = v, u
				}
				if hld.head[u] == hld.head[v] {
					break
				}
				l := hld.in[hld.head[v]]
				r := hld.in[v] + 1
				s[l] -= q - i
				s[r] += q - i
				v = hld.par[hld.head[v]]
			}
			l := hld.in[u] + 1
			r := hld.in[v] + 1
			s[l] -= q - i
			s[r] += q - i
		}
	}
	var ans [200020]int
	ss := 0
	for i := 0; i < n; i++ {
		ss += s[i]
		ans[i] = ss
	}
	for i := 0; i < n-1; i++ {
		fmt.Fprintln(out, ans[max(hld.in[a[i]], hld.in[b[i]])])
	}
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

func minPair(a, b P) P {
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

func maxPair(a, b P) P {
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
