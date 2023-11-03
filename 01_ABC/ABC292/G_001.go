package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod = 998244353
const N = 45

var m int
var dp [N][N][N][10]int
var s [N]string

func dfs(l, r, t, d int) int {
	if d > 9 || d < 0 {
		return 0
	}
	if dp[l][r][t][d] != -1 {
		return dp[l][r][t][d]
	}
	if t == m+1 {
		if l < r {
			dp[l][r][t][d] = 0
		} else {
			dp[l][r][t][d] = 1
		}
		return dp[l][r][t][d]
	}
	dp[l][r][t][d] = 0
	dp[l][r][t][d] = dfs(l, r, t, d+1)
	for i := l; i <= r; i++ {
		if s[i][t] != '?' && s[i][t] != '0'+byte(d) {
			break
		}
		if i == r {
			dp[l][r][t][d] = (dp[l][r][t][d] + dfs(l, r, t+1, 0)) % mod
		} else {
			dp[l][r][t][d] = (dp[l][r][t][d] + dfs(l, i, t+1, 0)*dfs(i+1, r, t, d+1)) % mod
		}
	}
	return dp[l][r][t][d]
}

func main() {
	in := bufio.NewReader(os.Stdin)

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			for k := 0; k < N; k++ {
				for l := 0; l < 10; l++ {
					dp[i][j][k][l] = -1
				}
			}
		}
	}

	var n int
	fmt.Fscan(in, &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &s[i])
		s[i] = " " + s[i]
	}
	fmt.Println(dfs(1, n, 1, 0))
}
