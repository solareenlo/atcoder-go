package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 100100

var (
	head  = [N]int{}
	cnt   int
	ds    = [N]int{}
	depth = [N]int{}
	f     = [N][19]int{}
	edges = make([]edge, N*2)
)

type edge struct {
	t    int
	next int
}

func adde(f, t int) {
	cnt++
	edges[cnt] = edge{t, head[f]}
	head[f] = cnt
	cnt++
	edges[cnt] = edge{f, head[t]}
	head[t] = cnt
}

func dfs(u, fa int) {
	depth[u] = depth[fa] + 1
	f[u][0] = fa
	for i := 1; i < 19; i++ {
		f[u][i] = f[f[u][i-1]][i-1]
	}
	for i := head[u]; i > 0; i = edges[i].next {
		if edges[i].t != fa {
			dfs(edges[i].t, u)
		}
	}
}

func LCA(x, y int) int {
	if depth[x] < depth[y] {
		x ^= y
		y ^= x
		x ^= y
	}
	for i := 18; i >= 0; i-- {
		if depth[f[x][i]] >= depth[y] {
			x = f[x][i]
		}
	}
	if x == y {
		return x
	}
	for i := 18; i >= 0; i-- {
		if f[x][i] != f[y][i] {
			x = f[x][i]
			y = f[y][i]
		}
	}
	return f[x][0]
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, q int
	fmt.Fscan(in, &n, &q)

	for i := 1; i < n; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		adde(a, b)
	}
	for i := 1; i < n+1; i++ {
		ds[i] = 1e7
	}
	var s string
	fmt.Fscan(in, &s)
	s = "#" + s
	tp := make([]int, 0)
	for i := 1; i <= n; i++ {
		if s[i] == '1' {
			tp = append(tp, i)
			ds[i] = 0
		}
	}
	for len(tp) > 0 {
		x := tp[0]
		tp = tp[1:]
		for i := head[x]; i > 0; i = edges[i].next {
			if ds[edges[i].t] > ds[x]+1 {
				ds[edges[i].t] = ds[x] + 1
				tp = append(tp, edges[i].t)
			}
		}
	}

	dfs(1, 0)

	v1 := make([][19]int, n+1)
	v2 := make([][19]int, n+1)
	for i := 1; i <= n; i++ {
		v1[i][0] = ds[i] * 3
		v2[i][0] = ds[i] * 3
	}
	for j := 1; j <= 18; j++ {
		for i := 1; i <= n; i++ {
			v1[i][j] = min(v1[i][j-1], (1<<(j-1))+v1[f[i][j-1]][j-1])
		}
	}
	for j := 1; j <= 18; j++ {
		for i := 1; i <= n; i++ {
			v2[i][j] = min(v2[i][j-1]+(1<<(j-1)), v2[f[i][j-1]][j-1])
		}
	}

	for j := 0; j < q; j++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		l := LCA(a, b)
		le := depth[a] + depth[b] - 2*depth[l]
		ans := 2 * le
		l1 := 0
		ans = min(ans, le+depth[a]-depth[l]+3*ds[l])
		for i := 18; i >= 0; i-- {
			if depth[f[a][i]] >= depth[l] {
				ans = min(ans, le+l1+v1[a][i])
				a = f[a][i]
				l1 += 1 << i
			}
		}
		l1 = 0
		for i := 18; i >= 0; i-- {
			if depth[f[b][i]] >= depth[l] {
				l1 += 1 << i
				ans = min(ans, le*2-l1+v2[b][i]+1)
				b = f[b][i]
			}
		}
		fmt.Fprintln(out, ans)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
