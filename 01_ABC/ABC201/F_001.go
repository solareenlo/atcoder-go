package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Scan(&n)

	P := make([]int, n)
	pos := make([]int, n+1)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &P[i])
		pos[P[i]] = i + 1
	}

	A := make([]int, n)
	B := make([]int, n)
	C := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &A[i], &B[i], &C[i])
	}

	Asum := make([]int, n+1)
	Bsum := make([]int, n+1)
	Csum := make([]int, n+1)
	for i := 0; i < n; i++ {
		Asum[i+1] = Asum[i] + A[i]
		Bsum[i+1] = Bsum[i] + min(B[i], A[i])
		Csum[i+1] = Csum[i] + min(C[i], A[i])
	}

	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = 1 << 60
	}

	seg := newSegtree(n+1, e, min)
	res := 1 << 60
	for i := 1; i < n+1; i++ {
		dp[i] = min(Bsum[i-1], seg.Prod(0, pos[i])+Asum[i-1])
		res = min(res, dp[i]+Csum[n]-Csum[i])
		seg.Set(pos[i], dp[i]-Asum[i])
	}
	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func e() int {
	return 1 << 60
}

type E func() int
type Op func(a, b int) int
type Compare func(v int) bool
type Segtree struct {
	n    int
	size int
	log  int
	d    []int
	e    E
	op   Op
}

func newSegtree(n int, e E, op Op) *Segtree {
	seg := new(Segtree)
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
		seg.d[seg.size+i] = i
	}
	for i := seg.size - 1; i >= 1; i-- {
		seg.Update(i)
	}
	return seg
}

func (seg *Segtree) Update(k int) {
	seg.d[k] = seg.op(seg.d[2*k], seg.d[2*k+1])
}

func (seg *Segtree) Set(p, x int) {
	if p < 0 || seg.n <= p {
		panic("")
	}
	p += seg.size
	seg.d[p] = x
	for i := 1; i <= seg.log; i++ {
		seg.Update(p >> uint(i))
	}
}

func (seg *Segtree) Get(p int) int {
	if p < 0 || seg.n <= p {
		panic("")
	}
	return seg.d[p+seg.size]
}

func (seg *Segtree) Prod(l, r int) int {
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

func (seg *Segtree) AllProd() int {
	return seg.d[1]
}

func (seg *Segtree) MaxRight(l int, cmp Compare) int {
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

func (seg *Segtree) MinLeft(r int, cmp Compare) int {
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

func (seg *Segtree) ceilPow2(n int) int {
	x := 0
	for (uint(1) << x) < uint(n) {
		x++
	}
	return x
}
