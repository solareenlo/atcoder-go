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

	a := make([]int, N+1)
	for i := 1; i <= N; i++ {
		fmt.Fscan(in, &a[i])
	}

	type pair struct{ x, y int }
	adj := make([][]pair, 100005)
	for i := 0; i < M; i++ {
		var u, v, w int
		fmt.Fscan(in, &u, &v, &w)
		adj[u] = append(adj[u], pair{v, w})
		adj[v] = append(adj[v], pair{u, w})
	}

	q := &Heap{}
	heap.Push(q, tuple{0, 0, 1})
	vis := make([]bool, 100005)
	for q.Len() > 0 {
		d := (*q)[0].x
		s := (*q)[0].y
		u := (*q)[0].z
		heap.Pop(q)
		if u == N {
			fmt.Println(s + a[N])
			return
		}
		if vis[u] {
			continue
		}
		vis[u] = true
		for i := range adj[u] {
			v := adj[u][i].x
			w := adj[u][i].y
			heap.Push(q, tuple{d - w, s + a[u], v})
		}
	}
}

type tuple struct{ x, y, z int }
type Heap []tuple

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i].x > h[j].x || (h[i].x == h[j].x && h[i].y > h[j].y) }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(tuple)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
