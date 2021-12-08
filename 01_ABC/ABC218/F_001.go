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

	var n, m int
	fmt.Fscan(in, &n, &m)

	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = -1
		}
	}

	type pair struct{ a, b int }
	edge := make([]pair, m)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		g[a][b] = i
		edge[i] = pair{a, b}
	}

	dist := make([]int, n)
	for i := range dist {
		dist[i] = -1
	}
	dist[0] = 0

	memo := make([]pair, n)
	que := make([]int, 0)
	que = append(que, 0)
	for len(que) > 0 {
		i := que[0]
		que = que[1:]
		for j := 0; j < n; j++ {
			if dist[j] == -1 && g[i][j] != -1 {
				dist[j] = dist[i] + 1
				memo[j] = pair{i, g[i][j]}
				que = append(que, j)
			}
		}
	}

	if dist[n-1] == -1 {
		for i := 0; i < m; i++ {
			fmt.Fprintln(out, -1)
		}
		return
	}

	shortest_path := make([]int, 0)
	now := n - 1
	for now != 0 {
		shortest_path = append(shortest_path, memo[now].b)
		now = memo[now].a
	}

	res := make([]int, m)
	for i := range res {
		res[i] = dist[n-1]
	}

	for _, e := range shortest_path {
		g[edge[e].a][edge[e].b] = -1
		q := make([]int, 0)
		q = append(q, 0)
		dis := make([]int, n)
		for i := range dis {
			dis[i] = -1
		}
		dis[0] = 0
		for len(q) > 0 {
			i := q[0]
			q = q[1:]
			for j := 0; j < n; j++ {
				if dis[j] == -1 && g[i][j] != -1 {
					dis[j] = dis[i] + 1
					q = append(q, j)
				}
			}
		}
		res[e] = dis[n-1]
		g[edge[e].a][edge[e].b] = e
	}

	for i := range res {
		fmt.Fprintln(out, res[i])
	}
}
