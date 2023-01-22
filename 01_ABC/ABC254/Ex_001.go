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

	a := &Heap{}
	heap.Init(a)
	for i := 1; i <= n; i++ {
		var x int
		fmt.Fscan(in, &x)
		heap.Push(a, x)
	}
	b := &Heap{}
	heap.Init(b)
	for i := 1; i <= n; i++ {
		var x int
		fmt.Fscan(in, &x)
		heap.Push(b, x)
	}

	ans := 0
	for a.Len() != 0 && b.Len() != 0 {
		x := heap.Pop(a).(int)
		y := heap.Pop(b).(int)
		if x != y {
			ans++
		}
		if x > y {
			heap.Push(a, x/2)
			heap.Push(b, y)
		}
		if x < y && (y&1) != 0 {
			fmt.Println("-1")
			os.Exit(0)
		}
		if x < y {
			heap.Push(a, x)
			heap.Push(b, y/2)
		}
	}
	fmt.Println(ans)
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
