package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	N   int
	vis = make([]int, 105)
	adj = [105][105]bool{}
)

func dfs(x int) {
	vis[x] = 1
	for i := 1; i <= N; i++ {
		if adj[x][i] && vis[i] == 0 {
			dfs(i)
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var Q int
	fmt.Fscan(in, &N, &Q)

	for i := 0; i < Q; i++ {
		var x, u, v int
		fmt.Fscan(in, &x, &u, &v)
		if x == 1 {
			adj[u][v] = !adj[u][v]
			adj[v][u] = !adj[v][u]
		} else {
			vis = make([]int, 105)
			dfs(u)
			if vis[v] != 0 {
				fmt.Fprintln(out, "Yes")
			} else {
				fmt.Fprintln(out, "No")
			}
		}
	}
}
