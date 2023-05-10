package main

import (
	"bufio"
	"fmt"
	"os"
)

type dat struct {
	A, B, C, D int
}

var N int
var B, W, sum [1 << 17]int

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &B[i], &W[i])
		if i%2 == 1 {
			B[i], W[i] = W[i], B[i]
		}
	}
	Init := make([]dat, N)
	for i := 0; i < N; i++ {
		Init[i] = dat{0, 0, 0, W[i] - B[i]}
		sum[i+1] = sum[i] + B[i]
	}
	P := NewSegTree(Init, op, e)
	var Q int
	fmt.Fscan(in, &Q)
	for Q > 0 {
		Q--
		var L, R int
		fmt.Fscan(in, &L, &R)
		L--
		ret := sum[R] - sum[L]
		ad := P.Prod(L, R).A
		if ad > 0 {
			ret += ad
		}
		fmt.Println(ret)
	}
}

func op(l, r dat) dat {
	if l.A == -1 {
		return r
	} else if r.A == -1 {
		return l
	}
	nxt := dat{0, 0, 0, 0}
	nxt.A = maxSlice(l.A+r.A, l.B+r.A, l.A+r.C)
	nxt.B = maxSlice(l.A+r.B, l.A+r.D, l.B+r.B)
	nxt.C = maxSlice(l.C+r.A, l.C+r.C, l.D+r.A)
	nxt.D = maxSlice(l.C+r.B, l.C+r.D, l.D+r.B)
	return nxt
}

func e() dat {
	x := dat{0, 0, 0, 0}
	x.A = -1
	return x
}

type E func() dat
type Op func(a, b dat) dat
type Compare func(v dat) bool
type SegTree struct {
	n    int
	size int
	log  int
	d    []dat
	e    E
	op   Op
}

func NewSegTree(n []dat, op Op, e E) *SegTree {
	seg := new(SegTree)
	seg.n = len(n)
	seg.log = seg.ceilPow2(seg.n)
	seg.size = 1 << uint(seg.log)
	seg.d = make([]dat, 2*seg.size)
	seg.e = e
	seg.op = op
	for i, _ := range seg.d {
		seg.d[i] = seg.e()
	}
	for i := 0; i < seg.n; i++ {
		seg.d[seg.size+i] = n[i]
	}
	for i := seg.size - 1; i >= 1; i-- {
		seg.Update(i)
	}
	return seg
}

func (seg *SegTree) Update(k int) {
	seg.d[k] = seg.op(seg.d[2*k], seg.d[2*k+1])
}

func (seg *SegTree) Set(p int, x dat) {
	if p < 0 || seg.n <= p {
		panic("")
	}
	p += seg.size
	seg.d[p] = x
	for i := 1; i <= seg.log; i++ {
		seg.Update(p >> uint(i))
	}
}

func (seg *SegTree) Get(p int) dat {
	if p < 0 || seg.n <= p {
		panic("")
	}
	return seg.d[p+seg.size]
}

func (seg *SegTree) Prod(l, r int) dat {
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

func (seg *SegTree) AllProd() dat {
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

func maxSlice(a ...int) int {
	res := a[0]
	for i := range a {
		if res < a[i] {
			res = a[i]
		}
	}
	return res
}
