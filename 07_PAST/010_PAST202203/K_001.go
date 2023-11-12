package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

const INF = int(1e18)

func dijkstra(graph [][]pair, s int) []int {
	dist := make([]int, len(graph))
	for i := range dist {
		dist[i] = INF
	}
	q := make(HeapPair, 0)
	heap.Init(&q)
	dist[s] = 0
	heap.Push(&q, pair{dist[s], s})
	for q.Len() > 0 {
		tmp := heap.Pop(&q).(pair)
		uc, ui := tmp.x, tmp.y
		if uc != dist[ui] {
			continue
		}
		for _, tmp := range graph[ui] {
			vi, vc := tmp.x, tmp.y
			if dist[vi] > uc+vc {
				dist[vi] = uc + vc
				heap.Push(&q, pair{dist[vi], vi})
			}
		}
	}
	return dist
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(in, &N, &M)
	G := make([][]pair, N)
	H := make([][]pair, N)
	for M > 0 {
		M--
		var u, v, c int
		fmt.Fscan(in, &u, &v, &c)
		u--
		v--
		G[u] = append(G[u], pair{v, c})
		H[v] = append(H[v], pair{u, c})
	}

	ds := dijkstra(G, 0)
	dt := dijkstra(H, N-1)
	for k := 0; k < N; k++ {
		if ds[k] == INF || dt[k] == INF {
			fmt.Println(-1)
		} else {
			fmt.Println(ds[k] + dt[k])
		}
	}
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
