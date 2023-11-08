package main

import (
	"bufio"
	"fmt"
	"os"
)

var g [200200][]int
var siz [200200]int
var ans int = 0
var n int

func dfs(u int) {
	siz[u] = 1
	for i := 0; i < len(g[u]); i++ {
		v := g[u][i]
		if siz[v] != 0 {
			continue
		}
		dfs(v)
		siz[u] += siz[v]
		ans += siz[v] * (n - siz[v])
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)
	for i := 1; i < n; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}
	dfs(1)
	fmt.Println(n*(n-1)*(n-2)/6 - ans + n*(n-1)/2)
}
