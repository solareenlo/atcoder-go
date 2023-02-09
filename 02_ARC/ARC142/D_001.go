package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod = 998244353
const N = 200005

var g, dp [][]int

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	dp = make([][]int, N)
	for i := range dp {
		dp[i] = make([]int, 6)
	}
	g = make([][]int, N)
	for i := 1; i < n; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}
	dfs(1, 0)
	fmt.Println((dp[1][1] + dp[1][2] + dp[1][3]) % mod)
}

func dfs(u, topf int) {
	mul1, mul2, mul3 := 1, 1, 1
	for _, v := range g[u] {
		if v != topf {
			dfs(v, u)
			dp[u][1] = (dp[u][1]*dp[v][1] + dp[u][4]*dp[v][5] + dp[u][5]*dp[v][4]) % mod
			dp[u][2] = (dp[u][2]*dp[v][3] + dp[v][4]*mul1) % mod
			dp[u][3] = (dp[u][3]*dp[v][2] + dp[v][5]*mul2) % mod
			dp[u][4] = (dp[u][4]*dp[v][1] + dp[v][4]*mul3) % mod
			dp[u][5] = (dp[u][5]*dp[v][1] + dp[v][5]*mul3) % mod
			mul1 = mul1 * dp[v][3] % mod
			mul2 = mul2 * dp[v][2] % mod
			mul3 = mul3 * dp[v][1] % mod
		}
	}
	dp[u][4] = (dp[u][4] + mul2) % mod
	dp[u][5] = (dp[u][5] + mul1) % mod
}
