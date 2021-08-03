package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var q int
	fmt.Scan(&q)

	ih := &intHeap{}
	heap.Init(ih)

	offset := 0
	for i := 0; i < q; i++ {
		var p int
		fmt.Fscan(in, &p)
		if p == 3 {
			fmt.Fprintln(out, heap.Pop(ih).(int)+offset)
			continue
		}
		var x int
		fmt.Fscan(in, &x)
		if p == 1 {
			heap.Push(ih, x-offset)
			continue
		}
		offset += x
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
