package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	c = [200002]int{}
	p = [200002]int{}
	g = make([][]int, 2<<17)
)

func dfs(u, v int) {
	p[u] = v
	for _, e := range g[u] {
		if e != v {
			c[e] += c[u]
			dfs(e, u)
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n)
	b := make([]int, n)
	for i := 1; i < n; i++ {
		fmt.Fscan(in, &a[i], &b[i])
		g[a[i]] = append(g[a[i]], b[i])
		g[b[i]] = append(g[b[i]], a[i])
	}

	var q int
	fmt.Fscan(in, &q)

	dfs(1, 0)

	for i := 0; i < q; i++ {
		var t, e, x int
		fmt.Fscan(in, &t, &e, &x)
		u := a[e]
		v := b[e]
		if t == 2 {
			u, v = v, u
		}
		if p[u] == v {
			c[u] += x
		} else {
			c[1] += x
			c[v] -= x
		}
	}

	dfs(1, 0)

	for i := 1; i < n+1; i++ {
		fmt.Fprintln(out, c[i])
	}
}
