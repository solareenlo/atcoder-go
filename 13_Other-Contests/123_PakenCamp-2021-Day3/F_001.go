package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var A, S, T, ans [3 << 17]int
	var N, M int
	fmt.Fscan(in, &N, &M)
	uv := make([]pair, 0)
	for i := 0; i < M; i++ {
		fmt.Fscan(in, &A[i])
		A[i]--
		if i > 0 {
			u := A[i-1]
			v := A[i]
			if u > v {
				u, v = v, u
			}
			uv = append(uv, pair{u, v})
			ans[0] += v - u
		}
	}

	var Q int
	fmt.Fscan(in, &Q)
	for i := 1; i < Q; i++ {
		ans[i] = ans[0]
	}
	st := make([]int, Q)
	for i := 0; i < Q; i++ {
		st[i] = i
		fmt.Fscan(in, &S[i], &T[i])
		S[i]--
		T[i]--
	}
	{ // uSvT,uSTv
		P := NewSegTreeDat(N, e, op)
		sort.Slice(st, func(i, j int) bool {
			return S[st[i]] < S[st[j]]
		})
		sortPair(uv)
		uvi := 0
		for _, id := range st {
			s := S[id]
			t := T[id]
			for uvi < len(uv) && uv[uvi].x < s {
				now := P.Get(uv[uvi].y)
				now.x += uv[uvi].y
				now.y++
				P.Set(uv[uvi].y, now)
				uvi++
			}
			// s+t+1<2v
			vl := max(s, (s+t+1+1)/2)
			if vl <= t {
				now := P.Prod(vl, t+1)
				ans[id] += (s+t+1)*now.y - 2*now.x
			}
			{
				now := P.Prod(t+1, N)
				ans[id] += (-(t - s) + 1) * now.y
			}
		}
	}
	{ // SuTv
		P := NewSegTreeDat(N, e, op)
		sort.Slice(st, func(i, j int) bool {
			return T[st[i]] > T[st[j]]
		})
		sort.Slice(uv, func(i, j int) bool {
			return uv[i].y > uv[j].y
		})
		uvi := 0
		for _, id := range st {
			s := S[id]
			t := T[id]
			for uvi < len(uv) && uv[uvi].y > t {
				now := P.Get(uv[uvi].x)
				now.x += uv[uvi].x
				now.y++
				P.Set(uv[uvi].x, now)
				uvi++
			}
			// 2u<s+t-1
			ur := min(t, (s+t-1)/2)
			if s <= ur {
				now := P.Prod(s, ur+1)
				ans[id] += 2*now.x - (s+t-1)*now.y
			}
		}
	}
	{ // SuvT
		{ // uvST,uSvT,uSTv
			P := NewSegTreeDat(N, e, op)
			sort.Slice(st, func(i, j int) bool {
				return S[st[i]] < S[st[j]]
			})
			sortPair(uv)
			uvi := 0
			for _, id := range st {
				s := S[id]
				t := T[id]
				for uvi < len(uv) && uv[uvi].x < s {
					vu := uv[uvi].y - uv[uvi].x
					now := P.Get(vu)
					now.x += vu
					now.y++
					P.Set(vu, now)
					uvi++
				}
				// t-s+1<2(v-u)
				l := (t - s + 1 + 1) / 2
				now := P.Prod(l, N)
				ans[id] -= (t-s+1)*now.y - 2*now.x
			}
		}
		{ // uvST,uSvT,SuvT
			P := NewSegTreeDat(N, e, op)
			sort.Slice(st, func(i, j int) bool {
				return T[st[i]] < T[st[j]]
			})
			sort.Slice(uv, func(i, j int) bool {
				return uv[i].y < uv[j].y
			})
			uvi := 0
			for _, id := range st {
				s := S[id]
				t := T[id]
				for uvi < len(uv) && uv[uvi].y <= t {
					vu := uv[uvi].y - uv[uvi].x
					now := P.Get(vu)
					now.x += vu
					now.y++
					P.Set(vu, now)
					uvi++
				}
				// t-s+1<2(v-u)
				l := (t - s + 1 + 1) / 2
				now := P.Prod(l, N)
				ans[id] += (t-s+1)*now.y - 2*now.x
			}
		}
		{ // uSTv
			P := NewSegTreeDat(N, e, op)
			sort.Slice(st, func(i, j int) bool {
				return S[st[i]] < S[st[j]]
			})
			sortPair(uv)
			uvi := 0
			for _, id := range st {
				s := S[id]
				t := T[id]
				for uvi < len(uv) && uv[uvi].x < s {
					now := P.Get(uv[uvi].y)
					now.x += uv[uvi].y - uv[uvi].x
					now.y++
					P.Set(uv[uvi].y, now)
					uvi++
				}
				now := P.Prod(t+1, N)
				ans[id] += (t-s+1)*now.y - 2*now.x
			}
		}
	}
	for i := 0; i < Q; i++ {
		fmt.Fprintln(out, ans[i])
	}
}

type dat struct {
	x, y int
}

func op(a, b dat) dat { return dat{a.x + b.x, a.y + b.y} }
func e() dat          { return dat{0, 0} }

type E func() dat
type Op func(a, b dat) dat
type Compare func(v dat) bool
type SetTreeDat struct {
	n    int
	size int
	log  int
	d    []dat
	e    E
	op   Op
}

func NewSegTreeDat(n int, e E, op Op) *SetTreeDat {
	seg := new(SetTreeDat)
	seg.n = n
	seg.log = seg.ceilPow2(seg.n)
	seg.size = 1 << uint(seg.log)
	seg.d = make([]dat, 2*seg.size)
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

func (seg *SetTreeDat) Update(k int) {
	seg.d[k] = seg.op(seg.d[2*k], seg.d[2*k+1])
}

func (seg *SetTreeDat) Set(p int, x dat) {
	if p < 0 || seg.n <= p {
		panic("")
	}
	p += seg.size
	seg.d[p] = x
	for i := 1; i <= seg.log; i++ {
		seg.Update(p >> uint(i))
	}
}

func (seg *SetTreeDat) Get(p int) dat {
	if p < 0 || seg.n <= p {
		panic("")
	}
	return seg.d[p+seg.size]
}

func (seg *SetTreeDat) Prod(l, r int) dat {
	if l < 0 || r < l || seg.n < r {
		fmt.Println("l,r,n:", l, r, seg.n)
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

func (seg *SetTreeDat) AllProd() dat {
	return seg.d[1]
}

func (seg *SetTreeDat) MaxRight(l int, cmp Compare) int {
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

func (seg *SetTreeDat) MinLeft(r int, cmp Compare) int {
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

func (seg *SetTreeDat) ceilPow2(n int) int {
	x := 0
	for (uint(1) << x) < uint(n) {
		x++
	}
	return x
}

type pair struct {
	x, y int
}

func sortPair(tmp []pair) {
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].x == tmp[j].x {
			return tmp[i].y < tmp[j].y
		}
		return tmp[i].x < tmp[j].x
	})
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
