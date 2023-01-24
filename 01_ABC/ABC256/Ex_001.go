package main

import (
	"bufio"
	"fmt"
	"os"
)

const MAXN = 500005

var b [MAXN]int
var sum, lazy [MAXN * 4]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &b[i])
	}

	build(1, n, 1)

	var op, l, r, x int
	for m > 0 {
		fmt.Fscan(in, &op, &l, &r)
		switch op {
		case 1:
			fmt.Fscan(in, &x)
			modify1(1, n, l, r, x, 1)
		case 2:
			fmt.Fscan(in, &x)
			modify2(1, n, l, r, x, 1)
		case 3:
			fmt.Fprintln(out, query(1, n, l, r, 1))
		}
		m--
	}
}

func query(x, y, l, r, p int) int {
	if l <= x && r >= y {
		return sum[p]
	}
	pushdown(p, x, y)
	ans := 0
	if l <= mid(x, y) {
		ans += query(x, mid(x, y), l, r, lson(p))
	}
	if r > mid(x, y) {
		ans += query(mid(x, y)+1, y, l, r, rson(p))
	}
	return ans
}

func modify2(x, y, l, r, v, p int) {
	if l <= x && r >= y {
		push(p, y-x+1, v)
		return
	}
	pushdown(p, x, y)
	if l <= mid(x, y) {
		modify2(x, mid(x, y), l, r, v, lson(p))
	}
	if r > mid(x, y) {
		modify2(mid(x, y)+1, y, l, r, v, rson(p))
	}
	sum[p] = sum[lson(p)] + sum[rson(p)]
}

func modify1(x, y, l, r, v, p int) {
	if sum[p] == 0 {
		return
	}
	if x == y {
		sum[p] /= v
		return
	}
	if l <= x && r >= y && lazy[p] != -1 {
		lazy[p] /= v
		sum[p] = lazy[p] * (y - x + 1)
		return
	}
	pushdown(p, x, y)
	if l <= mid(x, y) {
		modify1(x, mid(x, y), l, r, v, lson(p))
	}
	if r > mid(x, y) {
		modify1(mid(x, y)+1, y, l, r, v, rson(p))
	}
	sum[p] = sum[lson(p)] + sum[rson(p)]
}

func build(x, y, p int) {
	if x == y {
		sum[p] = b[x]
		return
	}
	build(x, mid(x, y), lson(p))
	build(mid(x, y)+1, y, rson(p))
	sum[p] = sum[lson(p)] + sum[rson(p)]
	lazy[p] = -1
}

func pushdown(p, x, y int) {
	if lazy[p] != -1 {
		push(lson(p), mid(x, y)-x+1, lazy[p])
		push(rson(p), y-mid(x, y), lazy[p])
		lazy[p] = -1
	}
}

func push(p, siz, x int) {
	sum[p] = siz * x
	lazy[p] = x
}

func mid(x, y int) int {
	return (x + y) >> 1
}

func lson(x int) int {
	return x << 1
}

func rson(x int) int {
	return (x << 1) | 1
}
