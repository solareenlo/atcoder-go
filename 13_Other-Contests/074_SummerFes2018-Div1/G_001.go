package main

import (
	"bufio"
	"fmt"
	"os"
)

const M = 1000000007

var adj [][]int
var memo, fact []int

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	adj = make([][]int, n)
	memo = make([]int, n)
	for i := range memo {
		memo[i] = -1
	}
	fact = make([]int, n)
	fact[0] = 1
	for i := 0; i < n-1; i++ {
		fact[i+1] = fact[i] * (i + 1) % M
		var u, v int
		fmt.Fscan(in, &u, &v)
		adj[u-1] = append(adj[u-1], v-1)
		adj[v-1] = append(adj[v-1], u-1)
	}
	if n == 2 {
		fmt.Println(4)
		return
	}
	fmt.Println(dfs2(0))
}

func Dfs(p, par int) int {
	if memo[p] != -1 {
		return memo[p]
	}
	next := -1
	for _, c := range adj[p] {
		if c == par {
			continue
		}
		if len(adj[c]) > 1 {
			if next != -1 {
				memo[p] = 0
				return memo[p]
			}
			next = c
		}
	}
	cc := len(adj[p]) - 1
	if next == -1 {
		memo[p] = fact[cc] * 2 % M
		return memo[p]
	} else {
		memo[p] = fact[cc-1] * Dfs(next, p) % M
		return memo[p]
	}
}

func dfs2(p int) int {
	nexts := make([]int, 0)
	for _, c := range adj[p] {
		if len(adj[c]) > 1 {
			nexts = append(nexts, c)
		}
	}
	cc := len(adj[p])
	if cc == 1 {
		return dfs2(adj[p][0])
	}
	switch len(nexts) {
	case 0:
		return fact[cc] * 4 % M
	case 1:
		return fact[cc-1] * Dfs(nexts[0], p) * 4 % M
	case 2:
		return fact[cc-2] * Dfs(nexts[0], p) % M * Dfs(nexts[1], p) * 2 % M
	default:
		return 0
	}
}
