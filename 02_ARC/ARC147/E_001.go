package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 1000001

	var n int
	fmt.Fscan(in, &n)

	v1 := make([]int, N+10)
	v2 := make([]int, N)
	a := &HeapIntLess{}
	b := &HeapIntLess{}
	q1 := &HeapPair{}
	ans := 0
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &v1[i], &v2[i])
		if v1[i] < v2[i] {
			heap.Push(a, v1[i])
			heap.Push(b, v2[i])
		} else {
			heap.Push(q1, pair{v1[i], v2[i]})
			ans++
		}
	}

	c := &HeapIntGreater{}
	for b.Len() > 0 {
		val := heap.Pop(b).(int)
		if a.Len() > 0 && (*a)[0] >= val {
			heap.Pop(a)
			continue
		}
		for q1.Len() > 0 && (*q1)[0].x >= val {
			heap.Push(c, heap.Pop(q1).(pair).y)
		}
		if c.Len() == 0 {
			fmt.Println(-1)
			return
		}
		heap.Push(b, heap.Pop(c).(int))
		ans--
	}
	fmt.Println(ans)
}

type HeapIntLess []int

func (h HeapIntLess) Len() int            { return len(h) }
func (h HeapIntLess) Less(i, j int) bool  { return h[i] > h[j] }
func (h HeapIntLess) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapIntLess) Push(x interface{}) { *h = append(*h, x.(int)) }

func (h *HeapIntLess) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type HeapIntGreater []int

func (h HeapIntGreater) Len() int            { return len(h) }
func (h HeapIntGreater) Less(i, j int) bool  { return h[i] < h[j] }
func (h HeapIntGreater) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapIntGreater) Push(x interface{}) { *h = append(*h, x.(int)) }

func (h *HeapIntGreater) Pop() interface{} {
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

func (h HeapPair) Len() int { return len(h) }
func (h HeapPair) Less(i, j int) bool {
	if h[i].x == h[j].x {
		return h[i].y > h[j].y
	}
	return h[i].x > h[j].x
}
func (h HeapPair) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapPair) Push(x interface{}) { *h = append(*h, x.(pair)) }

func (h *HeapPair) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
