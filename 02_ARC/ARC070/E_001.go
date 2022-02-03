package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

var (
	p int
	q int
	a int
	b int
	L = &HeapL{}
	R = &HeapR{}
)

func topL() int {
	p = (*L)[0] + a
	return p
}

func topR() int {
	q = (*R)[0] + b
	return q
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	heap.Init(L)
	heap.Init(R)
	ans := 0
	for i := 0; i < n; i++ {
		var l, r int
		fmt.Fscan(in, &l, &r)
		leng := r - l
		b += leng
		heap.Push(L, r-a)
		heap.Push(R, r-b)
		for topL() > topR() {
			ans += p - q
			heap.Pop(L)
			heap.Pop(R)
			heap.Push(L, q-a)
			heap.Push(R, p-b)
		}
		a -= leng
	}
	fmt.Println(ans)
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
