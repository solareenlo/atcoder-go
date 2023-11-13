package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, C, D int
	fmt.Fscan(in, &N, &C, &D)
	A := make([][2]int, N+1)
	for i := 1; i <= N; i++ {
		fmt.Fscan(in, &A[i][0], &A[i][1])
	}
	A[0] = [2]int{1, 0}
	K := D
	for i := N; i >= 0; i-- {
		A[i][0] = K - A[i][0]
		K -= A[i][0]
	}
	pq := make(Heap, 0)
	heap.Init(&pq)
	ans := 0
	t := C
	for i := N; i >= 0; i-- {
		t -= A[i][0]
		for t < 0 && pq.Len() != 0 {
			t += pq[0]
			ans += 1
			heap.Pop(&pq)
		}
		if t < 0 {
			fmt.Println(-1)
			return
		}
		heap.Push(&pq, A[i][1])
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
