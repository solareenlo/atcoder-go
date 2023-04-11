package main

import (
	"bufio"
	"fmt"
	"os"
)

var dp [100005]int
var e [100005][]int

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	for i := 1; i <= m; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		e[u] = append(e[u], v)
	}

	ans := 0
	for i := 1; i <= n; i++ {
		ans = max(ans, dfs(i))
	}
	fmt.Println(ans)
}

func dfs(u int) int {
	if dp[u] == 0 {
		for _, i := range e[u] {
			dp[u] = max(dfs(i)+1, dp[u])
		}
	}
	return dp[u]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
