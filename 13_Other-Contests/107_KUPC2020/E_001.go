package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 600006

var n int
var a, b, c, s [N]int

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = int(1e18)

	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		s[i] = s[i-1] + a[i]
	}
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &b[i])
	}
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &c[i])
	}

	l, r := -INF, INF
	res := 0
	for l <= r {
		m := (l + r) >> 1
		if chk(m) {
			res = m
			l = m + 1
		} else {
			r = m - 1
		}
	}
	fmt.Println(res)
}

func chk(lim int) bool {
	mx := b[1] - s[0]
	for i := 1; i < n; i++ {
		if mx+c[i]+s[i] >= lim {
			mx = max(mx, b[i+1]-s[i])
		}
	}
	if mx+c[n]+s[n] >= lim {
		return true
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
