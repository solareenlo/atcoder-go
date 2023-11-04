package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, K int
	fmt.Fscan(in, &N, &K)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
	}
	pq := make(Heap, 0)
	heap.Init(&pq)
	heap.Push(&pq, 0)
	cnt := -1
	pr := -1
	for {
		x := heap.Pop(&pq).(int)
		if pr != x {
			pr = x
			cnt++
			if cnt == K {
				fmt.Println(x)
				break
			}
			for i := 0; i < N; i++ {
				heap.Push(&pq, x+A[i])
			}
		}
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
