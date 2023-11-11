package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	const N = 10005
	const S = 5005
	const INF = int(1e18)

	type edge struct {
		to, nxt int
	}

	var e [N]edge
	var tot int
	var h, fa, siz [N]int
	var n, A, B, C int
	var f [N][S][2]int

	var add func(int, int)
	add = func(u, v int) {
		tot++
		e[tot] = edge{v, h[u]}
		h[u] = tot
	}

	var dfs func(int)
	dfs = func(u int) {
		for i := 0; i < C+1; i++ {
			f[u][i][0] = -INF
			f[u][i][1] = -INF
		}
		f[u][0][0] = 0
		if h[u] > 0 {
			siz[u] = 1
			f[u][1][1] = 1
		} else {
			siz[u] = 0
			f[u][0][1] = 1
		}
		for i := h[u]; i > 0; i = e[i].nxt {
			v := e[i].to
			dfs(v)
			for j := min(siz[u], C); j >= 0; j-- {
				for k := min(siz[v], C-j); k >= 0; k-- {
					f[u][j+k][0] = max(f[u][j+k][0], f[u][j][0]+max(f[v][k][0], f[v][k][1]))
					f[u][j+k][1] = max(f[u][j+k][1], f[u][j][1]+f[v][k][0])
				}
			}
			siz[u] += siz[v]
		}
	}

	var t int
	fmt.Fscan(in, &t)
	for t > 0 {
		t--
		fmt.Fscan(in, &n, &A, &B, &C)
		tot = 0
		for i := 1; i <= n; i++ {
			h[i] = 0
		}
		for i := 2; i <= n; i++ {
			fmt.Fscan(in, &fa[i])
			add(fa[i], i)
		}
		if (C & 1) != 0 {
			fmt.Fprintln(out, "No")
			continue
		}
		C /= 2
		dfs(1)
		if (B >= C && f[1][C][0] >= B) || (B+1 >= C && f[1][C][1] >= B+1) {
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
		}
	}
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
