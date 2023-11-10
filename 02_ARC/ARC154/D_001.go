package main

import (
	"bufio"
	"fmt"
	"os"
)

const maxn = 2020

var out = bufio.NewWriter(os.Stdout)
var loc int
var a, b [maxn]int

func ask(i, j, k int) bool {
	fmt.Fprintln(out, "?", i, j, k)
	out.Flush()
	var s string
	fmt.Scan(&s)
	if s == "-1" {
		os.Exit(0)
	}
	return s == "Yes"
}

func cmp(x, y int) bool {
	if x == loc {
		return true
	}
	if y == loc {
		return false
	}
	return !ask(x, loc, y)
}

func cdq(l, r int) {
	if l == r {
		return
	}
	mid := (l + r) >> 1
	cdq(l, mid)
	cdq(mid+1, r)
	c1 := l
	c2 := mid + 1
	cnt := 0
	for c1 <= mid || c2 <= r {
		if c2 > r {
			cnt++
			b[cnt] = a[c1]
			c1++
		} else if c1 > mid {
			cnt++
			b[cnt] = a[c2]
			c2++
		} else if cmp(a[c1], a[c2]) {
			cnt++
			b[cnt] = a[c1]
			c1++
		} else {
			cnt++
			b[cnt] = a[c2]
			c2++
		}
	}
	for i, j := l, 1; i <= r; i, j = i+1, j+1 {
		a[i] = b[j]
	}
}

func main() {
	var ans [maxn]int

	var n int
	fmt.Scan(&n)
	loc = 1
	for i := 2; i <= n; i++ {
		if ask(loc, loc, i) {
			loc = i
		}
	}
	for i := 1; i <= n; i++ {
		a[i] = i
	}
	cdq(1, n)
	for i := 1; i <= n; i++ {
		ans[a[i]] = i
	}
	fmt.Fprint(out, "!")
	out.Flush()
	for i := 1; i <= n; i++ {
		fmt.Fprintf(out, " %d", ans[i])
		out.Flush()
	}
	fmt.Fprintln(out)
	out.Flush()
}
