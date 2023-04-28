package main

import (
	"bufio"
	"fmt"
	"os"
)

const INF = 4611686018427387902
const INF32 = 1073741822

type edge struct {
	to, cost int
}

func main() {
	IN := bufio.NewReader(os.Stdin)
	OUT := bufio.NewWriter(os.Stdout)
	defer OUT.Flush()

	type tuple struct {
		x, y, z int
	}

	var n, q int
	fmt.Fscan(IN, &n, &q)
	g := make([][]edge, n)
	a := make([]int, n-1)
	b := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(IN, &a[i], &b[i])
		a[i]--
		b[i]--
		var c int
		fmt.Fscan(IN, &c)
		g[a[i]] = append(g[a[i]], edge{b[i], c})
		g[b[i]] = append(g[b[i]], edge{a[i], c})
	}

	in := make([]int, n)
	out := make([]int, n)
	par := make([]int, n)
	lines := make([]tuple, n)
	lines[0] = tuple{INF, INF, 0}
	time := 0
	var dfs func(int, int, int, int)
	dfs = func(x, p, cnt, sum int) {
		in[x] = time
		time++
		par[x] = p
		for _, e := range g[x] {
			if e.to == p {
				continue
			}
			lines[e.to] = tuple{e.cost, sum - e.cost*cnt, cnt}
			dfs(e.to, x, cnt+1, sum+e.cost)
		}
		out[x] = time
	}
	dfs(0, -1, 0, 0)

	v := make([]int, q)
	k := make([]int, q)
	ans := make([]int, q)
	for i := range ans {
		ans[i] = INF
	}
	for i := 0; i < q; i++ {
		var r int
		fmt.Fscan(IN, &r)
		r--
		fmt.Fscan(IN, &k[i])
		x := a[r]
		y := b[r]
		if par[x] != y {
			x, y = y, x
		}
		v[i] = x
	}

	inv := make([]int, n)
	for i := 0; i < n; i++ {
		inv[in[i]] = i
	}

	var cht LiChaoTree
	cht.init(0, INF32)
	w := make([][]int, n)
	for i := 0; i < q; i++ {
		w[in[v[i]]] = append(w[in[v[i]]], i)
	}
	for i := 0; i < n; i++ {
		if i > 0 {
			for _, j := range w[i] {
				ans[j] = min(ans[j], cht.get(k[j]))
			}
			a := lines[inv[i]].x
			b := lines[inv[i]].y
			c := lines[inv[i]].z
			cht.add_segment(a, b, c, INF32)
		}
	}
	cht.init(0, INF32)
	w = make([][]int, n+1)
	for i := 0; i < q; i++ {
		w[out[v[i]]] = append(w[out[v[i]]], i)
	}
	for i := n - 1; i >= 0; i-- {
		if i > 0 {
			a := lines[inv[i]].x
			b := lines[inv[i]].y
			c := lines[inv[i]].z
			cht.add_segment(a, b, c, INF32)
			for _, j := range w[i] {
				ans[j] = min(ans[j], cht.get(k[j]))
			}
		}
	}

	for i := 0; i < q; i++ {
		if ans[i] < INF {
			fmt.Fprintln(OUT, ans[i])
		} else {
			fmt.Fprintln(OUT, -1)
		}
	}
}

type Segment struct {
	a, b, l, r int
}

func Comp(a, b Segment, p int) bool {
	return a.Point(p) < b.Point(p)
}

func (s Segment) Point(x int) int {
	return s.a*x + s.b
}

func (s Segment) Len() int {
	return s.r - s.l
}

type Node struct {
	s        Segment
	lch, rch int
}

type LiChaoTree struct {
	v      []Node
	root   int
	xl, xr int
}

func (lct *LiChaoTree) init(xl, xr int) {
	lct.xl = xl
	lct.xr = xr
	lct.v = make([]Node, 0)
	lct.root = -1
}

func (lct *LiChaoTree) add_line(a, b int) {
	lct.add_segment(a, b, lct.xl, lct.xr)
}

func (lct *LiChaoTree) add_segment(a, b, L, R int) { // y=a*x+b, L <= x < R
	lct.add_segment1(&lct.root, Segment{a, b, L, R}, lct.xl, lct.xr)
}

func (lct *LiChaoTree) add_segment1(_i *int, s Segment, l, r int) {
	if s.r <= l || r <= s.l {
		return
	}
	s.l = clamp(s.l, l, r)
	s.r = clamp(s.r, l, r)
	if *_i < 0 {
		*_i = len(lct.v)
		lct.v = append(lct.v, Node{s, -1, -1})
		return
	}
	i := *_i
	m := l + (r-l)/2
	if l == m {
		if Comp(s, lct.v[i].s, m) {
			lct.v[i].s, s = s, lct.v[i].s
		}
		return
	}
	if lct.v[i].s.Len() == r-l && s.Len() == r-l {
		L := Comp(s, lct.v[i].s, l)
		R := Comp(s, lct.v[i].s, r)
		M := Comp(s, lct.v[i].s, m)
		if (L && M && R) || (L && M && !R) || (L && !M && R) || (!L && M && R) {
			lct.v[i].s, s = s, lct.v[i].s
		}
		if L == R {
			return
		}
		if L != M {
			lct.add_segment1(&lct.v[i].lch, s, l, m)
			return
		}
		if R != M {
			lct.add_segment1(&lct.v[i].rch, s, m, r)
			return
		}
	}
	if lct.v[i].s.Len() < s.Len() {
		lct.v[i].s, s = s, lct.v[i].s
	}
	if lct.v[i].s.l <= s.l && s.r <= lct.v[i].s.r && !Comp(s, lct.v[i].s, s.l) && !Comp(s, lct.v[i].s, s.r) {
		return
	}
	lct.add_segment1(&lct.v[i].lch, s, l, m)
	lct.add_segment1(&lct.v[i].rch, s, m, r)
}

func (lct LiChaoTree) get(x int) int {
	l := lct.xl
	r := lct.xr
	i := lct.root
	ans := INF
	for i >= 0 {
		if lct.v[i].s.l <= x && x < lct.v[i].s.r {
			ans = min(ans, lct.v[i].s.Point(x))
		}
		m := l + (r-l)/2
		if x < m {
			r = m
			i = lct.v[i].lch
		} else {
			l = m
			i = lct.v[i].rch
		}
	}
	return ans
}

func clamp(v, low, high int) int {
	return min(max(v, low), high)
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
