package main

import (
	"bufio"
	"fmt"
	"os"
)

type P struct {
	x, y int
}

var G [1 << 18][]P
var par, to [18][1 << 18]int
var dep, dist [1 << 18]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, k int
	fmt.Fscan(in, &n, &k)
	for i := 0; i < n-1; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		a--
		b--
		G[a] = append(G[a], P{b, c})
		G[b] = append(G[b], P{a, c})
	}
	dfs(0)
	for i := 0; i < 17; i++ {
		for u := 0; u < n; u++ {
			par[i+1][u] = par[i][par[i][u]]
		}
	}
	for u := 0; u < n; u++ {
		if dist[u] < k {
			to[0][u] = -1
		} else {
			v := u
			for i := 17; i >= 0; i-- {
				if dist[u]-dist[par[i][v]] < k {
					v = par[i][v]
				}
			}
			to[0][u] = par[0][v]
		}
	}
	for i := 0; i < 17; i++ {
		for u := 0; u < n; u++ {
			if to[i][u] == -1 {
				to[i+1][u] = -1
			} else {
				to[i+1][u] = to[i][to[i][u]]
			}
		}
	}

	var q int
	fmt.Fscan(in, &q)
	for q > 0 {
		q--
		var s, t int
		fmt.Fscan(in, &s, &t)
		s--
		t--
		l := lca(s, t)
		ans := 0
		for i := 17; i >= 0; i-- {
			if to[i][s] != -1 && dep[to[i][s]] >= dep[l] {
				ans += 1 << i
				s = to[i][s]
			}
			if to[i][t] != -1 && dep[to[i][t]] >= dep[l] {
				ans += 1 << i
				t = to[i][t]
			}
		}
		if dist[s]+dist[t]-2*dist[l] >= k {
			ans++
		}
		fmt.Fprintln(out, ans)
	}
}

func dfs(u int) {
	for _, tmp := range G[u] {
		v := tmp.x
		d := tmp.y
		if v != par[0][u] {
			par[0][v] = u
			dep[v] = dep[u] + 1
			dist[v] = dist[u] + d
			dfs(v)
		}
	}
}

func lca(u, v int) int {
	if dep[u] > dep[v] {
		u, v = v, u
	}
	d := dep[v] - dep[u]
	for i := 0; i < 18; i++ {
		if ((d >> i) & 1) != 0 {
			v = par[i][v]
		}
	}
	if u == v {
		return u
	}
	for i := 17; i >= 0; i-- {
		if par[i][u] != par[i][v] {
			u = par[i][u]
			v = par[i][v]
		}
	}
	return par[0][u]
}
