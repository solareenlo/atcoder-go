package main

import (
	"bufio"
	"fmt"
	"os"
)

type edge struct{ to, cost int }

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	g := make([][]edge, n+1)

	var a, b, c int
	for i := 1; i < n; i++ {
		fmt.Fscan(in, &a, &b, &c)
		g[a] = append(g[a], edge{b, c})
		g[b] = append(g[b], edge{a, c})
	}

	var q, k int
	fmt.Fscan(in, &q, &k)

	d := make([]int, n+1)
	for i := range d {
		d[i] = int(1e18)
	}
	d[k] = 0

	que := make([]int, 0)
	que = append(que, k)
	for len(que) > 0 {
		u := que[0]
		que = que[1:]
		for _, e := range g[u] {
			v := e.to
			if d[v] > d[u]+e.cost {
				d[v] = d[u] + e.cost
				que = append(que, v)
			}
		}
	}

	var x, y int
	for i := 0; i < q; i++ {
		fmt.Fscan(in, &x, &y)
		fmt.Fprintln(out, d[x]+d[y])
	}
}
