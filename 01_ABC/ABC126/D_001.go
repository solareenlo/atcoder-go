package main

import (
	"bufio"
	"fmt"
	"os"
)

type pair struct{ to, color int }

var G [][]pair = make([][]pair, 100001)
var res [100001]int

func dfs(now, pre, color int) {
	res[now] = color
	for _, g := range G[now] {
		if g.to != pre {
			dfs(g.to, now, color^(g.color&1))
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	var u, v, w int
	for i := 1; i < n; i++ {
		fmt.Fscan(in, &u, &v, &w)
		G[u] = append(G[u], pair{v, w})
		G[v] = append(G[v], pair{u, w})
	}
	dfs(1, 0, 0)
	for i := 1; i < n+1; i++ {
		fmt.Println(res[i])
	}
}
