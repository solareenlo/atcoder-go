package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	type tuple4 struct {
		a, b, c, d int
	}

	var L [2][2][]tuple4

	var n int
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		var h, w, r, c int
		fmt.Fscan(in, &h, &w, &r, &c)
		r--
		c--
		L[r%2][c%2] = append(L[r%2][c%2], tuple4{r / 2, r/2 + (h+1)/2, c / 2, c/2 + (w+1)/2})
		L[r%2^1][c%2^1] = append(L[r%2^1][c%2^1], tuple4{(r + 1) / 2, (r+1)/2 + h/2, (c + 1) / 2, (c+1)/2 + w/2})
	}
	ans := 0
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if len(L[i][j]) == 0 {
				continue
			}
			A := make([]int, 0)
			Q := make([]tuple, 0)
			for _, tmp := range L[i][j] {
				a, b, c, d := tmp.a, tmp.b, tmp.c, tmp.d
				A = append(A, c)
				A = append(A, d)
				Q = append(Q, tuple{a, c, d})
				Q = append(Q, tuple{b, c, d})
			}
			sort.Ints(A)
			A = unique(A)
			z := len(A) - 1
			B := make([]P, z)
			for k := 0; k < z; k++ {
				B[k] = P{0, A[k+1] - A[k]}
			}
			seg := NewLazySegTree(B, op, e, mp, cmp, id)
			t := 0
			sortTuple(Q)
			for _, tmp := range Q {
				k, c, d := tmp.x, tmp.y, tmp.z
				ans += seg.AllProd().x * (k - t)
				idx0 := lowerBound(A, c)
				idx1 := lowerBound(A, d)
				seg.RangeApply(idx0, idx1, 1)
				t = k
			}
		}
	}
	fmt.Println(ans)
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

type tuple struct {
	x, y, z int
}

func sortTuple(tup []tuple) {
	sort.Slice(tup, func(i, j int) bool {
		if tup[i].x == tup[j].x {
			if tup[i].y == tup[j].y {
				return tup[i].z < tup[j].z
			}
			return tup[i].y < tup[j].y
		}
		return tup[i].x < tup[j].x
	})
}

func op(a, b P) P { return P{a.x + b.x, a.y + b.y} }
func e() P        { return P{0, 0} }
func mp(f int, x P) P {
	if f != 0 {
		return P{x.y - x.x, x.y}
	}
	return x
}
func cmp(f, g int) int { return f ^ g }
func id() int          { return 0 }

type P struct{ x, y int }

type E func() P
type Op func(a, b P) P
type Mapping func(f int, x P) P
type Composition func(f, g int) int
type Id func() int
type Compare func(v P) bool
type LazySegTree struct {
	n           int
	size        int
	log         int
	d           []P
	lz          []int
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
	lseg.lz = make([]int, lseg.size+1)
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
func (lseg *LazySegTree) allapply(k int, f int) {
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
func (lseg *LazySegTree) Apply(p int, f int) {
	p += lseg.size
	for i := lseg.log; i >= 1; i-- {
		lseg.push(p >> i)
	}
	lseg.d[p] = lseg.mapping(f, lseg.d[p])
	for i := 1; i <= lseg.log; i++ {
		lseg.update(p >> i)
	}
}
func (lseg *LazySegTree) RangeApply(l, r int, f int) {
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
