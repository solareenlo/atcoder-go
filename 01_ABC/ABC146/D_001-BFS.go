package main

import (
	"container/heap"
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	type Edge struct{ color, i int }
	g := make([][]Edge, n)
	for i := 0; i < n-1; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		a--
		b--
		g[a] = append(g[a], Edge{b, i})
		g[b] = append(g[b], Edge{a, i})
	}

	maxi := 0
	color := make([]int, n)
	color[0] = -1
	res := make([]int, n-1)
	q := &Heap{}
	heap.Push(q, 0)
	for q.Len() > 0 {
		v := (*q)[0]
		heap.Pop(q)
		cnt := 0
		for _, e := range g[v] {
			if color[e.color] != 0 {
				continue
			}
			heap.Push(q, e.color)
			cnt++
			if cnt == color[v] {
				cnt++
			}
			res[e.i] = cnt
			color[e.color] = cnt
			maxi = max(maxi, cnt)
		}
	}

	fmt.Println(maxi)
	for i := 0; i < n-1; i++ {
		fmt.Println(res[i])
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Heap []int

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i] < h[j] }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(int)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
