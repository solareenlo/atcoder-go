package main

import (
	"bufio"
	"fmt"
	"os"
)

var w, v [100]int
var G [100][]int
var vis [100]bool

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M, W int
	fmt.Fscan(in, &N, &M, &W)

	for i := 0; i < N; i++ {
		fmt.Fscan(in, &w[i], &v[i])
	}
	for i := 0; i < M; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		G[a] = append(G[a], b)
		G[b] = append(G[b], a)
	}

	var dp [10101]int
	for i := 0; i < N; i++ {
		if vis[i] {
			continue
		}
		p := dfs(i)
		for j := W - p.x; j >= 0; j-- {
			dp[j+p.x] = max(dp[j+p.x], dp[j]+p.y)
		}
	}
	fmt.Println(dp[W])
}

type pair struct {
	x, y int
}

func dfs(u int) pair {
	vis[u] = true
	ret := pair{w[u], v[u]}
	for _, v := range G[u] {
		if !vis[v] {
			q := dfs(v)
			ret.x += q.x
			ret.y += q.y
		}
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
