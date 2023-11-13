package main

import (
	"bufio"
	"fmt"
	"os"
)

type pair struct {
	x, y int
}

var G [2 << 17][]pair
var n, ans int

func dfs(u, p, l, r int) {
	if u != 0 {
		ans += l * (n - r)
	}
	for _, tmp := range G[u] {
		v, i := tmp.x, tmp.y
		if v != p {
			dfs(v, u, min(l, i), max(r, i))
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)
	for i := 1; i < n; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		G[u] = append(G[u], pair{v, i})
		G[v] = append(G[v], pair{u, i})
	}
	ans = n * (n - 1) / 2
	dfs(0, -1, n, 1)
	fmt.Println(ans)
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
