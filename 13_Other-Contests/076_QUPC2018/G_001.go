package main

import (
	"bufio"
	"fmt"
	"os"
)

type pair struct {
	x, y int
}

var G [1 << 19][]pair
var C [1 << 19]int
var dp [1 << 19][2]int

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, x, y int
	fmt.Fscan(in, &n, &x, &y)
	for i := 0; i < n-1; i++ {
		var a, b, v int
		fmt.Fscan(in, &a, &b, &v)
		a--
		b--
		G[a] = append(G[a], pair{b, v})
		G[b] = append(G[b], pair{a, v})
	}
	for i := 0; i < n; i++ {
		C[i] = -1
	}
	for i := 0; i < x+y; i++ {
		var u int
		fmt.Fscan(in, &u)
		u--
		if i < x {
			C[u] = 1
		} else {
			C[u] = 0
		}
	}
	dfs(0, -1)
	fmt.Println(min(dp[0][0], dp[0][1]))
}

func dfs(u, p int) {
	if C[u] != -1 {
		dp[u][C[u]] = 1 << 60
	}
	for _, i := range G[u] {
		v := i.x
		d := i.y
		if v != p {
			dfs(v, u)
			for j := 0; j < 2; j++ {
				dp[u][j] += min(dp[v][j], dp[v][j^1]+d)
			}
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
