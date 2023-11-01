package main

import (
	"bufio"
	"fmt"
	"os"
)

var K int
var V [200005]bool
var G [200005][]int

func dfs(i int) {
	if K >= 1e6 {
		return
	}
	V[i] = true
	K++
	for _, j := range G[i] {
		if !V[j] {
			dfs(j)
		}
	}
	V[i] = false
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
	fmt.Println(K)
}
