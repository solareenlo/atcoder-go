package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	type P struct {
		x, y int
	}

	var n, m int
	fmt.Fscan(in, &n, &m)
	e := make([]P, 0)
	deg1 := make([]int, n)
	deg2 := make([]int, n)
	t := make([]int, 0)
	var g [201][201]int
	for i := 0; i < m; i++ {
		var u, v int
		var c string
		fmt.Fscan(in, &u, &v, &c)
		u--
		v--
		g[u][v] = 1
		if c == "O" {
			e = append(e, P{u, v})
			deg1[u] ^= 1
			deg2[v] ^= 1
		}
	}
	for i := 0; i < n; i++ {
		g[i][i] = 1
		if (deg1[i] ^ deg2[i]) != 0 {
			t = append(t, i)
		}
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				g[i][j] |= g[i][k] & g[k][j]
			}
		}
	}

	var check func(int, int) int
	check = func(s, t int) int {
		for _, it := range e {
			if g[s][it.x] == 0 || g[it.y][t] == 0 {
				return 0
			}
		}
		return g[s][t]
	}

	ans := 0
	if len(t) == 2 {
		ans = check(t[0], t[1]) + check(t[1], t[0])
	}
	if len(t) == 0 {
		for i := 0; i < n; i++ {
			ans += check(i, i)
		}
	}
	fmt.Println(ans)
}
