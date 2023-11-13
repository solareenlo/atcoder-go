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

	var N, M int
	fmt.Fscan(in, &N, &M)

	adj := make([][]int, N)
	deg := make([]int, N)
	for i := 0; i < M; i++ {
		var A, B int
		fmt.Fscan(in, &A, &B)
		A--
		B--
		adj[A] = append(adj[A], B)
		deg[B]++
	}

	q := make(Heap, 0)
	heap.Init(&q)
	for i := 0; i < N; i++ {
		if deg[i] == 0 {
			heap.Push(&q, i)
		}
	}

	for i := 0; i < N; i++ {
		x := q[0]
		if i == N-1 {
			fmt.Fprintln(out, x+1)
		} else {
			fmt.Fprintf(out, "%d ", x+1)
		}
		heap.Pop(&q)

		for _, y := range adj[x] {
			deg[y]--
			if deg[y] == 0 {
				heap.Push(&q, y)
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
