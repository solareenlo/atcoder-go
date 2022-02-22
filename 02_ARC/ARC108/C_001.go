package main

import (
	"bufio"
	"fmt"
	"os"
)

type pair struct{ i, w int }

var (
	adj = make([][]pair, 100005)
	c   = make([]int, 100005)
)

func dfs(cur int) {
	for j := range adj[cur] {
		i := adj[cur][j].i
		w := adj[cur][j].w
		if c[i] == 0 {
			if w == c[cur] {
				tmp := 0
				if w == 1 {
					tmp = 1
				}
				c[i] = 1 + tmp
			} else {
				c[i] = w
			}
			dfs(i)
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	for i := 0; i < m; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		adj[a] = append(adj[a], pair{b, c})
		adj[b] = append(adj[b], pair{a, c})
	}

	c[1] = 1
	dfs(1)
	for i := 1; i <= n; i++ {
		fmt.Fprintln(out, c[i])
	}
}
