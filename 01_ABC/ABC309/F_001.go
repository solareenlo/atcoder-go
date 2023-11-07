package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var h [3][1 << 18]int

	var N int
	fmt.Fscan(in, &N)
	for i := 0; i < N; i++ {
		y := make([]int, 3)
		for j := 0; j < 3; j++ {
			fmt.Fscan(in, &y[j])
		}
		sort.Ints(y)
		for j := 0; j < 3; j++ {
			h[j][i] = y[j]
		}
	}
	v := make([]int, N)
	for i := 0; i < N; i++ {
		v[i] = h[1][i]
	}
	sort.Ints(v)
	v = unique(v)
	for i := 0; i < N; i++ {
		h[1][i] = lowerBound(v, h[1][i])
	}
	x := make([]tuple, N)
	for i := 0; i < N; i++ {
		x[i] = tuple{h[0][i], -h[1][i], h[2][i]}
	}
	sortTuple(x)
	s := NewSegTree(N, op, e)
	a := 0
	for i := 0; i < N; i++ {
		_, W, D := x[i].x, x[i].y, x[i].z
		if s.Prod(0, -W) < D {
			a++
		}
		if s.Get(-W) > D {
			s.Set(-W, D)
		}
	}
	if a != 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func op(a, b int) int { return min(a, b) }
func e() int          { return (1 << 61) }

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

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
