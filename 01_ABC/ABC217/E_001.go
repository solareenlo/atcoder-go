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

	var n int
	fmt.Fscan(in, &n)

	pq := &Heap{}
	q := make([]int, 0)
	for i := 0; i < n; i++ {
		var op int
		fmt.Fscan(in, &op)
		switch op {
		case 1:
			var x int
			fmt.Fscan(in, &x)
			q = append(q, x)
		case 2:
			if pq.Len() > 0 {
				fmt.Fprintln(out, heap.Pop(pq).(int))
			} else {
				fmt.Fprintln(out, q[0])
				q = q[1:]
			}
		case 3:
			for len(q) > 0 {
				heap.Push(pq, q[0])
				q = q[1:]
			}
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
