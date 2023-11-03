package main

import (
	"bufio"
	"fmt"
	"os"
)

func cross(x1, y1, x2, y2 int) int {
	return x1*y2 - x2*y1
}

var n, q int
var x, y [200005]int

func solve(a, b int) int {
	o1 := cross(a-x[0], b-y[0], x[n-1]-x[0], y[n-1]-y[0])
	o2 := cross(a-x[0], b-y[0], x[1]-x[0], y[1]-y[0])
	if !(o1 >= 0 && o2 <= 0) {
		return 1
	}
	if (o1 == 0 && min(x[0], x[n-1]) <= a && a <= max(x[0], x[n-1]) && min(y[0], y[n-1]) <= b && b <= max(y[0], y[n-1])) || (o2 == 0 && min(x[0], x[1]) <= a && a <= max(x[0], x[1]) && min(y[0], y[1]) <= b && b <= max(y[0], y[1])) {
		return 0
	}
	l, r, ans := 1, n-2, 1
	for l <= r {
		mid := (l + r) >> 1
		o := cross(a-x[0], b-y[0], x[mid]-x[0], y[mid]-y[0])
		if o <= 0 {
			l = mid + 1
			ans = mid
		} else {
			r = mid - 1
		}
	}
	o := cross(a-x[ans], b-y[ans], x[(ans+1)%n]-x[ans], y[(ans+1)%n]-y[ans])
	if o == 0 {
		return 0
	}
	if o > 0 {
		return 1
	}
	return -1
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x[i], &y[i])
	}
	fmt.Fscan(in, &q)
	for q > 0 {
		q--
		var a, b int
		fmt.Fscan(in, &a, &b)
		op := solve(a, b)
		if op == 0 {
			fmt.Fprintln(out, "ON")
		} else if op == 1 {
			fmt.Fprintln(out, "OUT")
		} else {
			fmt.Fprintln(out, "IN")
		}
	}
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
