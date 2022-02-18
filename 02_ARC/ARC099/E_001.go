package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 705

func Get(x int) int {
	return x * (x - 1) >> 1
}

var (
	num = [3]int{}
	n   int
	m   int
	col = [N]int{}
	dp  = [N][N]bool{}
	g   = [N][N]bool{}
)

func dfs(x, c int) bool {
	col[x] = c
	num[c]++
	for i := 1; i <= n; i++ {
		if !g[x][i] {
			if col[i] == 0 {
				if !dfs(i, 3-c) {
					return false
				}
			} else if col[i] == c {
				return false
			}
		}
	}
	return true
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n, &m)
	for i := 1; i <= m; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		g[u][v] = true
		g[v][u] = true
	}

	dp[0][0] = true
	for i := 1; i <= n; i++ {
		g[i][i] = true
	}

	now := 0
	for i := 1; i <= n; i++ {
		if col[i] == 0 {
			num[1] = 0
			num[2] = 0
			if !dfs(i, 1) {
				fmt.Println(-1)
				return
			}
			now++
			for j := 0; j <= n; j++ {
				if dp[now-1][j] {
					dp[now][j+num[1]] = true
					dp[now][j+num[2]] = true
				}
			}
		}
	}

	ans := 1 << 60
	for i := 0; i <= n; i++ {
		if dp[now][i] {
			ans = min(ans, Get(i)+Get(n-i))
		}
	}

	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
