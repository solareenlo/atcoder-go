package main

import (
	"bufio"
	"fmt"
	"os"
)

const inf = 1073741723

type tuple struct {
	a, b, c, d int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	type pair struct {
		to, Cap int
	}

	var n, m int
	fmt.Fscan(in, &n, &m)
	g := make([][]int, n)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		g[a] = append(g[a], b)
	}
	var s scc
	s.init(g)
	vs := s.s
	var mf maxflow
	mf.init(vs*2 + 2)
	src := vs * 2
	dst := src + 1
	es := make([]tuple, 0)
	ini := make([]tuple, 0)
	fin := make([]tuple, 0)
	for i := 0; i < vs; i++ {
		for _, j := range s.da[i] {
			es = append(es, mf.ae(i, vs+j, 1))
			es = append(es, mf.ae(vs+i, vs+j, inf))
		}
	}
	for i := 0; i < vs; i++ {
		ini = append(ini, mf.ae(src, i, 1))
	}
	for i := 0; i < vs; i++ {
		fin = append(fin, mf.ae(vs+i, dst, 1))
	}
	mf.calc(src, dst)
	nx := make([]int, vs)
	for i := range nx {
		nx[i] = -1
	}

	h := make([][]pair, vs*2)
	for _, tmp := range es {
		u := tmp.a
		v := tmp.b
		j := tmp.d
		h[u] = append(h[u], pair{v, mf.g[v][j].Cap})
	}
	good := make([]bool, vs*2)
	for _, tmp := range fin {
		u := tmp.a
		v := tmp.b
		j := tmp.d
		if mf.g[v][j].Cap != 0 {
			good[u] = true
		}
	}
	for _, tmp := range ini {
		v := tmp.b
		j := tmp.d
		if mf.g[v][j].Cap != 0 {
			tmp := v
			for !good[v] {
				if h[v][len(h[v])-1].Cap == 0 {
					h[v] = h[v][:len(h[v])-1]
					continue
				}
				h[v][len(h[v])-1].Cap--
				v = h[v][len(h[v])-1].to
			}
			nx[tmp] = v - vs
			good[v] = false
		}
	}

	d := make([]int, vs)
	for i := range d {
		d[i] = -1
	}
	used := 0
	for i := 0; i < vs; i++ {
		if d[i] == -1 {
			d[i] = used
			used++
		}
		if nx[i] != -1 {
			d[nx[i]] = d[i]
		}
	}

	for i := 0; i < n; i++ {
		if i < n-1 {
			fmt.Fprintf(out, "%d ", d[s.bl[i]]+1)
		} else {
			fmt.Fprintln(out, d[s.bl[i]]+1)
		}
	}
}

type scc struct {
	n                int
	g                [][]int
	ord, low, bl, st []int
	s, head          int
	idx, da          [][]int
}

func (s *scc) dfs(v int) {
	s.ord[v] = s.head
	s.low[v] = s.head
	s.head++
	s.st = append(s.st, v)
	for _, to := range s.g[v] {
		if s.ord[to] == -1 {
			s.dfs(to)
			s.low[v] = min(s.low[v], s.low[to])
		} else if s.bl[to] == -1 {
			s.low[v] = min(s.low[v], s.ord[to])
		}
	}
	if s.ord[v] == s.low[v] {
		c := len(s.idx)
		s.idx = append(s.idx, []int{})
		for {
			a := s.st[len(s.st)-1]
			s.st = s.st[:len(s.st)-1]
			s.bl[a] = c
			s.idx[len(s.idx)-1] = append(s.idx[len(s.idx)-1], a)
			if v == a {
				break
			}
		}
	}
}

func (scc *scc) init(gg [][]int) {
	scc.n = len(gg)
	scc.g = make([][]int, len(gg))
	for i := range scc.g {
		scc.g[i] = make([]int, len(gg[i]))
		copy(scc.g[i], gg[i])
	}
	scc.ord = make([]int, scc.n)
	scc.low = make([]int, scc.n)
	scc.bl = make([]int, scc.n)
	for i := 0; i < scc.n; i++ {
		scc.ord[i] = -1
		scc.low[i] = -1
		scc.bl[i] = -1
	}
	scc.head = 0
	for i := 0; i < scc.n; i++ {
		if scc.ord[i] == -1 {
			scc.dfs(i)
		}
	}
	scc.s = len(scc.idx)
	for i := 0; i < scc.n; i++ {
		scc.bl[i] = scc.s - 1 - scc.bl[i]
	}
	scc.idx = reverseOrderSlice(scc.idx)
	u := make([]bool, scc.s)
	scc.da = make([][]int, scc.s)
	for i := 0; i < scc.s; i++ {
		for _, v := range scc.idx[i] {
			for _, to := range scc.g[v] {
				if scc.bl[v] < scc.bl[to] {
					if !u[scc.bl[to]] {
						scc.da[scc.bl[v]] = append(scc.da[scc.bl[v]], scc.bl[to])
					}
					u[scc.bl[to]] = true
				}
			}
		}
		for _, v := range scc.idx[i] {
			for _, to := range scc.g[v] {
				u[scc.bl[to]] = false
			}
		}
	}
}

type E struct {
	to, rev, Cap int
}

type maxflow struct {
	g          [][]E
	itr, lv, q []int
}

func (m *maxflow) init(n int) {
	m.g = make([][]E, n)
	m.itr = make([]int, n)
	m.lv = make([]int, n)
	m.q = make([]int, n)
}

func (m *maxflow) ae(s, t, c int) tuple {
	m.g[s] = append(m.g[s], E{t, len(m.g[t]), c})
	m.g[t] = append(m.g[t], E{s, len(m.g[s]) - 1, 0})
	return tuple{s, t, len(m.g[s]) - 1, len(m.g[t]) - 1}
}

func (m *maxflow) aeU(s, t, c int) {
	m.g[s] = append(m.g[s], E{t, len(m.g[t]), c})
	m.g[t] = append(m.g[t], E{s, len(m.g[s]) - 1, c})
}

func (m *maxflow) bfs(s, t int) {
	for i := range m.lv {
		m.lv[i] = -1
	}
	m.lv[s] = 0
	l := 0
	r := 0
	m.q[r] = s
	r++
	for l < r {
		v := m.q[l]
		l++
		if v == t {
			break
		}
		for _, e := range m.g[v] {
			if e.Cap > 0 && m.lv[e.to] == -1 {
				m.lv[e.to] = m.lv[v] + 1
				m.q[r] = e.to
				r++
			}
		}
	}
}

func (m *maxflow) dfs(v, t, f int) int {
	if v == t {
		return f
	}
	res := 0
	for i := &m.itr[v]; m.itr[v] < len(m.g[v]); (*i)++ {
		e := &m.g[v][*i]
		if (*e).Cap > 0 && m.lv[(*e).to] == m.lv[v]+1 {
			a := m.dfs((*e).to, t, min(f, (*e).Cap))
			if a > 0 {
				(*e).Cap -= a
				m.g[(*e).to][(*e).rev].Cap += a
				res += a
				f -= a
				if f <= 0 {
					break
				}
			}
		}
	}
	return res
}

func (m *maxflow) calc(s, t int) int {
	f := 0
	for f < inf {
		m.bfs(s, t)
		if m.lv[t] == -1 {
			return f
		}
		for i := range m.itr {
			m.itr[i] = 0
		}
		f += m.dfs(s, t, inf-f)
	}
	return f
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func reverseOrderSlice(a [][]int) [][]int {
	n := len(a)
	res := make([][]int, n)
	n = copy(res, a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
