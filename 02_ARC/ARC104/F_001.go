package main

import "fmt"

const N = 105
const mod = 1_000_000_007

var (
	vis = [N][N][N]bool{}
	dp  = [N][N][N]int{}
	tp  = [N]int{}
)

func dfs(l, r, v int) int {
	if l > r {
		return 1
	}
	if v == 0 {
		return 0
	}
	if vis[l][r][v] {
		return dp[l][r][v]
	}
	as := 0
	for i := l; i <= r; i++ {
		as = (as + dfs(l, i-1, min(v, tp[i]))*dfs(i+1, r, min(v, tp[i])-1)) % mod
	}
	vis[l][r][v] = true
	dp[l][r][v] = as
	return dp[l][r][v]
}

func main() {
	var n int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Scan(&tp[i])
		if tp[i] > n {
			tp[i] = n
		}
	}
	fmt.Println(dfs(1, n, n))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
