package main

import (
	"bufio"
	"fmt"
	"os"
)

type edge struct{ to, cost int }

var (
	g     [][]edge
	depth []int
)

func dfs(i, p, d int) {
	depth[i] = d
	for _, e := range g[i] {
		if e.to != p {
			dfs(e.to, i, d+e.cost)
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	g = make([][]edge, n+1)

	var a, b, c int
	for i := 1; i < n; i++ {
		fmt.Fscan(in, &a, &b, &c)
		g[a] = append(g[a], edge{b, c})
		g[b] = append(g[b], edge{a, c})
	}

	var q, k int
	fmt.Fscan(in, &q, &k)
	depth = make([]int, n+1)
	dfs(k, 0, 0)

	var x, y int
	for i := 0; i < q; i++ {
		fmt.Fscan(in, &x, &y)
		fmt.Fprintln(out, depth[x]+depth[y])
	}
}
