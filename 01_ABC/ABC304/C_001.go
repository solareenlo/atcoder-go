package main

import (
	"bufio"
	"fmt"
	"os"
)

var n, d int
var vis [2020]bool
var x, y [2020]int

func dfs(u int) {
	vis[u] = true
	for v := 1; v <= n; v++ {
		if !vis[v] && (x[u]-x[v])*(x[u]-x[v])+(y[u]-y[v])*(y[u]-y[v]) <= d {
			dfs(v)
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &n, &d)
	d *= d
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &x[i], &y[i])
	}
	dfs(1)
	for i := 1; i <= n; i++ {
		if vis[i] {
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
		}
	}
}
