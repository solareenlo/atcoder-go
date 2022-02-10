package main

import (
	"container/list"
	"fmt"
)

type pair struct {
	u, v int
}

func main() {
	var k int
	fmt.Scan(&k)

	visited := make([]bool, k)

	deque := list.New()
	deque.PushFront(pair{1, 1})

	for deque.Len() > 0 {
		p := deque.Remove(deque.Front())
		u, v := p.(pair).u, p.(pair).v
		if visited[u] {
			continue
		}
		visited[u] = true
		if u == 0 {
			fmt.Println(v)
			return
		}
		deque.PushFront(pair{(u * 10) % k, v})
		deque.PushBack(pair{(u + 1) % k, v + 1})
	}
}
