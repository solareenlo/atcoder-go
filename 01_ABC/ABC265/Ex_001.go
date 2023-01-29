package main

import (
	"bufio"
	"fmt"
	"os"
)

type P [6]int
type F [3]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, q int
	fmt.Fscan(in, &n, &q)
	v := make([]P, n)
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		v[i][a] = 1
	}
	seg := newLazySegtree(v, op, e, mapping, composition, id)

	for q > 0 {
		var ty, l, r int
		fmt.Fscan(in, &ty, &l, &r)
		l--
		if ty == 1 {
			ret := seg.Prod(l, r)
			fmt.Fprintln(out, ret[3]+ret[4]+ret[5])
		} else {
			var s, t, u int
			fmt.Fscan(in, &s, &t, &u)
			f := F{s, t, u}
			seg.RangeApply(l, r, f)
		}
		q--
	}
}

func op(x, y P) P {
	var ret P
	ret[0] = x[0] + y[0]
	ret[1] = x[1] + y[1]
	ret[2] = x[2] + y[2]
	ret[3] = x[3] + y[3] + x[1]*y[0]
	ret[4] = x[4] + y[4] + x[2]*y[0]
	ret[5] = x[5] + y[5] + x[2]*y[1]
	return ret
}

func e() P { return P{0, 0, 0, 0, 0, 0} }

func mapping(f F, x P) P {
	var ret P
	ret[f[0]] += x[0]
	ret[f[1]] += x[1]
	ret[f[2]] += x[2]
	var cnt [3][3]int
	cnt[f[1]][f[0]] += x[3]
	cnt[f[2]][f[0]] += x[4]
	cnt[f[2]][f[1]] += x[5]
	cnt[f[0]][f[1]] += x[0]*x[1] - x[3]
	cnt[f[0]][f[2]] += x[0]*x[2] - x[4]
	cnt[f[1]][f[2]] += x[1]*x[2] - x[5]
	ret[3] = cnt[1][0]
	ret[4] = cnt[2][0]
	ret[5] = cnt[2][1]
	return ret
}

func composition(f, g F) F {
	var ret F
	ret[0] = f[g[0]]
	ret[1] = f[g[1]]
	ret[2] = f[g[2]]
	return ret
}

func id() F { return F{0, 1, 2} }

type E func() P
type Op func(a, b P) P
type Mapping func(f F, x P) P
type Composition func(f, g F) F
type Id func() F
type Compare func(v P) bool
type LazySegtree struct {
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

func newLazySegtree(v []P, op Op, e E, mapping Mapping, composition Composition, id Id) *LazySegtree {
	lseg := new(LazySegtree)
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
func (lseg *LazySegtree) update(k int) {
	lseg.d[k] = lseg.op(lseg.d[2*k], lseg.d[2*k+1])
}
func (lseg *LazySegtree) allapply(k int, f F) {
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
func (lseg *LazySegtree) Set(p int, x P) {
	p += lseg.size
	for i := lseg.log; i <= 1 && i >= 0; i-- {
		lseg.push(p >> i)
	}
	lseg.d[p] = x
	for i := 1; i <= lseg.log; i++ {
		lseg.update(p >> i)
	}
}
func (lseg *LazySegtree) Get(p int) P {
	p += lseg.size
	for i := lseg.log; i >= 1; i-- {
		lseg.push(p >> i)
	}
	return lseg.d[p]
}
func (lseg *LazySegtree) Prod(l, r int) P {
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
func (lseg *LazySegtree) AllProd() P {
	return lseg.d[1]
}
func (lseg *LazySegtree) Apply(p int, f F) {
	p += lseg.size
	for i := lseg.log; i >= 1; i-- {
		lseg.push(p >> i)
	}
	lseg.d[p] = lseg.mapping(f, lseg.d[p])
	for i := 1; i <= lseg.log; i++ {
		lseg.update(p >> i)
	}
}
func (lseg *LazySegtree) RangeApply(l, r int, f F) {
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
