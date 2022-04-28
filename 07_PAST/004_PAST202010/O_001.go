package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(in, &N, &M)
	Z := 0
	G := make([][]pair, N+1)
	for i := 0; i < N; i++ {
		var a int
		fmt.Fscan(in, &a)
		Z += a
		G[i] = append(G[i], pair{i + 1, a})
		G[i+1] = append(G[i+1], pair{i, 0})
	}
	for i := 0; i < M; i++ {
		var l, r, c int
		fmt.Fscan(in, &l, &r, &c)
		G[l-1] = append(G[l-1], pair{r, c})
	}

	pq := &Heap{}
	heap.Push(pq, pair{0, 0})
	seen := make([]int, N+1)
	for i := range seen {
		seen[i] = -1 << 60
	}
	for pq.Len() > 0 {
		a := (*pq)[0].x
		b := (*pq)[0].y
		heap.Pop(pq)
		if seen[b] > a {
			continue
		}
		for _, x := range G[b] {
			i := x.x
			d := x.y
			if seen[i] < a-d {
				seen[i] = a - d
				heap.Push(pq, pair{a - d, i})
			}
		}
	}
	fmt.Println(Z + seen[N])
}

type pair struct{ x, y int }
type Heap []pair

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i].x > h[j].x }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(pair)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
