package main

import (
	"bufio"
	"fmt"
	"os"
)

const INF = 5e8
const MAX_N = 300000

var ar, pos []int
var cnt int = 0
var seg *segTree
var n int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &n)
	ar = make([]int, 100005)
	pos = make([]int, 100005)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &ar[i])
		ar[i]--
		pos[ar[i]] = i
	}
	seg = NewSegTree()
	seg.init(n, ar)
	for i := 0; i < n-1; i++ {
		p := pos[i]
		q := pos[i+1]
		cnt += check(p, q)
	}
	var q int
	fmt.Fscan(in, &q)
	for i := 0; i < q; i++ {
		var x int
		fmt.Fscan(in, &x)
		x--
		cnt -= doit(ar[x], ar[x+1])
		ar[x], ar[x+1] = ar[x+1], ar[x]
		pos[ar[x]] = x
		pos[ar[x+1]] = x + 1
		seg.set(x, ar[x])
		seg.set(x+1, ar[x+1])
		cnt += doit(ar[x], ar[x+1])
		if cnt == n-1 {
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
		}
	}
}

func check(l, r int) int {
	if l >= r {
		return 1
	}
	if seg.query(l, r, 0, 0, seg.n) <= ar[r] {
		return 1
	}
	return 0
}

func doit(v, v2 int) int {
	res := 0
	if v > 0 {
		res += check(pos[v-1], pos[v])
	}
	if v < n-1 {
		res += check(pos[v], pos[v+1])
	}
	if v2 > 0 && v+1 != v2 {
		res += check(pos[v2-1], pos[v2])
	}
	if v2 < n-1 && v-1 != v2 {
		res += check(pos[v2], pos[v2+1])
	}
	return res
}

type segTree struct {
	val []int
	n   int
}

func NewSegTree() *segTree {
	seg := new(segTree)
	seg.n = 0
	seg.val = make([]int, MAX_N)
	return seg
}

func (s *segTree) init(n_ int, a []int) {
	s.n = 1
	s.val = make([]int, MAX_N)
	for s.n < n_ {
		s.n <<= 1
	}
	for i := 0; i < s.n*2; i++ {
		s.val[i] = -INF
	}
	for i := 0; i < n_; i++ {
		s.val[i+s.n-1] = a[i]
	}
	for i := s.n - 2; i >= 0; i-- {
		s.val[i] = max(s.val[i*2+1], s.val[i*2+2])
	}
}

func (s *segTree) set(k, a int) {
	k += s.n - 1
	s.val[k] = a
	for k > 0 {
		k = (k - 1) / 2
		s.val[k] = max(s.val[k*2+1], s.val[k*2+2])
	}
}

func (s segTree) query(a, b, i, l, r int) int {
	if a <= l && r <= b {
		return s.val[i]
	}
	if b <= l || r <= a {
		return -INF
	}
	md := (l + r) >> 1
	return max(s.query(a, b, i*2+1, l, md), s.query(a, b, i*2+2, md, r))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
