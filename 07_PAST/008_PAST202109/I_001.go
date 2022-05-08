package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	pq := &Heap{}
	cnt := 0
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		for a%2 == 0 {
			cnt++
			a /= 2
		}
		heap.Push(pq, -a)
	}

	for i := 0; i < cnt; i++ {
		heap.Push(pq, 3*(*pq)[0])
		heap.Pop(pq)
	}
	fmt.Println(-(*pq)[0])
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
