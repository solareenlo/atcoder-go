package main

import (
	"bufio"
	"fmt"
	"os"
)

func op(a, b int) int {
	if a == -2 {
		return b
	}
	if b == -2 {
		return a
	}
	if a == -1 || b == -1 {
		return -1
	}
	if a == b {
		return a
	}
	return -1
}

func e() int { return -2 }

func mapping(a, b int) int {
	if a == -1 {
		return b
	}
	return a
}

func composition(a, b int) int {
	if a == -1 {
		return b
	}
	if b == -1 {
		return a
	}
	return a
}

func id() int { return -1 }

func f(a int) bool {
	if a != -1 {
		return true
	}
	return false
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N int
	fmt.Fscan(in, &N)
	mp := make(map[int]int)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
		mp[A[i]]++
	}

	seg := newLazySegtree(A, op, e, mapping, composition, id)
	cur := 0
	for _, a := range mp {
		cur += a * (a - 1) / 2
	}

	var Q int
	fmt.Fscan(in, &Q)
	for i := 0; i < Q; i++ {
		var l, r, x int
		fmt.Fscan(in, &l, &r, &x)
		l--
		for l != r {
			ok := min(r, seg.MaxRight(l, f))
			v := seg.Prod(l, ok)
			cur -= mp[v] * (mp[v] - 1) / 2
			mp[v] -= ok - l
			cur += mp[v] * (mp[v] - 1) / 2
			cur -= mp[x] * (mp[x] - 1) / 2
			mp[x] += ok - l
			cur += mp[x] * (mp[x] - 1) / 2
			seg.RangeApply(l, ok, x)
			l = ok
		}
		fmt.Fprintln(out, cur)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type E func() int
type Op func(a, b int) int
type Mapping func(f int, x int) int
type Composition func(f, g int) int
type Id func() int
type Compare func(v int) bool
type LazySegtree struct {
	n           int
	size        int
	log         int
	d           []int
	lz          []int
	e           E
	op          Op
	mapping     Mapping
	composition Composition
	id          Id
}

func newLazySegtree(v []int, op Op, e E, mapping Mapping, composition Composition, id Id) *LazySegtree {
	lseg := new(LazySegtree)
	lseg.n = len(v)
	lseg.log = lseg.ceilPow2(lseg.n)
	lseg.size = 1 << uint(lseg.log)
	lseg.d = make([]int, 2*(lseg.size+1))
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
func (lseg *LazySegtree) update(k int) {
	lseg.d[k] = lseg.op(lseg.d[2*k], lseg.d[2*k+1])
}
func (lseg *LazySegtree) allapply(k int, f int) {
	lseg.d[k] = lseg.mapping(f, lseg.d[k])
	if k < lseg.size {
		lseg.lz[k] = lseg.composition(f, lseg.lz[k])
	}
}
func (lseg *LazySegtree) push(k int) {
	lseg.allapply(2*k, lseg.lz[k])
	lseg.allapply(2*k+1, lseg.lz[k])
	lseg.lz[k] = lseg.id()
}
func (lseg *LazySegtree) Set(p int, x int) {
	p += lseg.size
	for i := lseg.log; i <= 1 && i >= 0; i-- {
		lseg.push(p >> i)
	}
	lseg.d[p] = x
	for i := 1; i <= lseg.log; i++ {
		lseg.update(p >> i)
	}
}
func (lseg *LazySegtree) Get(p int) int {
	p += lseg.size
	for i := lseg.log; i >= 1; i-- {
		lseg.push(p >> i)
	}
	return lseg.d[p]
}
func (lseg *LazySegtree) Prod(l, r int) int {
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
func (lseg *LazySegtree) AllProd() int {
	return lseg.d[1]
}
func (lseg *LazySegtree) Apply(p int, f int) {
	p += lseg.size
	for i := lseg.log; i >= 1; i-- {
		lseg.push(p >> i)
	}
	lseg.d[p] = lseg.mapping(f, lseg.d[p])
	for i := 1; i <= lseg.log; i++ {
		lseg.update(p >> i)
	}
}
func (lseg *LazySegtree) RangeApply(l, r int, f int) {
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
func (lseg *LazySegtree) MaxRight(l int, cmp Compare) int {
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
func (lseg *LazySegtree) MinLeft(r int, cmp Compare) int {
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
func (lseg *LazySegtree) ceilPow2(n int) int {
	x := 0
	for (1 << uint(x)) < n {
		x++
	}
	return x
}
