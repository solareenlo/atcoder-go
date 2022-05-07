package main

import (
	"bufio"
	"fmt"
	"os"
)

type pair struct{ x, y int }

var G = make([][]pair, 3000)

func dfs(u, p, d int) bool {
	if d == 0 {
		return true
	}
	for _, e := range G[u] {
		v := e.x
		if v != p && dfs(v, u, d-e.y) {
			return true
		}
	}
	return false
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, X int
	fmt.Fscan(in, &N, &X)
	for i := 1; i < N; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		a--
		b--
		G[a] = append(G[a], pair{b, c})
		G[b] = append(G[b], pair{a, c})
	}

	for i := 0; i < N; i++ {
		if dfs(i, -1, X) {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
