package main

import (
	"bufio"
	"fmt"
	"os"
)

const MAXN = 500005

var fa, siz [MAXN]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m, e int
	fmt.Fscan(in, &n, &m, &e)

	for i := 1; i <= n; i++ {
		fa[i] = i
		siz[i] = 1
		fa[i+m] = n + m + 1
	}

	var a, b [MAXN]int
	for i := 1; i <= e; i++ {
		fmt.Fscan(in, &a[i], &b[i])
	}

	var q int
	fmt.Fscan(in, &q)
	var x, pd [MAXN]int
	for i := 1; i <= q; i++ {
		fmt.Fscan(in, &x[i])
		pd[x[i]] = 1
	}
	for i := 1; i <= e; i++ {
		if pd[i] == 0 {
			merge(a[i], b[i])
		}
	}

	var ans [MAXN]int
	for i := q; i >= 1; i-- {
		ans[i] = siz[find(n+m+1)]
		merge(a[x[i]], b[x[i]])
	}

	for i := 1; i <= q; i++ {
		fmt.Fprintln(out, ans[i])
	}

}

func find(x int) int {
	if fa[x] == x {
		return x
	}
	fa[x] = find(fa[x])
	return fa[x]
}

func merge(x, y int) {
	x = find(x)
	y = find(y)
	if x == y {
		return
	}
	if siz[x] < siz[y] {
		x, y = y, x
	}
	fa[y] = x
	siz[x] += siz[y]
}
