package main

import (
	"bufio"
	"fmt"
	"os"
)

var n, ans int
var vis [15]bool
var a [15][15]int

func dfs(k, sum int) {
	vis[k] = true
	for i := 1; i <= n; i++ {
		if a[k][i] != 0 && !vis[i] {
			dfs(i, sum+a[k][i])
		}
	}
	vis[k] = false
	ans = max(ans, sum)
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var m int
	fmt.Fscan(in, &n, &m)
	for i := 1; i <= m; i++ {
		var u, v, c int
		fmt.Fscan(in, &u, &v, &c)
		a[u][v] = c
		a[v][u] = c
	}
	for i := 1; i <= n; i++ {
		dfs(i, 0)
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
