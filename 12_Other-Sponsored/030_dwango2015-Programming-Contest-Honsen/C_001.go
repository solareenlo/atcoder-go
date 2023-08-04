package main

import (
	"bufio"
	"fmt"
	"os"
)

type edge struct {
	to, id int
}

const MOD = 1000000007

var pw [500009]int
var par, deg [400009]int
var g [400009][]edge
var ret int

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(in, &N, &M)
	pw[0] = 1
	for i := 0; i < N; i++ {
		par[i] = i
	}
	for i := 0; i < M; i++ {
		pw[i+1] = 2 * pw[i] % MOD
		ret = (ret + pw[i]) % MOD
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		if root(a) != root(b) {
			g[a] = append(g[a], edge{b, i})
			g[b] = append(g[b], edge{a, i})
			par[root(a)] = root(b)
		} else {
			deg[a] ^= 1
			deg[b] ^= 1
		}
	}
	dfs(0, -1)
	fmt.Println(2 * ret % MOD)
}

func root(x int) int {
	if x == par[x] {
		return x
	}
	par[x] = root(par[x])
	return par[x]
}

func dfs(pos, pre int) {
	for _, i := range g[pos] {
		if i.to == pre {
			continue
		}
		dfs(i.to, pos)
		if deg[i.to] == 1 {
			deg[i.to] ^= 1
			deg[pos] ^= 1
		} else {
			ret = (ret + pw[i.id]) % MOD
		}
	}
}
