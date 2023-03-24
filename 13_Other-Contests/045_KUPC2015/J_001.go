package main

import (
	"bufio"
	"fmt"
	"os"
)

const maxN = 100005

var (
	N, Q    int
	A, B, C [maxN]int
	X, Y, Z [maxN]int
	G       [maxN][]int
	dist    [maxN]int
	par, dp [maxN][22]int
)

func dfs(pos int, dep int) {
	dist[pos] = dep
	for _, i := range G[pos] {
		if dist[i] != -1 {
			continue
		}
		dfs(i, dep+1)
		par[i][0] = pos
	}
}

func prevs(pos int, x int) int {
	for i := 19; i >= 0; i-- {
		if x >= (1 << i) {
			pos = par[pos][i]
			x -= (1 << i)
		}
	}
	return pos
}

func lca(u int, v int) int {
	if dist[u] > dist[v] {
		u, v = v, u
	}
	v = prevs(v, dist[v]-dist[u])
	if u == v {
		return u
	}

	for i := 19; i >= 0; i-- {
		if par[u][i] == par[v][i] {
			continue
		}
		u = par[u][i]
		v = par[v][i]
	}
	return par[u][0]
}

func solve(u int, v int, x int) int {
	w := lca(u, v)
	cx := u
	for {
		for i := 19; i >= 0; i-- {
			if dp[cx][i] > x {
				cx = par[cx][i]
			}
		}
		if dist[cx] < dist[w] {
			break
		}
		x = (x % A[cx])
		cx = par[cx][0]
	}

	dx := w
	for {
		ex := v
		tx := -1
		lim := -1
		for i := 19; i >= 0; i-- {
			if dist[ex]-dist[dx] < (1 << i) {
				continue
			}
			if dp[ex][i] <= x {
				tx = ex
				lim = i
			}
			ex = par[ex][i]
		}
		if lim == -1 {
			break
		}
		for i := lim - 1; i >= 0; i-- {
			nex := par[tx][i]
			if dp[nex][i] <= x {
				tx = par[tx][i]
			}
		}
		x = (x % A[tx])
	}
	return x
}
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &N, &Q)
	for i := 1; i <= N; i++ {
		fmt.Fscan(in, &A[i])
	}
	for i := 1; i <= N-1; i++ {
		fmt.Fscan(in, &B[i], &C[i])
	}
	for i := 1; i <= Q; i++ {
		fmt.Fscan(in, &X[i], &Y[i], &Z[i])
	}
	for i := 1; i <= N-1; i++ {
		G[B[i]] = append(G[B[i]], C[i])
		G[C[i]] = append(G[C[i]], B[i])
	}

	for i := 0; i <= N; i++ {
		dist[i] = -1
	}
	dfs(1, 0)
	for i := 1; i <= 19; i++ {
		for j := 1; j <= N; j++ {
			par[j][i] = par[par[j][i-1]][i-1]
		}
	}
	for i := 1; i <= N; i++ {
		dp[i][0] = A[i]
	}
	for i := 1; i <= 19; i++ {
		for j := 1; j <= N; j++ {
			dp[j][i] = min(dp[j][i-1], dp[par[j][i-1]][i-1])
		}
	}

	for i := 1; i <= Q; i++ {
		Answer := solve(Y[i], Z[i], X[i])
		fmt.Fprintln(out, Answer)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
