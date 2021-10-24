package main

import (
	"container/heap"
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	a := &Heap{}
	for i := 0; i < n; i++ {
		var A int
		fmt.Scan(&A)
		heap.Push(a, A)
	}

	for i := 0; i < m; i++ {
		top := (*a)[0]
		heap.Pop(a)
		heap.Push(a, top/2)
	}

	sum := 0
	for a.Len() > 0 {
		sum += (*a)[0]
		heap.Pop(a)
	}
	fmt.Println(sum)
}

type Heap []int

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i] > h[j] }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(int)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
