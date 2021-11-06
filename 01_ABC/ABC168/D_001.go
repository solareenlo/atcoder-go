package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	g := make([][]int, n+1)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	dist := make([]int, n+1)
	dist[1] = 1

	q := make([]int, 0)
	q = append(q, 1)
	v := 0
	for len(q) > 0 {
		v, q = q[0], q[1:]
		for _, i := range g[v] {
			if dist[i] == 0 {
				dist[i] = v
				q = append(q, i)
			}
		}
	}

	fmt.Println("Yes")
	for i := 2; i < n+1; i++ {
		fmt.Println(dist[i])
	}
}
