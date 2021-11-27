package main

import "fmt"

var (
	mod = 1000000007
	g   = make([][]int, 2020)
	dp  = [2020][2020][3]int{}
)

func dfs(u, v int) int {
	s := 1
	dp[v][0][0] = 1
	dp[v][1][1] = 1
	dp[v][1][2] = 1
	for _, w := range g[v] {
		if w != u {
			t := dfs(v, w)
			for k := s; k >= 0; k-- {
				for l := 1; l < t+1; l++ {
					dp[v][k+l][0] += dp[v][k][0] * ((dp[w][l][0] + dp[w][l][1] - dp[w][l-1][0] + mod) % mod) % mod
					dp[v][k+l][0] %= mod
					dp[v][k+l][1] += dp[v][k][1] * ((dp[w][l][0] + dp[w][l][1] + dp[w][l][2] - dp[w][l-1][0] + mod) % mod) % mod
					dp[v][k+l][1] %= mod
					dp[v][k+l][2] += dp[v][k][2] * (dp[w][l][1] + dp[w][l][2]) % mod
					dp[v][k+l][2] %= mod
				}
				dp[v][k][2] = 0
			}
			s += t
			s %= mod
		}
	}
	return s
}

func main() {
	var n int
	fmt.Scan(&n)
	for i := 1; i < n; i++ {
		var u, v int
		fmt.Scan(&u, &v)
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	dfs(0, 1)

	for k := 0; k < n+1; k++ {
		if k != 0 {
			fmt.Println((dp[1][k][0] + dp[1][k][1] + dp[1][k][2] - dp[1][k-1][0] + mod) % mod)
		} else {
			fmt.Println((dp[1][k][0] + dp[1][k][1] + dp[1][k][2]) % mod)
		}
	}
}
