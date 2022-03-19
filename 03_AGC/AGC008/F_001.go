package main

import (
	"bufio"
	"fmt"
	"os"
)

const inf = 1 << 60
const N = 200005

var (
	s   string
	mn  = make([]int, N)
	mx  = make([]int, N)
	sx  = make([]int, N)
	siz = make([]int, N)
	G   = make([][]int, N)
	ans int
)

func dfs1(u, fa int) {
	if s[u] == '1' {
		mn[u] = 0
		siz[u] = 1
	} else {
		mn[u] = inf
	}
	for _, v := range G[u] {
		if v == fa {
			continue
		}
		dfs1(v, u)
		siz[u] += siz[v]
		if mx[u] < mx[v]+1 {
			sx[u] = mx[u]
			mx[u] = mx[v] + 1
		} else {
			sx[u] = max(sx[u], mx[v]+1)
		}
		if siz[v] != 0 {
			mn[u] = min(mn[u], mx[v]+1)
		}
	}
}

func dfs2(u, fa int) {
	d := min(mx[u]-1, sx[u]+1)
	if d >= mn[u] {
		ans += d - mn[u] + 1
	}
	for _, v := range G[u] {
		if v == fa {
			continue
		}
		dis := mx[u] + 1
		if mx[u] == mx[v]+1 {
			dis = sx[u] + 1
		}
		if mx[v] < dis {
			sx[v] = mx[v]
			mx[v] = dis
		} else {
			sx[v] = max(sx[v], dis)
		}
		if siz[v] < siz[1] {
			mn[v] = min(mn[v], dis)
		}
		dfs2(v, u)
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	for i := 1; i < n; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		G[u] = append(G[u], v)
		G[v] = append(G[v], u)
	}
	fmt.Fscan(in, &s)
	s = " " + s
	dfs1(1, 0)
	dfs2(1, 0)
	fmt.Println(ans + 1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
