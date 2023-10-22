package main

import (
	"bufio"
	"fmt"
	"os"
)

var r, fa, ti [100100]int

func find(t, x int) int {
	if ti[x] > t {
		return x
	}
	return find(t, fa[x])
}

func uni(x, y, t int) {
	x = find(t, x)
	y = find(t, y)
	if x == y {
		return
	}
	if r[x] < r[y] {
		x, y = y, x
	}
	fa[y] = x
	ti[y] = t
	if r[x] == r[y] {
		r[x]++
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 0; i < n+1; i++ {
		fa[i] = i
		ti[i] = (1 << 60)
	}
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		uni(a, b, i+1)
	}

	var q int
	fmt.Fscan(in, &q)
	for q > 0 {
		q--
		var a, b int
		fmt.Fscan(in, &a, &b)
		if find(m, a) != find(m, b) {
			fmt.Fprintln(out, -1)
			continue
		}
		s, g := 0, m
		for g > 1+s {
			mid := (s + g) / 2
			if find(mid, a) == find(mid, b) {
				g = mid
			} else {
				s = mid
			}
		}
		fmt.Fprintln(out, g)
	}
}
