package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const INF = 3010000000000000000

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)

	x := make([]int, n)
	y := make([]int, n)
	px := make([]int, n)
	py := make([]int, n)
	c := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x[i], &y[i], &c[i])
	}

	Init := make([]S, n)
	for i := 0; i < n; i++ {
		Init[i] = S{INF, i}
	}
	sort.Slice(Init, func(a, b int) bool {
		return x[Init[a].y] < x[Init[b].y]
	})
	for i := 0; i < n; i++ {
		px[Init[i].y] = i
	}
	tx := NewLazySegTree(Init, op, e, mapping, composition, Id)

	for i := 0; i < n; i++ {
		Init[i].y = i
	}
	sort.Slice(Init, func(a, b int) bool {
		return y[Init[a].y] < y[Init[b].y]
	})
	for i := 0; i < n; i++ {
		py[Init[i].y] = i
	}
	ty := NewLazySegTree(Init, op, e, mapping, composition, Id)

	sx := make([]int, len(x))
	copy(sx, x)
	sy := make([]int, len(y))
	copy(sy, y)
	sort.Ints(sx)
	sort.Ints(sy)

	dist := make([]int, n)
	for i := range dist {
		dist[i] = INF
	}
	tx.Set(px[0], S{0, 0})
	ty.Set(py[0], S{0, 0})
	for {
		var cn, cd int
		cxS := tx.AllProd()
		cyS := ty.AllProd()
		if cxS.x == -1 {
			cd, cn = cyS.x, cyS.y
		} else if cyS.x == -1 {
			cd, cn = cxS.x, cxS.y
		} else {
			tmp := minS(cxS, cyS)
			cd, cn = tmp.x, tmp.y
		}
		if cd == -1 || cd == INF {
			break
		}

		dist[cn] = cd
		if cn != 0 {
			dist[cn] += c[cn]
		}
		tx.Set(px[cn], S{-1, cn})
		ty.Set(py[cn], S{-1, cn})

		cx := x[cn]
		cy := y[cn]
		nd := dist[cn] + c[cn]
		l := lowerBound(sx, cx-k)
		r := upperBound(sx, cx+k)
		tx.RangeApply(l, r, F(nd))
		l = lowerBound(sy, cy-k)
		r = upperBound(sy, cy+k)
		ty.RangeApply(l, r, F(nd))
	}
	ans := dist[n-1]
	if ans == INF {
		ans = -1
	}
	fmt.Println(ans)
}

func op(a, b S) S {
	if a.x == -1 {
		return b
	}
	if b.x == -1 {
		return a
	}
	return minS(a, b)
}
func e() S { return S{-1, -1} }
func mapping(f F, x S) S {
	return S{min(x.x, int(f)), x.y}
}
func composition(f, g F) F { return F(min(int(f), int(g))) }
func Id() F                { return F(INF) }

type S struct{ x, y int }

type F int

type E func() S
type Op func(a, b S) S
type Mapping func(f F, x S) S
type Composition func(f, g F) F
type ID func() F
type Compare func(v S) bool
type LazySegTree struct {
	n           int
	size        int
	log         int
	d           []S
	lz          []F
	e           E
	op          Op
	mapping     Mapping
	composition Composition
	id          ID
}

func NewLazySegTree(v []S, op Op, e E, mapping Mapping, composition Composition, id ID) *LazySegTree {
	lseg := new(LazySegTree)
	lseg.n = len(v)
	lseg.log = lseg.ceilPow2(lseg.n)
	lseg.size = 1 << uint(lseg.log)
	lseg.d = make([]S, 2*(lseg.size+1))
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
func (lseg *LazySegTree) Set(p int, x S) {
	p += lseg.size
	for i := lseg.log; i >= 1; i-- {
		lseg.push(p >> i)
	}
	lseg.d[p] = x
	for i := 1; i <= lseg.log; i++ {
		lseg.update(p >> i)
	}
}
func (lseg *LazySegTree) Get(p int) S {
	p += lseg.size
	for i := lseg.log; i >= 1; i-- {
		lseg.push(p >> i)
	}
	return lseg.d[p]
}
func (lseg *LazySegTree) Prod(l, r int) S {
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
func (lseg *LazySegTree) AllProd() S {
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minS(a, b S) S {
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

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}
