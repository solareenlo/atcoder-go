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
	fmt.Fscan(in, &q)

	l := &HeapL{}
	r := &HeapR{}
	var t, a, b int
	res := 0
	for i := 0; i < q; i++ {
		fmt.Fscan(in, &t)
		if t == 1 {
			fmt.Fscan(in, &a, &b)
			res += b
			heap.Push(l, a)
			heap.Push(r, a)
			x := heap.Pop(l).(int)
			y := heap.Pop(r).(int)
			if x > y {
				res += x - y
				x, y = y, x
			}
			heap.Push(l, x)
			heap.Push(r, y)
		} else {
			fmt.Fprintln(out, (*l)[0], res)
		}
	}
}

type HeapL []int

func (h HeapL) Len() int            { return len(h) }
func (h HeapL) Less(i, j int) bool  { return h[i] > h[j] }
func (h HeapL) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapL) Push(x interface{}) { *h = append(*h, x.(int)) }

func (h *HeapL) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type HeapR []int

func (h HeapR) Len() int            { return len(h) }
func (h HeapR) Less(i, j int) bool  { return h[i] < h[j] }
func (h HeapR) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapR) Push(x interface{}) { *h = append(*h, x.(int)) }

func (h *HeapR) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
