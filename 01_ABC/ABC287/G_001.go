package main

import (
	"bufio"
	"fmt"
	"os"
)

const MXN = 200200
const T = MXN << 5

var l, r, tot, p int
var sz, tr, ls, rs [T]int

func modify(x *int, l, r, p, k int) {
	if *x == 0 {
		tot++
		*x = tot
	}
	sz[*x] += k
	tr[*x] += p * k
	if l == r {
		return
	}
	if p <= (l+r)>>1 {
		modify(&ls[*x], l, (l+r)>>1, p, k)
	} else {
		modify(&rs[*x], ((l+r)>>1)+1, r, p, k)
	}
}

func query(x, l, r, k int) int {
	if x == 0 || k == 0 {
		return 0
	}
	if l == r {
		return k * l
	}
	if sz[rs[x]] >= k {
		return query(rs[x], ((l+r)>>1)+1, r, k)
	} else {
		return query(ls[x], l, (l+r)>>1, k-sz[rs[x]]) + tr[rs[x]]
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	const M = 1000000001

	var a, b [MXN]int

	var n int
	fmt.Fscan(in, &n)
	rt := 0
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i], &b[i])
		modify(&rt, 0, M, a[i], b[i])
	}
	var q int
	fmt.Fscan(in, &q)
	for q > 0 {
		q--
		var op, x int
		fmt.Fscan(in, &op, &x)
		switch op {
		case 1:
			{
				var y int
				fmt.Fscan(in, &y)
				modify(&rt, 0, M, a[x], -b[x])
				a[x] = y
				modify(&rt, 0, M, a[x], b[x])
			}
		case 2:
			{
				var y int
				fmt.Fscan(in, &y)
				modify(&rt, 0, M, a[x], -b[x])
				b[x] = y
				modify(&rt, 0, M, a[x], b[x])
			}
		case 3:
			{
				if sz[rt] < x {
					fmt.Fprintln(out, -1)
				} else {
					fmt.Fprintln(out, query(rt, 0, M, x))
				}
			}
		}
	}
}
