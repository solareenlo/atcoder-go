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

	var n, m int
	fmt.Fscan(in, &n, &m)

	adj := make([][]Edge, n)
	for i := 0; i < m; i++ {
		var a, b, c, d int
		fmt.Fscan(in, &a, &b, &c, &d)
		a--
		b--
		adj[a] = append(adj[a], Edge{b, c, d})
		adj[b] = append(adj[b], Edge{a, c, d})
	}

	dist := make([]int, n)
	for i := range dist {
		dist[i] = -1
	}
	pq := &Heap{}
	heap.Push(pq, pair{0, 0})
	for pq.Len() > 0 {
		t0 := -(*pq)[0].pos
		u := (*pq)[0].time
		heap.Pop(pq)
		if dist[u] != -1 {
			continue
		}
		dist[u] = t0
		for _, e := range adj[u] {
			if dist[e.v] == -1 {
				heap.Push(pq, pair{-e.tmin(t0), e.v})
			}
		}
	}

	fmt.Println(dist[n-1])
}

type Edge struct{ v, c, d int }

func (e Edge) arrive(t int) int { return t + e.c + e.d/(t+1) }
func (e Edge) tmin(t0 int) int {
	return e.arrive(max(t0, int(math.Round(math.Sqrt(float64(e.d)))-1)))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type pair struct{ pos, time int }
type Heap []pair

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i].pos > h[j].pos }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(pair)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
