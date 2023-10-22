package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)
	q := make(HeapPair, 0)
	heap.Init(&q)
	for n > 0 {
		n--
		var a, b int
		fmt.Fscan(in, &a, &b)
		heap.Push(&q, pair{-a, -b})
	}
	for k > 0 {
		k--
		a := heap.Pop(&q).(pair)
		n -= a.x
		a.x += a.y
		heap.Push(&q, a)
	}
	fmt.Println(n)
}

type pair struct {
	x, y int
}

type HeapPair []pair

func (h HeapPair) Len() int            { return len(h) }
func (h HeapPair) Less(i, j int) bool  { return h[i].x > h[j].x }
func (h HeapPair) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapPair) Push(x interface{}) { *h = append(*h, x.(pair)) }

func (h *HeapPair) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
