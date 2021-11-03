package main

import (
	"container/heap"
	"fmt"
)

func main() {
	var k int
	fmt.Scan(&k)

	q := &Heap{}
	for i := 0; i < 9; i++ {
		heap.Push(q, i+1)
	}

	for i := 0; i < k-1; i++ {
		num := (*q)[0]
		heap.Pop(q)
		rem := num % 10
		if rem != 0 {
			heap.Push(q, 10*num+rem-1)
		}
		heap.Push(q, 10*num+rem)
		if rem != 9 {
			heap.Push(q, 10*num+rem+1)
		}
	}

	fmt.Println((*q)[0])
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
