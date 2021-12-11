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

	var n int
	fmt.Fscan(in, &n)

	g := make([][]int, n)
	for i := 0; i < n-1; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		g[a-1] = append(g[a-1], b-1)
		g[b-1] = append(g[b-1], a-1)
	}

	sub := make([]int, n)
	for i := range sub {
		sub[i] = 1
	}
	res := make([]int, n)

	var dfs func(i, p, d int)
	dfs = func(i, p, d int) {
		res[0] += d
		for _, x := range g[i] {
			if x != p {
				dfs(x, i, d+1)
				sub[i] += sub[x]
			}
		}
	}
	dfs(0, -1, 0)

	var dfs2 func(i, p int)
	dfs2 = func(i, p int) {
		for _, x := range g[i] {
			if x != p {
				res[x] = res[i] - 2*sub[x] + n
				dfs2(x, i)
			}
		}
	}
	dfs2(0, -1)

	for i := range res {
		fmt.Fprintln(out, res[i])
	}
}
