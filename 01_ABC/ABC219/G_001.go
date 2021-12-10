package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	const M = 510
	type pair struct{ va, ti int }

	var n, m, q int
	fmt.Fscan(in, &n, &m, &q)

	v := make([]pair, n)
	for i := 0; i < n; i++ {
		v[i] = pair{i + 1, 0}
	}
	c := make([]pair, n)
	for i := range c {
		c[i] = pair{0, -1}
	}

	g := make([][]int, n)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	r := make([][]int, n)
	for i := 0; i < n; i++ {
		if len(g[i]) >= M {
			for _, p := range g[i] {
				r[p] = append(r[p], i)
			}
		}
	}

	for i := 0; i < q; i++ {
		var x int
		fmt.Fscan(in, &x)
		x--
		for _, p := range r[x] {
			if v[x].ti < c[p].ti {
				v[x] = c[p]
			}
		}
		v[x].ti = i + 1
		if len(g[x]) < M {
			for _, p := range g[x] {
				v[p] = v[x]
			}
		} else {
			c[x] = v[x]
		}
	}

	for i := 0; i < n; i++ {
		for _, p := range r[i] {
			if v[i].ti < c[p].ti {
				v[i] = c[p]
			}
		}
	}

	for i := 0; i < n; i++ {
		fmt.Fprint(out, v[i].va, " ")
	}
	fmt.Fprintln(out)
}
