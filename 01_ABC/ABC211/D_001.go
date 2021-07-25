package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Scan(&n, &m)

	path := make([][]int, n)
	var a, b int
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &a, &b)
		a--
		b--
		path[a] = append(path[a], b)
		path[b] = append(path[b], a)
	}

	cost := make([]int, n)
	for i := 0; i < n; i++ {
		cost[i] = 200001
	}
	cost[0] = 0

	to := make([]int, n)
	to[0] = 1

	q := make([]int, 0)
	q = append(q, 0)

	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		for _, u := range path[v] {
			if cost[u] > cost[v]+1 {
				cost[u] = cost[v] + 1
				to[u] = to[v]
				q = append(q, u)
			} else if cost[u] == cost[v]+1 {
				to[u] = to[u] + to[v]
				to[u] %= 1000000007
			}
		}
	}
	fmt.Println(to[n-1])
}
