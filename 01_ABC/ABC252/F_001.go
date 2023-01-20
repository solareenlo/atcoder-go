package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, l int
	fmt.Fscan(in, &n, &l)

	q := &Heap{}
	heap.Init(q)
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		heap.Push(q, -a)
		l -= a
	}

	if l != 0 {
		heap.Push(q, -l)
	}

	res := 0
	for q.Len() >= 2 {
		a := -heap.Pop(q).(int)
		b := -heap.Pop(q).(int)
		res += a + b
		heap.Push(q, -a-b)
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
