package main

import (
	"bufio"
	"fmt"
	"os"
)

var h [100010]int
var nex, to [200010]int
var M int
var d, v, f, g [100010]int

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	for i := 1; i < n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		add(x, y)
		add(y, x)
		d[x]++
		d[y]++
	}
	s := 0
	for i := 1; i <= n; i++ {
		if d[i] == 1 {
			s++
		}
		if d[i] == 2 {
			v[i] = 1
		} else {
			v[i] = 0
		}
	}
	dfs(0, 1)
	fmt.Println(s + f[1])
}

func add(a, b int) {
	M++
	to[M] = b
	nex[M] = h[a]
	h[a] = M
}

func dfs(fa, x int) {
	f[x] = v[x]
	g[x] = v[x]
	for i := h[x]; i > 0; i = nex[i] {
		if to[i] != fa {
			dfs(x, to[i])
			f[x] = max(f[x], max(g[x]+g[to[i]], f[to[i]]))
			g[x] = max(g[x], g[to[i]]+v[x])
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
