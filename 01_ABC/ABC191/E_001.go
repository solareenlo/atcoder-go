package main

import (
	"container/heap"
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	g := make([][]pair, n)
	for i := 0; i < m; i++ {
		var a, b, c int
		fmt.Scan(&a, &b, &c)
		a--
		b--
		g[a] = append(g[a], pair{b, c})
	}

	for i := 0; i < n; i++ {
		vis := make([]int, n)
		pq := &Heap{}
		heap.Push(pq, pair{0, i})
		ok := false
		for pq.Len() > 0 {
			x := (*pq)[0]
			heap.Pop(pq)
			if x.cost == i && x.to != 0 {
				fmt.Println(-x.to)
				ok = true
				break
			}
			if vis[x.cost] != 0 {
				continue
			}
			vis[x.cost] = 1
			for _, y := range g[x.cost] {
				heap.Push(pq, pair{x.to - y.cost, y.to})
			}
		}
		if !ok {
			fmt.Println(-1)
		}
	}
}

type pair struct{ to, cost int }
type Heap []pair

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i].to > h[j].to }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(pair)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
