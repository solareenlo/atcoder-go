package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N, M, Q, K int
	fmt.Fscan(in, &N, &M, &Q, &K)

	x := make([]int, K)
	for i := range x {
		fmt.Fscan(in, &x[i])
	}

	adj := make([][]int, 200005)
	for i := 0; i < M; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	d := [20][200005]int{}
	q := make([]int, 0)
	for i := 0; i < K; i++ {
		q = append(q, x[i])
		for len(q) > 0 {
			a := q[0]
			q = q[1:]
			for _, b := range adj[a] {
				if b == x[i] || d[i][b] != 0 {
					continue
				}
				d[i][b] = d[i][a] + 1
				q = append(q, b)
			}
		}
	}

	for i := 0; i < Q; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		ans := 1 << 60
		for i := 0; i < K; i++ {
			ans = min(ans, d[i][u]+d[i][v])
		}
		fmt.Fprintln(out, ans)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
