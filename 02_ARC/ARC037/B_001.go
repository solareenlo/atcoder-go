package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	ok   bool
	seen = make([]bool, 0)
	G    = make([][]int, 0)
)

func dfs(v, prev_v int) {
	seen[v] = true
	for _, next_v := range G[v] {
		if seen[next_v] {
			if next_v != prev_v {
				ok = false
			}
			continue
		}
		dfs(next_v, v)
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	G = make([][]int, n)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		G[u] = append(G[u], v)
		G[v] = append(G[v], u)
	}

	cnt := 0
	seen = make([]bool, n)
	for i := 0; i < n; i++ {
		ok = true
		if seen[i] {
			continue
		}
		dfs(i, -1)
		if ok {
			cnt++
		}
	}

	fmt.Println(cnt)
}
