package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 51

var (
	mp  = [N][N]int{}
	e   = make([][]int, N)
	g   = make([][]int, N)
	fa  = make([]int, N)
	vis = make([]int, N)
)

func dfs(u, f int) {
	fa[u] = f
	if vis[f] != 0 && mp[u][f] != 0 {
		vis[u] = 1
	}
	for _, v := range g[u] {
		if v^f != 0 {
			dfs(v, u)
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var T int
	fmt.Fscan(in, &T)

	for k := 0; k < T; k++ {
		var n int
		fmt.Fscan(in, &n)
		ans := n + 1
		for i := range mp {
			for j := range mp[i] {
				mp[i][j] = 0
			}
		}
		for i := 1; i <= n; i++ {
			e[i] = e[i][:0]
			g[i] = g[i][:0]
		}
		for i := 1; i < n; i++ {
			var u, v int
			fmt.Fscan(in, &u, &v)
			mp[u][v] = 1
			mp[v][u] = 1
			e[u] = append(e[u], v)
			e[v] = append(e[v], u)
		}
		for i := 1; i < n; i++ {
			var u, v int
			fmt.Fscan(in, &u, &v)
			g[u] = append(g[u], v)
			g[v] = append(g[v], u)
		}
		ndeg := make([]int, n+1)
		for i := 1; i <= n; i++ {
			ndeg[i] = len(e[i])
		}
		for i := 1; i <= n; i++ {
			cost := 0
			for j := 0; j < n+1; j++ {
				vis[j] = 0
			}
			vis[i] = 1
			dfs(i, 0)
			deg := make([]int, n+1)
			copy(deg, ndeg)
			var u int
			for {
				for u = 1; u <= n; u++ {
					if vis[u] == 0 && deg[u] == 1 && vis[fa[u]] != 0 {
						break
					}
				}
				if u > n {
					break
				} else {
					vis[u] = 1
					cost++
				}
				for _, v := range e[u] {
					deg[v]--
				}
			}
			for u = 1; u <= n; u++ {
				if vis[u] == 0 {
					break
				}
			}
			if u > n {
				ans = min(ans, cost)
			}
			if len(e[i]) > 1 || u > n || ans <= n {
				continue
			}
			for j := 0; j < n+1; j++ {
				vis[j] = 0
			}
			vis[0] = 1
			copy(deg, ndeg)
			for {
				for u = 1; u <= n; u++ {
					if vis[u] == 0 && deg[u] <= 1 && vis[fa[u]] != 0 {
						break
					}
				}
				if u > n {
					break
				} else {
					vis[u] = 1
				}
				for _, v := range e[u] {
					deg[v]--
				}
			}
			for u = 1; u <= n; u++ {
				if vis[u] == 0 {
					break
				}
			}
			if u > n {
				ans = n
			}
		}
		if ans > n {
			fmt.Println(-1)
		} else {
			fmt.Println(ans)
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
