package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var G [1 << 17][]int
	var dist [2 << 17][2]int

	var N, M int
	fmt.Fscan(in, &N, &M)
	U := make([]int, 2*M)
	V := make([]int, 2*M)
	W := make([]int, 2*M)
	for i := 0; i < M; i++ {
		var u, v, w int
		fmt.Fscan(in, &u, &v, &w)
		u--
		v--
		U[i*2] = u
		V[i*2] = v
		W[i*2] = w
		G[u] = append(G[u], i*2)
		U[i*2+1] = v
		V[i*2+1] = u
		W[i*2+1] = w
		G[v] = append(G[v], i*2+1)
	}
	for i := 0; i < 2*M; i++ {
		dist[i][0] = int(9e18)
		dist[i][1] = int(9e18)
	}
	pq := &HeapPair{}
	for _, e := range G[0] {
		dist[e][0] = W[e]
		dist[e][1] = W[e]
		heap.Push(pq, Pair{-W[e], pair{e, 0}})
		heap.Push(pq, Pair{-W[e], pair{e, 1}})
	}
	for pq.Len() > 0 {
		c := -(*pq)[0].x
		ei := (*pq)[0].y.x
		f := (*pq)[0].y.y
		heap.Pop(pq)
		if dist[ei][f] < c {
			continue
		}
		if V[ei] == N-1 {
			fmt.Println(c)
			return
		}
		for _, e := range G[V[ei]] {
			if (f == 0 && W[ei] <= W[e]) || (f != 0 && W[ei] >= W[e]) {
				continue
			}
			nc := c + W[e]
			tmp := 0
			if f == 0 {
				tmp = 1
			}
			if dist[e][tmp] > nc {
				dist[e][tmp] = nc
				heap.Push(pq, Pair{-nc, pair{e, tmp}})
			}
		}
	}
	fmt.Println(-1)
}

type pair struct {
	x, y int
}

type Pair struct {
	x int
	y pair
}

type HeapPair []Pair

func (h HeapPair) Len() int            { return len(h) }
func (h HeapPair) Less(i, j int) bool  { return h[i].x > h[j].x }
func (h HeapPair) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapPair) Push(x interface{}) { *h = append(*h, x.(Pair)) }

func (h *HeapPair) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
