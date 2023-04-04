package main

import (
	"bufio"
	"fmt"
	"os"
)

var n, ans int
var g [][]int
var used []bool

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)
	g = make([][]int, n)
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}
	used = make([]bool, n)
	dfs(0, -1)
	if !used[0] {
		ans++
	}
	fmt.Println(ans - 1)
}

func dfs(u, p int) {
	ch := 0
	for _, v := range g[u] {
		if v != p {
			dfs(v, u)
			if !used[v] {
				ch++
			}
		}
	}
	if ch >= 2 {
		ans += ch - 1
		used[u] = true
	}
}
