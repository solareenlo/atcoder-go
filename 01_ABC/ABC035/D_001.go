package main

import (
	"container/heap"
	"fmt"
)

type edge struct {
	to, cost int
}

func dijkstra(s int, g [][]edge) []int {
	dist := make([]int, len(g))
	for i := range dist {
		dist[i] = int(1e9 + 7)
	}
	dist[s] = 0
	h := &Heap{}
	heap.Init(h)
	heap.Push(h, edge{0, s})
	for h.Len() > 0 {
		head := heap.Pop(h).(edge)
		if dist[head.cost] < head.to {
			continue
		}
		for _, e := range g[head.cost] {
			if dist[e.to] > head.to+e.cost {
				dist[e.to] = head.to + e.cost
				heap.Push(h, edge{to: dist[e.to], cost: e.to})
			}
		}
	}
	return dist
}

func main() {
	var n, m, t int
	fmt.Scan(&n, &m, &t)
	A := make([]int, n)
	for i := range A {
		fmt.Scan(&A[i])
	}

	g1 := make([][]edge, n)
	g2 := make([][]edge, n)
	var a, b, c int
	for i := 0; i < m; i++ {
		fmt.Scan(&a, &b, &c)
		a--
		b--
		g1[a] = append(g1[a], edge{to: b, cost: c})
		g2[b] = append(g2[b], edge{to: a, cost: c})
	}

	dist1 := dijkstra(0, g1)
	dist2 := dijkstra(0, g2)

	res := 0
	for i := 0; i < n; i++ {
		res = max(res, max(t-dist1[i]-dist2[i], 0)*A[i])
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Heap []edge

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i].cost < h[j].cost }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(edge)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
