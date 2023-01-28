package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 17

var c, dp [1 << N][18]int
var n int

func main() {
	in := bufio.NewReader(os.Stdin)

	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	fmt.Fscan(in, &n)

	m := (1 << n)
	for i := 0; i <= m-1; i++ {
		for j := 1; j <= n; j++ {
			fmt.Fscan(in, &c[i][j])
		}
	}
	fmt.Println(dfs(1, 0))

}

func dfs(u, s int) int {
	if u >= (1 << n) {
		return c[u^(1<<n)][s]
	}
	if dp[u][s] >= 0 {
		return dp[u][s]
	}
	dp[u][s] = max(dfs(u<<1, s+1)+dfs(u<<1|1, 0), dfs(u<<1, 0)+dfs(u<<1|1, s+1))
	return dp[u][s]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
