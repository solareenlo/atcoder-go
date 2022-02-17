package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N = 100005

type edge struct{ x, y, w int }

var (
	e = make([]edge, N)
	f = make([]int, N)
)

func gf(k int) int {
	if f[k] != 0 {
		f[k] = gf(f[k])
		return f[k]
	}
	return k
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	a := make([]int, n+1)
	b := make([]int, n+1)
	s := make([]int, n+1)
	ss := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i], &b[i])
		s[i] = max(a[i], b[i])
		ss[i] = b[i]
	}

	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &e[i].x, &e[i].y)
		e[i].w = max(a[e[i].x]-b[e[i].x], a[e[i].y]-b[e[i].y])
	}
	tmp := e[1 : m+1]
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].w < tmp[j].w
	})

	for i := 1; i <= m; i++ {
		e[i].x = gf(e[i].x)
		e[i].y = gf(e[i].y)
		if e[i].x != e[i].y {
			s[e[i].x] = min(max(s[e[i].x], e[i].w)+ss[e[i].y], max(s[e[i].y], e[i].w)+ss[e[i].x])
			ss[e[i].x] += ss[e[i].y]
			f[e[i].y] = e[i].x
		}
	}

	fmt.Println(s[gf(1)])
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
