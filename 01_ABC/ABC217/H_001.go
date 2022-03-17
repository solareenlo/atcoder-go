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

	L := &Heap{}
	R := &Heap{}
	for i := 1; i <= n; i++ {
		heap.Push(L, 0)
		heap.Push(R, 0)
	}

	ans := 0
	for i := 1; i <= n; i++ {
		var t, d, x int
		fmt.Fscan(in, &t, &d, &x)
		l := -t
		r := t
		if d == 0 {
			ans += max(0, x-(-(*R)[0])-r)
			heap.Push(R, -(x - r))
			heap.Push(L, -(*R)[0]+r-l)
			heap.Pop(R)
		}
		if d == 1 {
			ans += max(0, -x+(*L)[0]+l)
			heap.Push(L, x-l)
			heap.Push(R, -((*L)[0] + l - r))
			heap.Pop(L)
		}
	}

	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
