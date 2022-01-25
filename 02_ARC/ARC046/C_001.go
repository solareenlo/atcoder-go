package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const MX = 150005
const INF = 1 << 60

type P struct{ x, y int }

var (
	a = make([]P, MX)
	b = make([]P, MX)
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	a = make([]P, n)
	b = make([]P, m)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i].x, &a[i].y)
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &b[i].x, &b[i].y)
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i].x < a[j].x
	})
	sort.Slice(b, func(i, j int) bool {
		return b[i].x < b[j].x
	})

	t := newsegtree(m + 5)
	for i := 0; i < m; i++ {
		t.d[t.x2+i] = b[i].y
	}
	t.fil()
	ans := 0
	for i := 0; i < n; i++ {
		l := lowerBound(b, P{a[i].y, -1})
		l--
		r := m
		for l+1 < r {
			c := (l + r) >> 1
			if t.get(l+1, c+1, 1, 0, -1) <= a[i].x {
				r = c
			} else {
				l = c
			}
		}
		if r == m {
			continue
		}
		ans++
		t.add(r, INF)
	}
	fmt.Println(ans)
}

func lowerBound(a []P, x P) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i].x >= x.x
	})
	return idx
}

type segtree struct {
	d  []int
	x2 int
}

func newsegtree(mx int) *segtree {
	seg := new(segtree)
	seg.x2 = 1
	for seg.x2 < mx {
		seg.x2 <<= 1
	}
	seg.d = make([]int, seg.x2<<1)
	for i := range seg.d {
		seg.d[i] = INF
	}
	return seg
}

func (seg *segtree) fil() {
	for i := seg.x2 - 1; i >= 0; i-- {
		seg.d[i] = min(seg.d[i<<1], seg.d[i<<1|1])
	}
}

func (seg *segtree) add(i, x int) {
	i += seg.x2
	seg.d[i] = x
	i >>= 1
	for ; i > 0; i >>= 1 {
		seg.d[i] = min(seg.d[i<<1], seg.d[i<<1|1])
	}
}

func (seg *segtree) get(a, b, i, l, r int) int {
	if r == -1 {
		r = seg.x2
	}
	if a <= l && r <= b {
		return seg.d[i]
	}
	c := (l + r) >> 1
	res := INF
	if a < c {
		res = min(res, seg.get(a, b, i<<1, l, c))
	}
	if c < b {
		res = min(res, seg.get(a, b, i<<1|1, c, r))
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
