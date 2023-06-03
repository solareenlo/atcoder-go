package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	C := make([]int, n+1)
	for i := range C {
		C[i] = int(1e9)
	}
	A := make([][]pair, n+1)
	B := &HeapPair{}
	for i := 0; i < m; i++ {
		var a, b, c, d int
		fmt.Fscan(in, &a, &b, &c, &d)
		A[a] = append(A[a], pair{c*10000 - d, b})
		A[b] = append(A[b], pair{c*10000 - d, a})
	}
	heap.Push(B, pair{0, 1})
	for B.Len() > 0 {
		tmp := heap.Pop(B).(pair)
		o := tmp.x
		p := tmp.y
		for _, tmp1 := range A[p] {
			q := tmp1.x
			r := tmp1.y
			if C[r] > o+q {
				C[r] = o + q
				heap.Push(B, pair{o + q, r})
			}
		}
	}
	a := (C[n] + 9999) / 10000
	fmt.Println(a, (C[n]-a*10000)*-1)
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
