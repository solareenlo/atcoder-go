package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var q int
	fmt.Scan(&q)

	ih := &intHeap{}
	heap.Init(ih)

	sum := 0
	for i := 0; i < q; i++ {
		var p, x int
		fmt.Fscan(in, &p)
		if p == 1 {
			fmt.Fscan(in, &x)
			heap.Push(ih, x-sum)
		} else if p == 2 {
			fmt.Fscan(in, &x)
			sum += x
		} else if p == 3 {
			fmt.Println(heap.Pop(ih).(int) + sum)
		}
	}
}

type intHeap []int

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
