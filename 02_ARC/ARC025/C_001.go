package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

type P struct{ c, v int }
type edge struct{ to, cost int }

var (
	G    = make([][]edge, 0)
	dist = make([]int, 0)
	N    int
)

func dijkstra(s int) {
	for i := 0; i < N; i++ {
		dist[i] = 1 << 60
	}
	dist[s] = 0
	Q := &Heap{}
	heap.Push(Q, P{0, s})
	heap.Init(Q)
	for Q.Len() > 0 {
		c := (*Q)[0].c
		v := (*Q)[0].v
		heap.Pop(Q)
		if dist[v] < c {
			continue
		}
		for _, e := range G[v] {
			if dist[e.to] > c+e.cost {
				dist[e.to] = c + e.cost
				heap.Push(Q, P{dist[e.to], e.to})
			}
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var M, R, T int
	fmt.Fscan(in, &N, &M, &R, &T)

	G = make([][]edge, N)
	for i := 0; i < M; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		a--
		b--
		G[a] = append(G[a], edge{b, c})
		G[b] = append(G[b], edge{a, c})
	}

	dist = make([]int, N)
	D := make([]int, N)
	ans := 0
	for i := 0; i < N; i++ {
		dijkstra(i)
		sort.Ints(dist)
		for j := 0; j < N; j++ {
			D[j] = dist[j] * T
		}
		for j := 1; j < N; j++ {
			idx := upperBound(D[1:], dist[j]*R) + 1
			ans += N - idx
		}
	}

	if T > R {
		fmt.Println(ans - N*(N-1))
	} else {
		fmt.Println(ans)
	}
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}

type Heap []P

func (h Heap) Len() int { return len(h) }
func (h Heap) Less(i, j int) bool {
	return h[i].c < h[j].c
}
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(P)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
