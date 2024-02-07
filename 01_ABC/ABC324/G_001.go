package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 200010

type D struct {
	ls, rs, x int
}

var t [N << 5]D
var cnt, n, m int
var xl, xr, yl, yr, a, rt, rk, ans [N]int

func add(p *int, q, l, r, x int) {
	cnt++
	(*p) = cnt
	t[*p].x = t[q].x + 1
	if l == r {
		return
	}
	mid := (l + r) >> 1
	if x <= mid {
		t[*p].rs = t[q].rs
		add(&(t[*p].ls), t[q].ls, l, mid, x)
	} else {
		t[*p].ls = t[q].ls
		add(&(t[*p].rs), t[q].rs, mid+1, r, x)
	}
}

func askx(p, q, l, r, x int) int {
	if l == r {
		return l
	}
	mid := (l + r) >> 1
	s := t[t[p].ls].x - t[t[q].ls].x
	if x <= s {
		return askx(t[p].ls, t[q].ls, l, mid, x)
	}
	return askx(t[p].rs, t[q].rs, mid+1, r, x-s)
}

func asky(p, q, l, r, pl, pr int) int {
	if pr < l || r < pl {
		return 0
	}
	if pl <= l && r <= pr {
		return t[p].x - t[q].x
	}
	mid := (l + r) >> 1
	s := 0
	if pl <= mid {
		s += asky(t[p].ls, t[q].ls, l, mid, pl, pr)
	}
	if pr > mid {
		s += asky(t[p].rs, t[q].rs, mid+1, r, pl, pr)
	}
	return s
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		rk[a[i]] = i
	}
	xl[0] = 1
	yl[0] = 1
	xr[0] = n
	yr[0] = n
	ans[0] = n
	for i := 1; i <= n; i++ {
		add(&rt[i], rt[i-1], 1, n, rk[i])
	}
	fmt.Fscan(in, &m)
	for i := 1; i <= m; i++ {
		var opt, id, x int
		fmt.Fscan(in, &opt, &id, &x)
		if opt == 2 {
			xl[i] = xl[id]
			xr[i] = xr[id]
			yr[i] = yr[id]
			if x < yl[id] {
				yl[i] = yl[id]
				yl[id] = yr[id] + 1
			} else if x >= yr[id] {
				yl[i] = yr[i] + 1
			} else {
				yl[i] = x + 1
				yr[id] = x
			}
		} else {
			xr[i] = xr[id]
			yl[i] = yl[id]
			yr[i] = yr[id]
			if x == 0 {
				xl[i] = xl[id]
				xl[id] = xr[id] + 1
			} else if x >= ans[id] {
				xl[i] = xr[id] + 1
			} else {
				X := asky(rt[yr[i]], rt[yl[i]-1], 1, n, 1, xl[id]-1)
				T := askx(rt[yr[i]], rt[yl[i]-1], 1, n, x+X)
				xr[id] = T
				xl[i] = T + 1
			}
		}
		ans[i] = asky(rt[yr[i]], rt[yl[i]-1], 1, n, xl[i], xr[i])
		ans[id] -= ans[i]
		fmt.Fprintln(out, ans[i])
	}
}
