package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	e    = make([][]int, 200002)
	used = make([]bool, 200002)
	x, y int
)

func dfs(k int) {
	used[k] = true
	x++
	y += len(e[k])
	for i := 0; i < len(e[k]); i++ {
		if !used[e[k][i]] {
			dfs(e[k][i])
		}
	}
	return
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		e[u] = append(e[u], v)
		e[v] = append(e[v], u)
	}

	res := 1
	for i := 0; i < n; i++ {
		if !used[i] {
			x = 0
			y = 0
			dfs(i)
			if y == (x * 2) {
				res = (res * 2) % 998244353
			} else {
				res = 0
			}
		}
	}
	fmt.Println(res)
}
