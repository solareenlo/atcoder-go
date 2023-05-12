package main

import (
	"bufio"
	"fmt"
	"os"
)

const INF = 1 << 60

var num int
var par, sz, in, nxt, A, C [1 << 17]int
var G [1 << 17][]int
var W []int

func main() {
	IN := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(IN, &n)
	for i := 1; i < n; i++ {
		fmt.Fscan(IN, &par[i])
		par[i]--
		G[par[i]] = append(G[par[i]], i)
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(IN, &A[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(IN, &C[i])
	}
	var q int
	fmt.Fscan(IN, &q)
	W = make([]int, n)
	dfs_sz(0)
	dfs_hld(0)
	seg := NewLazySegTree(W, op, e, mapping, composition, id)
	for i := 0; i < q+1; i++ {
		t := 1
		v := 0
		x := A[0]
		var k int
		if i != 0 {
			fmt.Fscan(IN, &t, &v, &x)
			v--
		}
		if t == 1 {
			k = x - A[v]
			A[v] = x
		} else {
			k = C[v] - x
			C[v] = x
		}
		for nxt[v] != 0 {
			seg.RangeApply(in[nxt[v]], in[v]+1, k)
			v = par[nxt[v]]
		}
		seg.RangeApply(0, in[v]+1, k)
		if seg.AllProd() >= 0 {
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
		}
	}
}

func dfs_sz(u int) {
	sz[u] = 1
	for v := range G[u] {
		dfs_sz(G[u][v])
		sz[u] += sz[G[u][v]]
		if sz[G[u][v]] > sz[G[u][0]] {
			G[u][v], G[u][0] = G[u][0], G[u][v]
		}
	}
}

func dfs_hld(u int) {
	in[u] = num
	num++
	W[in[u]] = A[u] - C[u]
	for _, v := range G[u] {
		if v == G[u][0] {
			nxt[v] = nxt[u]
		} else {
			nxt[v] = v
		}
		dfs_hld(v)
		W[in[u]] += W[in[v]]
	}
}

func op(a, b int) int          { return min(a, b) }
func e() int                   { return INF }
func mapping(f, x int) int     { return f + x }
func composition(f, g int) int { return f + g }
func id() int                  { return 0 }

type E func() int
type Op func(a, b int) int
type Mapping func(f int, x int) int
type Composition func(f, g int) int
type Id func() int
type Compare func(v int) bool
type LazySegTree struct {
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

func NewLazySegTree(v []int, op Op, e E, mapping Mapping, composition Composition, id Id) *LazySegTree {
	lseg := new(LazySegTree)
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
func (lseg *LazySegTree) Set(p int, x int) {
	p += lseg.size
	for i := lseg.log; i <= 1 && i >= 0; i-- {
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
