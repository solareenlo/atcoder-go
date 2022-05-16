package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	vis  = make([]bool, 2<<17)
	vis2 = make([]bool, 2<<17)
	G    = make([][]int, 2<<17)
)

func dfs(u int) {
	vis[u] = true
	for _, v := range G[u] {
		if !vis[v] {
			fmt.Println(u, v)
			dfs(v)
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(in, &N, &M)

	for i := 0; i < M; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		G[u] = append(G[u], v)
		G[v] = append(G[v], u)
	}

	dfs(1)
	Q := make([]int, 0)
	Q = append(Q, 1)
	vis2[1] = true
	for len(Q) > 0 {
		u := Q[0]
		Q = Q[1:]
		for _, v := range G[u] {
			if !vis2[v] {
				fmt.Println(u, v)
				vis2[v] = true
				Q = append(Q, v)
			}
		}
	}
}
