package main

import (
	"bufio"
	"fmt"
	"os"
)

var N, ans int
var G, C [100009][]int
var sha [100009]int
var used [100009]bool

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &N)

	for i := 1; i <= N-1; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		G[a] = append(G[a], b)
		G[b] = append(G[b], a)
	}

	dfs(1)
	dfs2(1)

	fmt.Println(ans)
}

func dfs(v int) {
	used[v] = true
	for _, vv := range G[v] {
		if used[vv] {
			continue
		}
		C[v] = append(C[v], vv)
		dfs(vv)
	}
}

func dfs2(v int) {
	for _, vv := range C[v] {
		dfs2(vv)
		sha[v] += sha[vv]
	}
	sha[v]++

	ans += sha[v] * (N - sha[v])
}
