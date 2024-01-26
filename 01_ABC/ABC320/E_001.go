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

	const N = 200009
	var ans [N]int

	var n, m int
	fmt.Fscan(in, &n, &m)
	q := &Heap{}
	pq := &HeapPair{}
	for i := 0; i < n; {
		heap.Push(q, i)
		i++
	}
	for ; m > 0; m-- {
		var t, w, s int
		fmt.Fscan(in, &t, &w, &s)
		for pq.Len() > 0 && (*pq)[0].x <= t {
			heap.Push(q, (*pq)[0].y)
			heap.Pop(pq)
		}
		if q.Len() > 0 {
			ans[(*q)[0]] += w
			heap.Push(pq, pair{t + s, (*q)[0]})
			heap.Pop(q)
		}
	}
	for i := 0; i < n; {
		fmt.Fprintln(out, ans[i])
		i++
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

type pair struct {
	x, y int
}

type HeapPair []pair

func (h HeapPair) Len() int            { return len(h) }
func (h HeapPair) Less(i, j int) bool  { return h[i].x < h[j].x }
func (h HeapPair) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapPair) Push(x interface{}) { *h = append(*h, x.(pair)) }

func (h *HeapPair) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
