package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var g [][]int

var out = bufio.NewWriter(os.Stdout)
var in = bufio.NewReader(os.Stdin)

func main() {
	defer out.Flush()

	var n int
	fmt.Scan(&n)

	g = make([][]int, n+1)
	var a, b int
	for i := 0; i < n-1; i++ {
		fmt.Fscan(in, &a, &b)
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	for i := 1; i <= n; i++ {
		sort.Ints(g[i])
	}

	dfs(1, 0)
}

func dfs(u, f int) {
	fmt.Fprint(out, u, " ")
	for _, v := range g[u] {
		if v != f {
			dfs(v, u)
			fmt.Fprint(out, u, " ")
		}
	}
}
