package main

import (
	"bufio"
	"fmt"
	"os"
)

type pair struct{ x, y int }

var (
	G  = [5001][]int{}
	dp = make([][]pair, 5001)
	ss = [5001]int{}
	a  = [5001]int{}
)

func dfs2(u, p int) {
	ss[u] = 1
	dp[u] = make([]pair, ss[u])
	if a[u] > 0 {
		dp[u][0] = pair{a[u], 1}
	} else {
		dp[u][0] = pair{a[u], 0}
	}
	for _, v := range G[u] {
		if v != p {
			dfs2(v, u)
			m := make([]pair, ss[u]+ss[v])
			for i := range m {
				m[i] = pair{1 << 60, 0}
			}
			for i := 0; i < ss[u]; i++ {
				for j := 0; j < ss[v]; j++ {
					m[i+j].x = min(m[i+j].x, dp[u][i].x+dp[v][j].x)
					if dp[u][i].y != 0 && dp[v][j].y != 0 {
						m[i+j].y = 1
					}
					if dp[v][j].y != 0 || dp[v][j].x < 0 {
						m[i+j+1].x = min(m[i+j+1].x, dp[u][i].x)
						if dp[u][i].y != 0 {
							m[i+j+1].y = 1
						}
					}
				}
			}
			ss[u] += ss[v]
			dp[u] = m
		}
	}
}
func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 1; i < n; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		G[u] = append(G[u], v)
		G[v] = append(G[v], u)
	}
	dfs2(0, -1)
	res := n
	for i := 0; i < n; i++ {
		if d[0][i].y != 0 || d[0][i].x < 0 {
			res = min(res, i)
		}
	}
	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
