package main

import (
	"bufio"
	"fmt"
	"os"
)

const MAXN = 2e5 + 5

var (
	pre = [MAXN]int{}
	m   int
)

func fnd(x int) int {
	if x != pre[x] {
		pre[x] = fnd(pre[x])
	}
	return pre[x]
}

var (
	cnte int
	h    = [MAXN]int{}
	to   = [MAXN << 2]int{}
	nx   = [MAXN << 2]int{}
)

func adde(u, v int) {
	cnte++
	nx[cnte] = h[u]
	to[cnte] = v
	h[u] = cnte
}

var (
	fa  = [MAXN][19]int{}
	dep = [MAXN]int{}
	sum = [MAXN]int{}
)

func Dfs(u int) {
	dep[u] = dep[fa[u][0]] + 1
	for i := 1; ; i++ {
		fa[u][i] = fa[fa[u][i-1]][i-1]
		if fa[u][i] <= 0 {
			break
		}
	}
	for i := h[u]; i > 0; i = nx[i] {
		v := to[i]
		if v == fa[u][0] {
			continue
		}
		fa[v][0] = u
		Dfs(v)
	}
}

func Lca(x, y int) int {
	if dep[x] > dep[y] {
		x, y = y, x
	}
	for i := 17; i >= 0; i-- {
		if dep[fa[y][i]] >= dep[x] {
			y = fa[y][i]
		}
	}
	if x == y {
		return x
	}
	for i := 17; i >= 0; i-- {
		if fa[x][i] != fa[y][i] {
			x = fa[x][i]
			y = fa[y][i]
		}
	}
	return fa[x][0]
}

type pair struct{ x, y int }

var (
	a   = make([]pair, MAXN*2)
	vis = [MAXN << 2]bool{}
	ans int
)

func Dfs2(u int) {
	for i := h[u]; i > 0; i = nx[i] {
		v := to[i]
		if v == fa[u][0] {
			continue
		}
		Dfs2(v)
		sum[u] += sum[v]
	}
	ans += min(sum[u], 2)
	if sum[u]&1 != 0 {
		sum[u]++
		sum[fa[u][0]]--
		m++
		a[m] = pair{u, fa[u][0]}
	}
}

func Dfs3(u int) {
	for i := h[u]; i > 0; {
		if !vis[i] {
			v := to[i]
			if i&1 != 0 {
				a[i>>1].x, a[i>>1].y = a[i>>1].y, a[i>>1].x
			}
			vis[i^1] = true
			vis[i] = true
			Dfs3(v)
		}
		i = nx[i]
		h[u] = i
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n, &m)
	for i := 1; i < n; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		adde(u, v)
		adde(v, u)
	}
	Dfs(1)
	for i := 1; i <= m; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		sum[u]++
		sum[v]++
		sum[Lca(u, v)] -= 2
		a[i] = pair{u, v}
	}
	_m := m
	Dfs2(1)
	cnte = 1
	for i := range h {
		h[i] = 0
	}
	for i := 1; i <= n; i++ {
		pre[i] = i
	}
	for i := 1; i <= m; i++ {
		u := a[i].x
		v := a[i].y
		adde(u, v)
		adde(v, u)
		pre[fnd(u)] = fnd(v)
	}
	for i := 1; i <= n; i++ {
		if pre[i] == i {
			Dfs3(i)
		}
	}
	fmt.Fprintln(out, ans)
	for i := 1; i <= _m; i++ {
		fmt.Fprintln(out, a[i].x, a[i].y)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
