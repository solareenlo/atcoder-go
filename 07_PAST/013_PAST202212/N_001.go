package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

var value int = 0

type Query struct {
	l, r, b, idx int
}

type Mo struct {
	res         []int
	data        []Query
	width, n, q int
}

func (m *Mo) init(n, q int) {
	m.n = n
	m.q = q
	m.width = int(math.Max(1.0, float64(n)/math.Max(1.0, math.Sqrt(float64(q)*2.0/3.0))))
	m.res = make([]int, q)
	m.data = make([]Query, q)
}

var idxInsert int = 0

func (m *Mo) insert(l, r int) {
	m.data[idxInsert] = Query{l, r, l / m.width, idxInsert}
	idxInsert++
}

func (m *Mo) build() {
	sort.Slice(m.data, func(i, j int) bool {
		if m.data[i].b != m.data[j].b {
			return m.data[i].b < m.data[j].b
		}
		if (m.data[i].b & 1) != 0 {
			return m.data[i].r > m.data[j].r
		}
		return m.data[i].r < m.data[j].r
	})
	nl, nr := 0, 0
	for _, v := range m.data {
		for nl > v.l {
			nl--
			m.add_left(nl)
		}
		for nr < v.r {
			m.add_right(nr)
			nr++
		}
		for nl < v.l {
			m.del_left(nl)
			nl++
		}
		for nr > v.r {
			nr--
			m.del_right(nr)
		}
		m.res[v.idx] = value
	}
}

var d, idx []int
var seg *SegTree

func (m *Mo) add_left(p int) {
	mx := seg.MaxRight(idx[p], cmp)
	mi := seg.MinLeft(idx[p], cmp)
	mi--
	if mx < m.n {
		value -= d[mx] * d[idx[p]]
	}
	if mi >= 0 {
		value -= d[mi] * d[idx[p]]
	}
	if mx < m.n && mi >= 0 {
		value += d[mx] * d[mi]
	}
	seg.Update(idx[p], 1)
}

func (m *Mo) add_right(p int) {
	m.add_left(p)
}

func (m *Mo) del_left(p int) {
	seg.Update(idx[p], -1)
	mx := seg.MaxRight(idx[p], cmp)
	mi := seg.MinLeft(idx[p], cmp)
	mi--
	if mx < m.n {
		value += d[mx] * d[idx[p]]
	}
	if mi >= 0 {
		value += d[mi] * d[idx[p]]
	}
	if mx < m.n && mi >= 0 {
		value -= d[mx] * d[mi]
	}
}

func (m *Mo) del_right(p int) {
	m.del_left(p)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	seg = NewSegTree(20010, opSeg, eSeg)

	var n int
	fmt.Fscan(in, &n)
	d = make([]int, n)
	I := make([]int, n)
	mi := NewSegTree(n, opMin, eMin)
	mx := NewSegTree(n, opMax, eMax)
	c := make([]int, n+1)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &d[i])
		I[i] = i
		mi.Set(i, d[i]*d[i])
		mx.Set(i, d[i]*d[i])
		c[i+1] = c[i] + d[i]*d[i]
	}
	mi.build()
	mx.build()
	for i := 0; i < n+1; i++ {
		seg.Set(i, -1)
	}
	seg.build()
	sort.Slice(I, func(i, j int) bool {
		return d[I[i]] < d[I[j]]
	})
	sort.Ints(d)
	idx = make([]int, n)
	for i := 0; i < n; i++ {
		idx[I[i]] = i
	}
	var q int
	fmt.Fscan(in, &q)
	var mo Mo
	mo.init(n, q)
	val := make([]int, q)
	for i := 0; i < q; i++ {
		var l, r int
		fmt.Fscan(in, &l, &r)
		l--
		mo.insert(l, r)
		val[i] += (c[r] - c[l]) * 2
		val[i] -= mi.Prod(l, r) + mx.Prod(l, r)
	}
	mo.build()
	for i := 0; i < q; i++ {
		fmt.Fprintln(out, mo.res[i]*2+val[i])
	}
}

func eMin() int          { return 9223372036854775807 }
func opMin(a, b int) int { return min(a, b) }

func eMax() int          { return -9223372036854775807 }
func opMax(a, b int) int { return max(a, b) }

func eSeg() int          { return -2147483647 }
func opSeg(a, b int) int { return max(a, b) }

func cmp(x int) bool { return x < 0 }

type E func() int
type Op func(a, b int) int
type Compare func(v int) bool
type SegTree struct {
	n    int
	size int
	log  int
	d    []int
	e    E
	op   Op
}

func NewSegTree(n int, op Op, e E) *SegTree {
	seg := new(SegTree)
	seg.n = n
	seg.log = seg.ceilPow2(seg.n)
	seg.size = 1 << uint(seg.log)
	seg.d = make([]int, 2*seg.size)
	seg.e = e
	seg.op = op
	for i, _ := range seg.d {
		seg.d[i] = seg.e()
	}
	for i := 0; i < seg.n; i++ {
		seg.d[seg.size+i] = seg.e()
	}
	return seg
}

func (seg *SegTree) build() {
	for k := seg.size - 1; k >= 1; k-- {
		seg.d[k] = seg.op(seg.d[2*k], seg.d[2*k+1])
	}
}

func (seg *SegTree) Set(k, x int) {
	seg.d[k+seg.size] = x
}

func (seg *SegTree) Update(p, x int) {
	p += seg.size
	seg.d[p] = x
	for p > 1 {
		p >>= 1
		seg.d[p] = seg.op(seg.d[p*2], seg.d[p*2+1])
	}
}

func (seg *SegTree) Get(p int) int {
	if p < 0 || seg.n <= p {
		panic("")
	}
	return seg.d[p+seg.size]
}

func (seg *SegTree) Prod(l, r int) int {
	if l < 0 || r < l || seg.n < r {
		panic("")
	}
	sml, smr := seg.e(), seg.e()
	l += seg.size
	r += seg.size
	for l < r {
		if (l & 1) == 1 {
			sml = seg.op(sml, seg.d[l])
			l++
		}
		if (r & 1) == 1 {
			r--
			smr = seg.op(seg.d[r], smr)
		}
		l >>= 1
		r >>= 1
	}
	return seg.op(sml, smr)
}

func (seg *SegTree) AllProd() int {
	return seg.d[1]
}

func (seg *SegTree) MaxRight(l int, cmp Compare) int {
	if l < 0 || seg.n < l {
		panic("")
	}
	if !cmp(seg.e()) {
		panic("")
	}
	if l == seg.n {
		return seg.n
	}
	l += seg.size
	sm := seg.e()
	for {
		for l%2 == 0 {
			l >>= 1
		}
		if !cmp(seg.op(sm, seg.d[l])) {
			for l < seg.size {
				l = 2 * l
				if cmp(seg.op(sm, seg.d[l])) {
					sm = seg.op(sm, seg.d[l])
					l++
				}
			}
			return l - seg.size
		}
		sm = seg.op(sm, seg.d[l])
		l++
		if l&-l == l {
			break
		}
	}
	return seg.n
}

func (seg *SegTree) MinLeft(r int, cmp Compare) int {
	if r < 0 || seg.n < r {
		panic("")
	}
	if !cmp(seg.e()) {
		panic("")
	}
	if r == 0 {
		return 0
	}
	r += seg.size
	sm := seg.e()
	for {
		r--
		for r > 1 && r%2 != 0 {
			r >>= 1
		}
		if !cmp(seg.op(seg.d[r], sm)) {
			for r < seg.size {
				r = 2*r + 1
				if cmp(seg.op(seg.d[r], sm)) {
					sm = seg.op(seg.d[r], sm)
					r--
				}
			}
			return r + 1 - seg.size
		}
		sm = seg.op(seg.d[r], sm)
		if r&-r == r {
			break
		}
	}
	return 0
}

func (seg *SegTree) ceilPow2(n int) int {
	x := 0
	for (uint(1) << x) < uint(n) {
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
