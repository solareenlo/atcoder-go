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

	g := make([][]int, n+1)
	for i := 1; i < n; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	p := make([]int, n+1)
	var dfs func(u, v int) [3]int
	dfs = func(u, v int) [3]int {
		p[v] = u
		s, d, c := v, 0, 1
		for _, w := range g[v] {
			if w != u {
				res := dfs(v, w)
				S, D, C := res[0], res[1], res[2]
				if d <= D {
					d = D + 1
					s = S
					c = 0
				}
				if d == D+1 {
					c += C
				}
			}
		}
		return [3]int{s, d, c}
	}

	res := dfs(0, dfs(0, 1)[0])
	s, d := res[0], res[1]
	for i := 0; i < d/2; i++ {
		s = p[s]
	}

	const mod = 998244353
	x, y := 1, 1
	if d&1 != 0 {
		x = dfs(s, p[s])[2]
		x *= dfs(p[s], s)[2]
		x %= mod
		x++
	} else {
		for _, v := range g[s] {
			res = dfs(s, v)
			D, C := res[1], res[2]
			if D == d/2-1 {
				x *= C + 1
				x %= mod
				y += C
				y %= mod
			}
		}
	}

	fmt.Println((x - y + mod) % mod)
}
