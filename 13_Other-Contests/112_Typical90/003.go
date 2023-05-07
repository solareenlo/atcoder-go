package main

import (
	"bufio"
	"fmt"
	"os"
)

var g [][]int

var maxd, dl int

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	g = make([][]int, n+1)
	for i := 0; i < n-1; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}
	dfs(1, 1, -1)
	dfs(dl, 1, -1)
	fmt.Println(maxd)
}

func dfs(u, depth, p int) {
	for i := 0; i < len(g[u]); i++ {
		nx := g[u][i]
		if nx != p {
			dfs(nx, depth+1, u)
		}
	}
	if depth > maxd {
		maxd = depth
		dl = u
	}
}
