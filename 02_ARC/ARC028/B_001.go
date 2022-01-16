package main

import (
	"container/heap"
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	q := &Heap{}
	heap.Init(q)
	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		heap.Push(q, pair{x, i + 1})
		if q.Len() > k {
			heap.Pop(q)
		}
		if k <= i+1 {
			fmt.Println((*q)[0].y)
		}
	}
}

type pair struct{ x, y int }
type Heap []pair

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i].x > h[j].x }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(pair)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
