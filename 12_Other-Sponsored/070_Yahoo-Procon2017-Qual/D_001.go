package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var T, D, A, Z [1 << 17]int

	var q, k int
	fmt.Fscan(in, &q, &k)
	for i := 0; i < q; i++ {
		fmt.Fscan(in, &T[i], &D[i])
		if T[i] == 1 {
			fmt.Fscan(in, &A[i])
		}
	}
	idx := make([]int, q)
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		if D[idx[i]] != D[idx[j]] {
			return D[idx[i]] < D[idx[j]]
		}
		return idx[i] < idx[j]
	})
	tmp := make([]int, q)
	seg := NewLazySegTree(tmp, op, e, mp, cmp, id)
	for i := 0; i < q; i++ {
		if T[idx[i]] == 1 {
			seg.RangeApply(idx[i], q, F{A[idx[i]], k * D[idx[i]]})
		} else {
			Z[idx[i]] = seg.Get(idx[i])
		}
	}
	for i := 0; i < q; i++ {
		if T[i] == 2 {
			fmt.Println(Z[i])
		}
	}
}

type F struct {
	a, b int
}

func op(a, b int) int   { return 0 }
func e() int            { return 0 }
func mp(f F, x int) int { return min(x+f.a, f.b) }
func cmp(f, g F) F      { return F{g.a + f.a, min(g.b+f.a, f.b)} }
func id() F             { return F{0, 1 << 60} }

type E func() int
type Op func(a, b int) int
type Mapping func(f F, x int) int
type Composition func(f, g F) F
type Id func() F
type Compare func(v int) bool
type LazySegTree struct {
	n           int
	size        int
	log         int
	d           []int
	lz          []F
	e           E
	op          Op
	mapping     Mapping
	composition Composition
	id          Id
}

func NewLazySegTree(v []int, op Op, e E, mapping Mapping, composition Composition, id Id) *LazySegTree {
	lseg := new(LazySegTree)
	lseg.n = len(v)
	lseg.log = lseg.ceilPow2(lseg.n)
	lseg.size = 1 << uint(lseg.log)
	lseg.d = make([]int, 2*(lseg.size+1))
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
func (lseg *LazySegTree) Set(p int, x int) {
	p += lseg.size
	for i := lseg.log; i >= 1; i-- {
		lseg.push(p >> i)
	}
	lseg.d[p] = x
	for i := 1; i <= lseg.log; i++ {
		lseg.update(p >> i)
	}
}
func (lseg *LazySegTree) Get(p int) int {
	p += lseg.size
	for i := lseg.log; i >= 1; i-- {
		lseg.push(p >> i)
	}
	return lseg.d[p]
}
func (lseg *LazySegTree) Prod(l, r int) int {
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
func (lseg *LazySegTree) AllProd() int {
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
