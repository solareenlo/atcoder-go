package main

import (
	"container/list"
	"fmt"
)

type pair struct{ u, dist int }

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	g := make([][]int, n+1)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Scan(&u, &v)
		g[u] = append(g[u], v)
	}

	var s, t int
	fmt.Scan(&s, &t)

	d := make([][3]bool, n+1)
	q := list.New()
	q.PushBack(pair{s, 0})
	for q.Len() > 0 {
		front := q.Front()
		q.Remove(front)
		u := front.Value.(pair).u
		dist := front.Value.(pair).dist
		if d[u][dist%3] {
			continue
		}
		d[u][dist%3] = true
		if u == t && dist%3 == 0 {
			fmt.Println(dist / 3)
			return
		}
		for _, v := range g[u] {
			if !d[v][(dist+1)%3] {
				q.PushBack(pair{v, dist + 1})
			}
		}
	}
	fmt.Println(-1)
}
