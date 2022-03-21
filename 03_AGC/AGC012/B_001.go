package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 1e5 + 3

var (
	f = [N][13]bool{}
	s = make([]int, N)
	g = make([][]int, N)
)

func work(v, d, c int) {
	if d < 0 || f[v][d] {
		return
	}
	f[v][d] = true
	if s[v] == 0 {
		s[v] = c
	}
	for _, i := range g[v] {
		work(i, d-1, c)
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	var q int
	fmt.Fscan(in, &q)
	v := make([]int, q+1)
	d := make([]int, q+1)
	c := make([]int, q+1)
	for i := 1; i <= q; i++ {
		fmt.Fscan(in, &v[i], &d[i], &c[i])
	}

	for i := q; i > 0; i-- {
		work(v[i], d[i], c[i])
	}

	for i := 1; i <= n; i++ {
		fmt.Fprintln(out, s[i])
	}
}
