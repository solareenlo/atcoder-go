package main

import (
	"bufio"
	"fmt"
	"os"
)

const MXN = 1 << 17

var sz, par [MXN]int
var G [MXN][]int
var dead [MXN]bool
var V []int

func dfs_sz(u, p int) {
	sz[u] = 1
	for _, v := range G[u] {
		if v != p && !dead[v] {
			dfs_sz(v, u)
			sz[u] += sz[v]
		}
	}
	V = append(V, u)
}

func dfs_cd(r, p int) {
	V = make([]int, 0)
	dfs_sz(r, -1)
	for _, v := range V {
		if 2*sz[v] > sz[r] {
			r = v
			if r != 0 {
				break
			}
		}
	}
	par[r] = p
	dead[r] = true
	for _, v := range G[r] {
		if !dead[v] {
			dfs_cd(v, r)
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		G[u] = append(G[u], v)
		G[v] = append(G[v], u)
	}
	dfs_cd(1, -1)
	for u := 1; u <= n; u++ {
		fmt.Printf("%d ", par[u])
	}
}
