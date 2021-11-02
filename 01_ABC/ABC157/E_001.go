package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	var s string
	fmt.Fscan(in, &n, &s)
	ss := strings.Split(s, "")

	pow := make([]int, 26)
	pow[0] = 1
	for i := 1; i < 26; i++ {
		pow[i] = pow[i-1] * 2
	}

	MaxE := func() int {
		return 0
	}
	MaxOp := func(a, b int) int {
		return a | b
	}
	seg := newSegtree(ss, MaxE, MaxOp)
	for i := 0; i < n; i++ {
		seg.Set(i, pow[s[i]-'a'])
	}

	var q int
	fmt.Fscan(in, &q)
	for i := 0; i < q; i++ {
		var t int
		fmt.Fscan(in, &t)
		if t == 1 {
			var itr int
			var c string
			fmt.Fscan(in, &itr, &c)
			itr--
			seg.Set(itr, pow[c[0]-'a'])
		} else {
			var l, r int
			fmt.Fscan(in, &l, &r)
			l--
			r--
			now := seg.Prod(l, r+1)
			cnt := 0
			for now > 0 {
				if now%2 == 1 {
					cnt++
				}
				now /= 2
			}
			fmt.Fprintln(out, cnt)
		}
	}
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

func newSegtree(v []string, e E, op Op) *Segtree {
	seg := new(Segtree)
	seg.n = len(v)
	seg.log = seg.ceilPow2(seg.n)
	seg.size = 1 << uint(seg.log)
	seg.d = make([]int, 2*seg.size)
	seg.e = e
	seg.op = op
	for i, _ := range seg.d {
		seg.d[i] = seg.e()
	}
	for i := 0; i < seg.n; i++ {
		val, _ := strconv.Atoi(v[i])
		seg.d[seg.size+i] = val
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
