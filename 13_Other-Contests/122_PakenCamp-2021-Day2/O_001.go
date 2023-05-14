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

	var s string
	fmt.Fscan(in, &s)

	n := len(s)

	sa := SuffixArrayString(s)
	lcp := LcpArrayString(s, sa)

	need := make([]int, n)
	for i := 0; i < n-1; i++ {
		need[sa[i]] = max(need[sa[i]], lcp[i]+1)
		need[sa[i+1]] = max(need[sa[i+1]], lcp[i]+1)
	}
	for i := 0; i < n; i++ {
		if need[i] > n-i {
			need[i] = n + 1
		}
	}

	dead_time := make([][]int, n+1)
	for i := 0; i < n; i++ {
		if need[i] == n+1 {
			continue
		}
		dead_time[i+need[i]] = append(dead_time[i+need[i]], i)
	}

	var q int
	fmt.Fscan(in, &q)
	l := make([]int, q)
	r := make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(in, &l[i], &r[i])
		l[i]--
	}

	q_idx := make([]int, q)
	for i := range q_idx {
		q_idx[i] = i
	}
	sort.Slice(q_idx, func(i, j int) bool {
		return r[q_idx[i]] < r[q_idx[j]]
	})

	ans := make([]int, q)
	rmq := NewSegTree(need, e, op)
	tmp := make([]int, n)
	for i := range tmp {
		tmp[i] = e()
	}
	dead := NewSegTree(tmp, e, op)
	cur_r := 0
	for _, i := range q_idx {
		for cur_r < r[i] {
			for _, j := range dead_time[cur_r] {
				rmq.Set(j, e())
				dead.Set(j, -j)
			}
			cur_r++
		}
		ans[i] = min(rmq.Prod(0, l[i]+1), r[i]+dead.Prod(0, l[i]+1))
	}

	for i := 0; i < q; i++ {
		fmt.Fprintln(out, ans[i])
	}
}

func SuffixArrayString(s string) []int {
	n := len(s)
	s2 := make([]int, n)
	for i := 0; i < n; i++ {
		s2[i] = int(s[i])
	}
	return SAIS(s2, 255, 10, 40)
}

func SANaive(ss []int) []int {
	s := make([]int, len(ss))
	copy(s, ss)
	n := len(s)
	sa := make([]int, n)
	for i := 0; i < n; i++ {
		sa[i] = i
	}
	sort.Slice(sa, func(i, j int) bool {
		l, r := sa[i], sa[j]
		if l == r {
			return false
		}
		for l < n && r < n {
			if s[l] != s[r] {
				return s[l] < s[r]
			}
			l++
			r++
		}
		return l == n
	})
	return sa
}

func SADoubling(ss []int) []int {
	s := make([]int, len(ss))
	copy(s, ss)
	n := len(s)
	sa, rnk, tmp := make([]int, n), s, make([]int, n)
	for i := 0; i < n; i++ {
		sa[i] = i
	}
	for k := 1; k < n; k *= 2 {
		cmp := func(i, j int) bool {
			x, y := sa[i], sa[j]
			if rnk[x] != rnk[y] {
				return rnk[x] < rnk[y]
			}
			rx, ry := -1, -1
			if x+k < n {
				rx = rnk[x+k]
			}
			if y+k < n {
				ry = rnk[y+k]
			}
			return rx < ry
		}
		sort.Slice(sa, cmp)
		tmp[sa[0]] = 0
		for i := 1; i < n; i++ {
			tmp[sa[i]] = tmp[sa[i-1]]
			if cmp(i-1, i) {
				tmp[sa[i]]++
			}
		}
		tmp, rnk = rnk, tmp
	}
	return sa
}

func SAIS(ss []int, upper, thresholdNaive, thresholdDoubling int) []int {
	s := make([]int, len(ss))
	copy(s, ss)
	n := len(s)
	if n == 0 {
		return []int{}
	}
	if n == 1 {
		return []int{0}
	}
	if n == 2 {
		if s[0] < s[1] {
			return []int{0, 1}
		} else {
			return []int{1, 0}
		}
	}
	if n < thresholdNaive {
		return SANaive(s)
	}
	if n < thresholdDoubling {
		return SADoubling(s)
	}

	sa := make([]int, n)
	ls := make([]bool, n)
	for i := n - 2; i >= 0; i-- {
		if s[i] == s[i+1] {
			ls[i] = ls[i+1]
		} else {
			ls[i] = s[i] < s[i+1]
		}
	}
	suml, sums := make([]int, upper+1), make([]int, upper+1)
	for i := 0; i < n; i++ {
		if !ls[i] {
			sums[s[i]]++
		} else {
			suml[s[i]+1]++
		}
	}
	for i := 0; i <= upper; i++ {
		sums[i] += suml[i]
		if i < upper {
			suml[i+1] += sums[i]
		}
	}

	induce := func(lms []int) {
		for i := 0; i < len(sa); i++ {
			sa[i] = -1
		}
		buf := make([]int, upper+1)
		copy(buf, sums)
		for _, d := range lms {
			if d == n {
				continue
			}
			sa[buf[s[d]]] = d
			buf[s[d]]++
		}
		copy(buf, suml)
		sa[buf[s[n-1]]] = n - 1
		buf[s[n-1]]++
		for i := 0; i < n; i++ {
			v := sa[i]
			if v >= 1 && !ls[v-1] {
				sa[buf[s[v-1]]] = v - 1
				buf[s[v-1]]++
			}
		}
		copy(buf, suml)
		for i := n - 1; i >= 0; i-- {
			v := sa[i]
			if v >= 1 && ls[v-1] {
				sa[buf[s[v-1]+1]-1] = v - 1
				buf[s[v-1]+1]--
			}
		}
	}
	lmsMap := make([]int, n+1)
	for i := 0; i < len(lmsMap); i++ {
		lmsMap[i] = -1
	}
	m := 0
	for i := 1; i < n; i++ {
		if !ls[i-1] && ls[i] {
			lmsMap[i] = m
			m++
		}
	}
	lms := make([]int, 0, m)
	for i := 1; i < n; i++ {
		if !ls[i-1] && ls[i] {
			lms = append(lms, i)
		}
	}
	induce(lms)
	if m > 0 {
		sortedLms := make([]int, 0, m)
		for _, v := range sa {
			if lmsMap[v] != -1 {
				sortedLms = append(sortedLms, v)
			}
		}
		recs := make([]int, m)
		recUpper := 0
		recs[lmsMap[sortedLms[0]]] = 0
		for i := 1; i < m; i++ {
			l, r := sortedLms[i-1], sortedLms[i]
			endl, endr := n, n
			if lmsMap[l]+1 < m {
				endl = lms[lmsMap[l]+1]
			}
			if lmsMap[r]+1 < m {
				endr = lms[lmsMap[r]+1]
			}
			same := true
			if endl-l != endr-r {
				same = false
			} else {
				for l < endl {
					if s[l] != s[r] {
						break
					}
					l++
					r++
				}
				if l == n || s[l] != s[r] {
					same = false
				}
			}
			if !same {
				recUpper++
			}
			recs[lmsMap[sortedLms[i]]] = recUpper
		}
		recSa := SAIS(recs, recUpper, thresholdNaive, thresholdDoubling)
		for i := 0; i < m; i++ {
			sortedLms[i] = lms[recSa[i]]
		}
		induce(sortedLms)
	}
	return sa
}

func LcpArrayInt(s, sa []int) []int {
	n := len(s)
	if n < 1 {
		panic("length of slice s must be more than or equal to 1")
	}
	rnk := make([]int, n)
	for i := 0; i < n; i++ {
		rnk[sa[i]] = i
	}
	lcp := make([]int, n-1)
	h := 0
	for i := 0; i < n; i++ {
		if h > 0 {
			h--
		}
		if rnk[i] == 0 {
			continue
		}
		j := sa[rnk[i]-1]
		for ; j+h < n && i+h < n; h++ {
			if s[j+h] != s[i+h] {
				break
			}
		}
		lcp[rnk[i]-1] = h
	}
	return lcp
}

func LcpArrayString(s string, sa []int) []int {
	n := len(s)
	s2 := make([]int, n)
	for i := 0; i < n; i++ {
		s2[i] = int(s[i])
	}
	return LcpArrayInt(s2, sa)
}

func e() int {
	return 1 << 30
}

func op(x, y int) int {
	return min(x, y)
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
