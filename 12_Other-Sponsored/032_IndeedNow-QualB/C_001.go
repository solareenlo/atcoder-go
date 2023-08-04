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
	G := make([][]int, N+1)
	for i := 0; i < N-1; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		G[a] = append(G[a], b)
		G[b] = append(G[b], a)
	}
	pq := &Heap{}
	heap.Push(pq, 1)
	seen := make([]bool, N+1)
	seen[1] = true
	for pq.Len() > 0 {
		u := heap.Pop(pq).(int)
		N--
		if N != 0 {
			fmt.Printf("%d ", u)
		} else {
			fmt.Println(u)
		}
		for _, v := range G[u] {
			if !seen[v] {
				seen[v] = true
				heap.Push(pq, v)
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
