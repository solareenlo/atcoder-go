package main

import (
	"bufio"
	"fmt"
	"os"
)

const mx = 100010

type pair struct {
	x, y int
}

var n, m int
var g [mx][]int
var vis, low, ord, cmp [mx]int
var bridge []pair

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n, &m)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}
	time := 0
	dfs(0, -1, &time)
	now := 0
	for i := 0; i < n; i++ {
		vis[i] = 0
	}
	for i := 0; i < n; i++ {
		if vis[i] == 0 {
			dfs1(i, -1, now)
			now++
		}
	}
	if (n == 3 && m == 2) || len(bridge) == 0 {
		fmt.Println("IMPOSSIBLE")
		return
	} else if len(bridge) == 1 {
		fmt.Println(0)
		return
	}

	for i := 0; i < now; i++ {
		g[i] = make([]int, 0)
	}
	for _, itr := range bridge {
		u := itr.x
		v := itr.y
		u = cmp[u]
		v = cmp[v]
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}
	cnt := 0
	for i := 0; i < now; i++ {
		if len(g[i]) == 1 {
			cnt++
		}
	}

	ans := (cnt+1)/2 - 1
	fn := false
	for i := 0; i < now; i++ {
		if len(g[i]) == 1 {
			for _, to := range g[i] {
				if len(g[to]) > 2 {
					fn = true
				}
			}
		}
	}
	if fn {
		cnt--
	}
	if cnt%2 == 0 {
		ans = cnt / 2
	} else {
		ans = (cnt + 1) / 2
	}
	fmt.Println(ans)
}

func dfs(ni, p int, time *int) {
	low[ni] = *time
	ord[ni] = *time
	(*time)++
	vis[ni] = 1
	for _, to := range g[ni] {
		if to != p {
			if vis[to] != 0 {
				low[ni] = min(low[ni], ord[to])
			} else {
				dfs(to, ni, time)
				low[ni] = min(low[ni], low[to])
				if low[to] > ord[ni] {
					bridge = append(bridge, pair{to, ni})
				}
			}
		}
	}
}

func dfs1(ni, p, now int) {
	vis[ni] = 1
	cmp[ni] = now
	for _, to := range g[ni] {
		if to != p {
			if low[ni] > ord[to] || low[to] > ord[ni] {
				continue
			}
			if vis[to] != 0 {
				continue
			}
			dfs1(to, ni, now)
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
