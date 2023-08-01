package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	g := make([][]int, n)
	used := make([]int, n)
	for i := range used {
		used[i] = -1
	}
	for i := 0; i < n; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}
	var dfs func(int, int, int) int
	dfs = func(v, p, d int) int {
		if used[v] != -1 {
			return d - used[v]
		}
		used[v] = d
		for _, u := range g[v] {
			if u == p {
				continue
			}
			res := dfs(u, v, d+1)
			if res != 0 {
				return res
			}
		}
		return 0
	}
	fmt.Println(dfs(0, -1, 0))
}
