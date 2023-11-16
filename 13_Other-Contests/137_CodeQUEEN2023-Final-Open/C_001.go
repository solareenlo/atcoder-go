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

	var n, s, t int
	fmt.Fscan(in, &n, &s, &t)
	s--
	t--
	G := make([][]int, n)
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		G[u] = append(G[u], v)
		G[v] = append(G[v], u)
	}
	var bfs func(int) []int
	bfs = func(start int) []int {
		dist := make([]int, n)
		for i := range dist {
			dist[i] = -1
		}
		dist[start] = 0
		Q := make([]int, 0)
		Q = append(Q, start)
		for len(Q) != 0 {
			p := Q[0]
			Q = Q[1:]
			for _, u := range G[p] {
				if dist[u] == -1 {
					dist[u] = dist[p] + 1
					Q = append(Q, u)
				}
			}
		}
		return dist
	}
	d1 := bfs(s)
	d2 := bfs(t)
	for i := 0; i < n; i++ {
		ans := (d1[i]+d2[i]-d1[t])/2 + 1
		fmt.Fprintln(out, ans)
	}
}
