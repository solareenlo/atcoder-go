package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

type node struct {
	u, dist int
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	g := make([][]node, n)
	var l, r, d int
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &l, &r, &d)
		l--
		r--
		g[l] = append(g[l], node{r, d})
		g[r] = append(g[r], node{l, -d})
	}

	mini := -(1 << 62)
	dist := make([]int, n)
	for i := range dist {
		dist[i] = mini
	}
	for i := 0; i < n; i++ {
		if dist[i] != mini {
			continue
		}
		dist[i] = 0
		deque := list.New()
		deque.PushFront(i)
		for deque.Len() > 0 {
			t := deque.Remove(deque.Front())
			for _, to := range g[t.(int)] {
				if dist[to.u] == mini {
					dist[to.u] = dist[t.(int)] + to.dist
					deque.PushBack(to.u)
				} else {
					if dist[to.u] != dist[t.(int)]+to.dist {
						fmt.Println("No")
						return
					}
				}
			}
		}
	}
	fmt.Println("Yes")
}
