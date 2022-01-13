package main

import (
	"bufio"
	"fmt"
	"os"
)

type P struct{ x, y int }

var (
	n int
	g = make([][]int, 100005)
)

func dfs(u, p int) P {
	res := P{0, u}
	for _, v := range g[u] {
		if v != p {
			p := dfs(v, u)
			p.x++
			if res.x < p.x {
				res = p
			}
		}
	}
	return res
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}
	u := dfs(1, 0)
	v := dfs(u.y, 0)
	fmt.Println(u.y, v.y)
}
