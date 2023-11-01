package main

import (
	"bufio"
	"fmt"
	"os"
)

const maxn = 1000010

var vis [maxn]bool
var col [maxn]int
var cnt [2]int
var G [maxn][]int
var flag bool

func dfs(u, c int) {
	vis[u] = true
	col[u] = c
	cnt[c]++
	for i := 0; i < len(G[u]); i++ {
		v := G[u][i]
		if !vis[v] {
			dfs(v, c^1)
		} else {
			if col[u] == col[v] {
				flag = true
			}
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 1; i <= m; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		G[u] = append(G[u], v)
		G[v] = append(G[v], u)
	}
	flag = false
	ans := n*(n-1)/2 - m
	for i := 1; i <= n; i++ {
		if !vis[i] {
			cnt[0] = 0
			cnt[1] = 0
			dfs(i, 0)
			ans -= cnt[0] * (cnt[0] - 1) / 2
			ans -= cnt[1] * (cnt[1] - 1) / 2
		}
	}
	if flag {
		fmt.Println(0)
	} else {
		fmt.Println(ans)
	}
}
