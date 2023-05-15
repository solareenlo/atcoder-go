package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, q int
	fmt.Fscan(in, &n, &q)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	seg := NewSegTree(a, e, op)

	mp := make(map[int]int)
	for i := 0; i < n; i++ {
		r := i
		for r != n {
			Y = seg.Prod(i, r+1)
			ok := seg.MaxRight(i, com)
			mp[seg.Prod(i, r+1)] += ok - r
			r = ok
		}
	}
	for q > 0 {
		q--
		var t int
		fmt.Fscan(in, &t)
		if t == 3 {
			var s int
			fmt.Fscan(in, &s)
			ans := 0
			if _, ok := mp[s]; ok {
				ans = mp[s]
			}
			fmt.Println(ans)
		} else {
			var k, a int
			fmt.Fscan(in, &k, &a)
			k--
			if t == 2 {
				a *= -1
			}
			r := k
			type pair struct {
				x, y int
			}
			L := make([]pair, 0)
			R := make([]pair, 0)
			for r != n {
				Y = seg.Prod(k, r+1)
				ok := seg.MaxRight(k, com)
				R = append(R, pair{Y, ok - r})
				r = ok
			}
			l := k
			for l != -1 {
				Y = seg.Prod(l, k+1)
				ok := seg.MinLeft(k+1, com)
				L = append(L, pair{Y, abs(l-ok) + 1})
				l = ok - 1
			}
			for x := 0; x < len(L); x++ {
				for y := 0; y < len(R); y++ {
					mp[op(L[x].x, R[y].x)] -= L[x].y * R[y].y
				}
			}
			seg.Set(k, seg.Get(k)+a)
			r1 := k
			L1 := make([]pair, 0)
			R1 := make([]pair, 0)
			for r1 != n {
				Y = seg.Prod(k, r1+1)
				ok := seg.MaxRight(k, com)
				R1 = append(R1, pair{Y, ok - r1})
				r1 = ok
			}
			l1 := k
			for l1 != -1 {
				Y = seg.Prod(l1, k+1)
				ok := seg.MinLeft(k+1, com)
				L1 = append(L1, pair{Y, abs(l1-ok) + 1})
				l1 = ok - 1
			}
			for x := 0; x < len(L1); x++ {
				for y := 0; y < len(R1); y++ {
					mp[op(L1[x].x, R1[y].x)] += L1[x].y * R1[y].y
				}
			}
		}
	}
}

const Inf = 1000000001

func op(a, b int) int {
	if a == Inf || b == Inf {
		return Inf
	}
	g := gcd(a, b)
	t := a / g
	t *= b / g
	t *= g
	if t >= Inf {
		return Inf
	}
	return t
}

func e() int {
	return 1
}

var Y int

func com(X int) bool {
	return X <= Y
}

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

func NewSegTree(n []int, e E, op Op) *SegTree {
	seg := new(SegTree)
	seg.n = len(n)
	seg.log = seg.ceilPow2(seg.n)
	seg.size = 1 << uint(seg.log)
	seg.d = make([]int, 2*seg.size)
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

func (seg *SegTree) Set(p, x int) {
	if p < 0 || seg.n <= p {
		panic("")
	}
	p += seg.size
	seg.d[p] = x
	for i := 1; i <= seg.log; i++ {
		seg.Update(p >> uint(i))
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

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
