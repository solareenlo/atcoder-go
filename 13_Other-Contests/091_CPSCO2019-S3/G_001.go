package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	EPS := 1e-12

	var n, X int
	fmt.Fscan(in, &n, &X)
	a := make([]int, n)
	A := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		A += a[i]
	}

	que := &HeapPair{}
	for i := 0; i < n; i++ {
		p := a[i] * X / A
		if p > 0 {
			heap.Push(que, Pair{1.0/float64(a[i]) + float64(i)*EPS, P{i, p}})
		}
		heap.Push(que, Pair{-1.0/float64(a[i]) + float64(i)*EPS, P{i, X}})
		tmp := math.Abs(float64(p*A-a[i]*X)) / float64(a[i]) / float64(A)
		tmp -= math.Abs(float64((p+1)*A-a[i]*X)) / float64(a[i]) / float64(A)
		heap.Push(que, Pair{tmp + float64(i)*EPS, P{i, 1}})
	}

	x := make([]int, n)
	r := X
	for r > 0 {
		tmp := heap.Pop(que).(Pair)
		i := tmp.b.a
		p := tmp.b.b
		p = min(p, r)
		x[i] += p
		r -= p
	}
	for i := 0; i < n; i++ {
		fmt.Println(x[i])
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type P struct {
	a, b int
}

type Pair struct {
	a float64
	b P
}

type HeapPair []Pair

func (h HeapPair) Len() int            { return len(h) }
func (h HeapPair) Less(i, j int) bool  { return h[i].a > h[j].a }
func (h HeapPair) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapPair) Push(x interface{}) { *h = append(*h, x.(Pair)) }

func (h *HeapPair) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
