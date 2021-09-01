package main

import "fmt"

func main() {
	var n, a, b, m int
	fmt.Scan(&n, &a, &b, &m)

	g := make([][]int, n+1)
	var x, y int
	for i := 0; i < m; i++ {
		fmt.Scan(&x, &y)
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	const mod int = int(1e9 + 7)

	dist := make([]int, n+1)
	cnt := make([]int, n+1)
	dist[a] = 1
	cnt[a] = 1

	q := make([]int, 0)
	q = append(q, a)
	for len(q) != 0 {
		v := q[0]
		q = q[1:]
		for _, e := range g[v] {
			if dist[e] == 0 {
				dist[e] = dist[v] + 1
				cnt[e] = cnt[v]
				q = append(q, e)
			} else if dist[e] == dist[v]+1 {
				cnt[e] = (cnt[e] + cnt[v]) % mod
			}
		}
	}
	fmt.Println(cnt[b])
}
