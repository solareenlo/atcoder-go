package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var Q int
	fmt.Fscan(in, &Q)
	pq := &Heap{}
	for Q > 0 {
		Q--
		var cmd int
		fmt.Fscan(in, &cmd)
		if cmd == 1 {
			var x int
			fmt.Fscan(in, &x)
			heap.Push(pq, x)
		} else if cmd == 2 {
			fmt.Println((*pq)[0])
		} else {
			heap.Pop(pq)
		}
	}
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
