package main

import (
	"bufio"
	"fmt"
	"os"
)

const maxn = 200200

type pair struct {
	x, y int
}

var s [maxn << 2]pair
var n, q int
var a, b [maxn]int

func upd(o int) {
	s[o].x = max(s[o<<1].x+s[o<<1|1].x, s[o<<1].y+s[o<<1|1].y)
	s[o].y = max(s[o<<1].y+s[o<<1|1].x, s[o<<1].x+s[o<<1|1].y)
}

func build(o, l, r int) {
	if l == r {
		s[o] = pair{b[l], a[l]}
		return
	}
	mid := (l + r) >> 1
	build(o<<1, l, mid)
	build(o<<1|1, mid+1, r)
	upd(o)
}

func mdy(o, l, r, k int, p pair) {
	if l == r {
		s[o] = p
		return
	}
	mid := (l + r) >> 1
	if k <= mid {
		mdy(o<<1, l, mid, k, p)
	} else {
		mdy(o<<1|1, mid+1, r, k, p)
	}
	upd(o)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &n, &q)
	for i := 1; i <= 2*n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 1; i <= 2*n; i++ {
		fmt.Fscan(in, &b[i])
	}
	if n > 1 {
		build(1, 2, 2*n-1)
	}
	for i := 1; i <= q; i++ {
		var p, x, y int
		fmt.Fscan(in, &p, &x, &y)
		if 2 <= p && p <= 2*n-1 {
			mdy(1, 2, 2*n-1, p, pair{y, x})
		} else {
			a[p] = x
			b[p] = y
		}
		fmt.Fprintln(out, a[1]+a[2*n]+s[1].x)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
