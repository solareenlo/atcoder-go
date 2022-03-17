package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 200200

var (
	n   int
	m   int
	g   = [N]int{}
	f   = [N]int{}
	b   = [N]int{}
	res int
)

func chk(x int) bool {
	mx := -1 << 60
	l := -1
	r := -1
	for i := 1; i <= n; i++ {
		if i-2 >= 0 {
			g[i] = g[i-2] + 1
		}
		tmp := 0
		if i == 1 {
			tmp = 1
			g[i] = 1
		}
		f[i] = f[i-2+tmp] + b[i] - x
		if f[i-1] == f[i] {
			g[i] = max(g[i-1], g[i])
		}
		if f[i-1] > f[i] {
			f[i] = f[i-1]
			g[i] = g[i-1]
		}
		if f[i] == mx {
			l = min(l, g[i])
			r = max(r, g[i])
		}
		if f[i] > mx {
			mx = f[i]
			l = g[i]
			r = g[i]
		}
	}
	if (l <= m) && (r >= m) {
		res = max(res, x)
	}
	return g[n] >= m
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n, &m)

	m = min(m, n-m)
	for i := 1; i < n; i++ {
		var k int
		fmt.Fscan(in, &k)
		b[i] += k
		b[i+1] += k
	}

	l := 0
	r := 1 << 60
	for l <= r {
		md := (l + r) >> 1
		if chk(md) {
			l = md + 1
		} else {
			r = md - 1
		}
	}

	chk(res)

	fmt.Println(res*m + f[n])
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
