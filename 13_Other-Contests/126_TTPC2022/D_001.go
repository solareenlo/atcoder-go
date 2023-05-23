package main

import (
	"bufio"
	"fmt"
	"os"
)

const MAXN = 100010

var dp [MAXN][2]int
var c []int
var g [MAXN][]int

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	c = make([]int, n+1)
	for i := 1; i < n+1; i++ {
		fmt.Fscan(in, &c[i])
	}

	for i := 1; i < n; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}
	dfs(1, 0)
	fmt.Println(max(dp[1][0], dp[1][1]))
}

func dfs(u, f int) {
	if len(g[u]) > 1 || g[u][0] != f {
		dp[u][1] = -MAXN
	}
	for _, v := range g[u] {
		if v == f {
			continue
		}
		dfs(v, u)
		ndp := [2]int{max(dp[u][0]+dp[v][0], dp[u][1]+dp[v][1]), max(dp[u][1]+dp[v][0], dp[u][0]+dp[v][1])}
		dp[u][0] = ndp[0]
		dp[u][1] = ndp[1]
	}
	dp[u][c[u]^1]++
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
