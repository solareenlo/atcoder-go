package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type query struct {
	Type, x, l, r, index int
}

const INF = int(1e18)

var p [18][100010]int
var dist [100010]int

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	yy := make([]int, 0)
	queries := make([]query, 0)
	queries = append(queries, query{0, -INF, -INF, INF, n})
	queries = append(queries, query{1, INF, -INF, INF, n})
	yy = append(yy, -INF)
	yy = append(yy, INF)
	for i := 0; i < n; i++ {
		var x, y, r int
		fmt.Fscan(in, &x, &y, &r)
		queries = append(queries, query{0, x + y - r, x - y - r, x - y + r, i})
		queries = append(queries, query{1, x + y + r, x - y - r, x - y + r, i})
		yy = append(yy, x-y-r)
		yy = append(yy, x-y+r)
	}
	var m int
	fmt.Fscan(in, &m)
	pos := make([][]int, m)
	for i := 0; i < m; i++ {
		pos[i] = make([]int, 2)
		var x1, y1, x2, y2 int
		fmt.Fscan(in, &x1, &y1, &x2, &y2)
		queries = append(queries, query{2, x1 + y1, x1 - y1, -1, i * 2})
		queries = append(queries, query{2, x2 + y2, x2 - y2, -1, i*2 + 1})
		yy = append(yy, x1-y1)
		yy = append(yy, x2-y2)
	}
	sort.Ints(yy)
	yy = unique(yy)
	sort.Slice(queries, func(a, b int) bool {
		if queries[a].x == queries[b].x {
			return queries[a].l < queries[b].l
		}
		return queries[a].x < queries[b].x
	})
	tmp := make([]P, len(yy)-1)
	for i := range tmp {
		tmp[i].x = INF
	}
	seg := NewLazySegTree(tmp, op, e, mp, cmp, id)
	p[0][n] = -1
	for _, q := range queries {
		Type, l, r, index := q.Type, q.l, q.r, q.index
		l = lowerBound(yy, l)
		r = lowerBound(yy, r)
		if Type == 0 {
			if index != n {
				p[0][index] = seg.Get(l).x
				dist[index] = dist[p[0][index]] + 1
			}
			seg.RangeApply(l, r, F{index})
		}
		if Type == 1 {
			seg.RangeApply(l, r, F{p[0][index]})
		}
		if Type == 2 {
			pos[index/2][index%2] = seg.Get(l).x
		}
	}
	for k := 0; k+1 < 18; k++ {
		for v := 0; v < n; v++ {
			if p[k][v] < 0 {
				p[k+1][v] = -1
			} else {
				p[k+1][v] = p[k][p[k][v]]
			}
		}
	}
	for i := 0; i < m; i++ {
		fmt.Println(dist[pos[i][0]] + dist[pos[i][1]] - 2*dist[lca(pos[i][0], pos[i][1])])
	}
}

func unique(a []int) []int {
	occurred := map[int]bool{}
	result := []int{}
	for i := range a {
		if occurred[a[i]] != true {
			occurred[a[i]] = true
			result = append(result, a[i])
		}
	}
	// sort.Ints(result)
	n := len(result)
	for i := 0; i < n; i++ {
		a[i] = result[i]
	}
	return result
}

func op(a, b P) P { return P{0} }
func e() P        { return P{0} }
func mp(f F, x P) P {
	if f.d == INF {
		return x
	}
	return P{f.d}
}
func cmp(f, g F) F {
	if f.d == INF {
		return g
	}
	return f
}
func id() F { return F{INF} }

type P struct{ x int }

type F struct{ d int }

type E func() P
type Op func(a, b P) P
type Mapping func(f F, x P) P
type Composition func(f, g F) F
type Id func() F
type Compare func(v P) bool
type LazySegTree struct {
	n           int
	size        int
	log         int
	d           []P
	lz          []F
	e           E
	op          Op
	mapping     Mapping
	composition Composition
	id          Id
}

func NewLazySegTree(v []P, op Op, e E, mapping Mapping, composition Composition, id Id) *LazySegTree {
	lseg := new(LazySegTree)
	lseg.n = len(v)
	lseg.log = lseg.ceilPow2(lseg.n)
	lseg.size = 1 << uint(lseg.log)
	lseg.d = make([]P, 2*(lseg.size+1))
	lseg.e = e
	lseg.lz = make([]F, lseg.size+1)
	lseg.op = op
	lseg.mapping = mapping
	lseg.composition = composition
	lseg.id = id
	for i := range lseg.d {
		lseg.d[i] = lseg.e()
	}
	for i := range lseg.lz {
		lseg.lz[i] = lseg.id()
	}
	for i := 0; i < lseg.n; i++ {
		lseg.d[lseg.size+i] = v[i]
	}
	for i := lseg.size - 1; i >= 1; i-- {
		lseg.update(i)
	}
	return lseg
}
func (lseg *LazySegTree) update(k int) {
	lseg.d[k] = lseg.op(lseg.d[2*k], lseg.d[2*k+1])
}
func (lseg *LazySegTree) allapply(k int, f F) {
	lseg.d[k] = lseg.mapping(f, lseg.d[k])
	if k < lseg.size {
		lseg.lz[k] = lseg.composition(f, lseg.lz[k])
	}
}
func (lseg *LazySegTree) push(k int) {
	lseg.allapply(2*k, lseg.lz[k])
	lseg.allapply(2*k+1, lseg.lz[k])
	lseg.lz[k] = lseg.id()
}
func (lseg *LazySegTree) Set(p int, x P) {
	p += lseg.size
	for i := lseg.log; i >= 1; i-- {
		lseg.push(p >> i)
	}
	lseg.d[p] = x
	for i := 1; i <= lseg.log; i++ {
		lseg.update(p >> i)
	}
}
func (lseg *LazySegTree) Get(p int) P {
	p += lseg.size
	for i := lseg.log; i >= 1; i-- {
		lseg.push(p >> i)
	}
	return lseg.d[p]
}
func (lseg *LazySegTree) Prod(l, r int) P {
	if l == r {
		return lseg.e()
	}
	l += lseg.size
	r += lseg.size
	for i := lseg.log; i >= 1; i-- {
		if (l>>i)<<i != l {
			lseg.push(l >> i)
		}
		if (r>>i)<<i != r {
			lseg.push(r >> i)
		}
	}
	sml, smr := lseg.e(), lseg.e()
	for l < r {
		if (l & 1) == 1 {
			sml = lseg.op(sml, lseg.d[l])
			l++
		}
		if (r & 1) == 1 {
			r--
			smr = lseg.op(lseg.d[r], smr)
		}
		l >>= 1
		r >>= 1
	}
	return lseg.op(sml, smr)
}
func (lseg *LazySegTree) AllProd() P {
	return lseg.d[1]
}
func (lseg *LazySegTree) Apply(p int, f F) {
	p += lseg.size
	for i := lseg.log; i >= 1; i-- {
		lseg.push(p >> i)
	}
	lseg.d[p] = lseg.mapping(f, lseg.d[p])
	for i := 1; i <= lseg.log; i++ {
		lseg.update(p >> i)
	}
}
func (lseg *LazySegTree) RangeApply(l, r int, f F) {
	if l == r {
		return
	}
	l += lseg.size
	r += lseg.size
	for i := lseg.log; i >= 1; i-- {
		if (l>>i)<<i != l {
			lseg.push(l >> i)
		}
		if (r>>i)<<i != r {
			lseg.push((r - 1) >> i)
		}
	}
	l2, r2 := l, r
	for l < r {
		if l&1 == 1 {
			lseg.allapply(l, f)
			l++
		}
		if r&1 == 1 {
			r--
			lseg.allapply(r, f)
		}
		l >>= 1
		r >>= 1
	}
	l, r = l2, r2
	for i := 1; i <= lseg.log; i++ {
		if (l>>i)<<i != l {
			lseg.update(l >> i)
		}
		if (r>>i)<<i != r {
			lseg.update((r - 1) >> i)
		}
	}
}
func (lseg *LazySegTree) MaxRight(l int, cmp Compare) int {
	if l == lseg.n {
		return lseg.n
	}
	l += lseg.size
	for i := lseg.log; i >= 1; i-- {
		lseg.push(l >> i)
	}
	sm := lseg.e()
	for {
		for l%2 == 0 {
			l >>= 1
		}
		if !cmp(lseg.op(sm, lseg.d[l])) {
			for l < lseg.size {
				lseg.push(l)
				l = 2 * l
				if cmp(lseg.op(sm, lseg.d[l])) {
					sm = lseg.op(sm, lseg.d[l])
					l++
				}
			}
			return l - lseg.size
		}
		sm = lseg.op(sm, lseg.d[l])
		l++
		if l&-l == l {
			break
		}
	}
	return lseg.n
}
func (lseg *LazySegTree) MinLeft(r int, cmp Compare) int {
	if r == 0 {
		return 0
	}
	r += lseg.size
	for i := lseg.log; i >= 1; i-- {
		lseg.push(r - 1>>i)
	}
	sm := lseg.e()
	for {
		r--
		for r > 1 && r%2 != 0 {
			r >>= 1
		}
		if !cmp(lseg.op(lseg.d[r], sm)) {
			for r < lseg.size {
				lseg.push(r)
				r = 2*r + 1
				if cmp(lseg.op(lseg.d[r], sm)) {
					sm = lseg.op(lseg.d[r], sm)
					r--
				}
			}
			return r + 1 - lseg.size
		}
		sm = lseg.op(lseg.d[r], sm)
		if r&-r == r {
			break
		}
	}
	return 0
}
func (lseg *LazySegTree) ceilPow2(n int) int {
	x := 0
	for (1 << uint(x)) < n {
		x++
	}
	return x
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

func lca(u, v int) int {
	if dist[u] > dist[v] {
		u, v = v, u
	}
	for k := 0; k < 18; k++ {
		if (((dist[v] - dist[u]) >> k) & 1) != 0 {
			v = p[k][v]
		}
		if u == v {
			return u
		}
	}
	for k := 17; k >= 0; k-- {
		if p[k][u] != p[k][v] {
			u = p[k][u]
			v = p[k][v]
		}
	}
	return p[0][u]
}
