package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	l := make([]int, n+2)
	r := make([]int, n+2)
	inf := 1 << 62
	for i := 1; i <= n; i++ {
		l[i] = inf
		r[i] = -inf
	}

	for i := 1; i <= n; i++ {
		var x, c int
		fmt.Fscan(in, &x, &c)
		l[c] = min(l[c], x)
		r[c] = max(r[c], x)
	}

	f := make([]int, n+2)
	g := make([]int, n+2)
	for i, la := 1, 0; i <= n+1; i++ {
		if l[i] == inf {
			continue
		}
		f[i] = min(f[la]+abs(l[la]-r[i]), g[la]+abs(r[la]-r[i])) + r[i] - l[i]
		g[i] = min(f[la]+abs(l[la]-l[i]), g[la]+abs(r[la]-l[i])) + r[i] - l[i]
		la = i
	}

	fmt.Println(f[n+1])
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
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
