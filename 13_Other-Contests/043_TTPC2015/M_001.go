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

	c := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &c[i])
	}

	E := make([][]int, n)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		E[u] = append(E[u], v)
		E[v] = append(E[v], u)
	}

	dis := make([]int, n)
	for i := 1; i < n; i++ {
		dis[i] = 1000000001
	}
	dis[0] = 0

	Q := make([]int, 0)
	Q = append(Q, 0)

	ans := 0
	for len(Q) != 0 {
		u := Q[0]
		Q = Q[1:]

		for i := 0; i < len(E[u]); i++ {
			v := E[u][i]
			if dis[v] != 1000000001 {
				continue
			}

			dis[v] = dis[u] + 1

			if dis[v]&1 != 0 {
				ans ^= c[v]
			}

			Q = append(Q, v)
		}
	}

	if ans == 0 {
		fmt.Println("Second")
	} else {
		fmt.Println("First")
	}
}
