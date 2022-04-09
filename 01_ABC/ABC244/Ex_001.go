package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const inf = 1 << 62

type pair struct{ x, y int }
type pair2 struct {
	x pair
	y int
}

var (
	x   = [200005]int{}
	y   = [200005]int{}
	all = make([]pair, 200005)
	tr1 lct1
	tr2 lct2
)

func F(i, p int) int {
	return x[i]*all[p].x + y[i]*all[p].y
}

func (lct *lct1) upd(i, l, r, a int) {
	if lct.d[i] == 0 {
		lct.d[i] = a
		return
	}
	m := (l + r) >> 1
	b := lct.d[i]
	resa := F(a, m)
	resb := F(b, m)
	if resa > resb {
		a, b = b, a
		lct.d[i] = b
	}
	if l == r {
		return
	}
	if x[a] < x[b] {
		lct.upd(i<<1, l, m, a)
	} else {
		lct.upd(i<<1|1, m+1, r, a)
	}
}

func (lct *lct1) qry(i, l, r, p int) int {
	if lct.d[i] == 0 {
		return -inf
	}
	ret := F(lct.d[i], p)
	if l != r {
		m := (l + r) >> 1
		if p <= m {
			ret = max(ret, lct.qry(i<<1, l, m, p))
		} else {
			ret = max(ret, lct.qry(i<<1|1, m+1, r, p))
		}
	}
	return ret
}

type lct1 struct{ d [525000]int }

func (lct *lct2) upd(i, l, r, a int) {
	if lct.d[i] == 0 {
		lct.d[i] = a
		return
	}
	m := (l + r) >> 1
	b := lct.d[i]
	resa := F(a, m)
	resb := F(b, m)
	if resa < resb {
		a, b = b, a
		lct.d[i] = b
	}
	if l == r {
		return
	}
	if x[a] < x[b] {
		lct.upd(i<<1|1, m+1, r, a)
	} else {
		lct.upd(i<<1, l, m, a)
	}
}

func (lct *lct2) qry(i, l, r, p int) int {
	if lct.d[i] == 0 {
		return inf
	}
	ret := F(lct.d[i], p)
	if l != r {
		m := (l + r) >> 1
		if p <= m {
			ret = min(ret, lct.qry(i<<1, l, m, p))
		} else {
			ret = min(ret, lct.qry(i<<1|1, m+1, r, p))
		}
	}
	return ret
}

type lct2 struct{ d [525000]int }

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var q int
	fmt.Fscan(in, &q)
	a := make([]int, q+1)
	b := make([]int, q+1)
	v := make([]pair2, 0)
	for i := 1; i <= q; i++ {
		fmt.Fscan(in, &x[i], &y[i], &a[i], &b[i])
		if b[i] == 0 {
			continue
		}
		if b[i] > 0 {
			v = append(v, pair2{pair{a[i], b[i]}, i})
		} else {
			v = append(v, pair2{pair{-a[i], -b[i]}, i})
		}
	}
	sort.Slice(v, func(i, j int) bool {
		return v[i].x.x*v[j].x.y < v[i].x.y*v[j].x.x
	})

	pos := make([]int, 200005)
	N := len(v)
	for i := 0; i < len(v); i++ {
		pos[v[i].y] = i
		all[i] = v[i].x
	}
	mnA := int(1e9)
	mxA := -int(1e9)
	for i := 1; i <= q; i++ {
		if N > 0 {
			tr1.upd(1, 0, N-1, i)
			tr2.upd(1, 0, N-1, i)
		}
		mnA = min(mnA, x[i])
		mxA = max(mxA, x[i])
		if b[i] == 0 {
			if a[i] < 0 {
				fmt.Fprintln(out, a[i]*mnA)
			} else {
				fmt.Fprintln(out, a[i]*mxA)
			}
			continue
		}
		var h int
		if b[i] > 0 {
			h = tr1.qry(1, 0, N-1, pos[i])
		} else {
			h = -tr2.qry(1, 0, N-1, pos[i])
		}
		fmt.Fprintln(out, h)
	}
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
