package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const INFLL = 4557430888798830399
const INF = 1061109567

type Segtree struct {
	N   int
	dat []int
}

func (seg *Segtree) Init(n int) {
	seg.N = 1
	for seg.N < n {
		seg.N <<= 1
	}
	seg.dat = make([]int, 2*seg.N)
	for i := range seg.dat {
		seg.dat[i] = INFLL
	}
}

func (seg *Segtree) Set(k, x int) {
	k += seg.N
	seg.dat[k] = min(seg.dat[k], x)
	for k > 1 {
		k >>= 1
		seg.dat[k] = min(seg.dat[k*2], seg.dat[k*2+1])
	}
}

func (seg *Segtree) Prod(l, r int) int {
	res := INFLL
	l += seg.N
	r += seg.N
	for l < r {
		if (l & 1) != 0 {
			res = min(res, seg.dat[l])
			l++
		}
		if (r & 1) != 0 {
			r--
			res = min(res, seg.dat[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}

type st struct {
	x, y, c, id int
}

var ans [200000]int

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, s, t int
	fmt.Fscan(in, &n, &s, &t)
	s = -s
	t = -t
	xs := make([]int, 0)
	ys := make([]int, 0)
	p := make([]st, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &p[i].x, &p[i].y, &p[i].c)
		xs = append(xs, p[i].x)
		ys = append(ys, p[i].y)
		p[i].id = i + 1
	}
	sort.Slice(p, func(a, b int) bool {
		if p[a].x == p[b].x {
			return p[a].y > p[b].y
		}
		return p[a].x < p[b].x
	})
	for i := 0; i < n; i++ {
		if s < 0 && p[i].id == -s {
			s = i
		}
		if t < 0 && p[i].id == -t {
			t = i
		}
	}
	if s > t {
		s, t = t, s
	}
	sort.Ints(xs)
	xs = unique(xs)
	sort.Ints(ys)
	ys = unique(ys)
	var segyx, segx, segy Segtree
	segyx.Init(len(ys))
	segx.Init(len(xs))
	segy.Init(len(ys))
	for i := 0; i < n; i++ {
		p[i].x = lowerBound(xs, p[i].x)
		p[i].y = lowerBound(ys, p[i].y)
		xL := segyx.Prod(0, p[i].y+1)
		cost := min(segx.Prod(xL, p[i].x), segy.Prod(0, p[i].y+1)) + p[i].c
		if i == s {
			cost = 0
		}
		ans[i] = cost
		segx.Set(p[i].x, cost+p[i].c)
		segy.Set(p[i].y, cost)
		segyx.Set(p[i].y, p[i].x)
	}
	Min := ans[t]
	for i := 0; i < n; i++ {
		if p[t].x <= p[i].x && p[t].y <= p[i].y {
			Min = min(Min, ans[i]+p[i].c)
		}
	}
	fmt.Println(Min)
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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
