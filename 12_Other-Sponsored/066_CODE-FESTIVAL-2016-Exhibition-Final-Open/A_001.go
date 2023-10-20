package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)
	Q := make(HeapPair, 0)
	heap.Init(&Q)
	for i := 0; i < N; i++ {
		var a int
		fmt.Fscan(in, &a)
		heap.Push(&Q, pair{a, -1})
	}
	for i := 0; i < N; i++ {
		var a int
		fmt.Fscan(in, &a)
		heap.Push(&Q, pair{a, 1})
	}
	a := 0
	A := 1
	const MOD = 1000000007
	for Q.Len() > 0 {
		i := Q[0].y
		if a != i*abs(a) {
			A = (A * abs(a)) % MOD
		}
		a += i
		heap.Pop(&Q)
	}
	fmt.Println(A)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
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
