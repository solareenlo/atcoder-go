package main

import (
	"container/heap"
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	job := make([][]int, 100001)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		job[a-1] = append(job[a-1], b)
	}

	res := 0
	q := &Heap{}
	for i := 0; i < m; i++ {
		for _, b := range job[i] {
			heap.Push(q, b)
		}
		if q.Len() > 0 {
			res += (*q)[0]
			heap.Pop(q)
		}
	}
	fmt.Println(res)
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
