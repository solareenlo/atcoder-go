package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

const INF = int(1e18)

type edge struct {
	to, wt int
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m, s, t int
	fmt.Fscan(in, &n, &m, &s, &t)
	s--
	t--
	G := make([][]edge, n)
	for i := 0; i < m; i++ {
		var u, v, c int
		fmt.Fscan(in, &u, &v, &c)
		u--
		v--
		G[u] = append(G[u], edge{v, c})
		G[v] = append(G[v], edge{u, c})
	}

	d1 := Dijkstra(G, s)
	d2 := Dijkstra(G, t)
	for i := 0; i < n; i++ {
		if d1[i] < INF && d1[i] == d2[i] {
			fmt.Println(i + 1)
			return
		}
	}
	fmt.Println(-1)
}

func Dijkstra(G [][]edge, s int) []int {
	n := len(G)
	d := make([]int, n)
	for i := range d {
		d[i] = INF
	}
	d[s] = 0
	Q := &HeapPair{}
	heap.Push(Q, pair{0, s})
	for Q.Len() > 0 {
		q := heap.Pop(Q).(pair)
		d0 := -q.x
		u := q.y
		if d0 > d[u] {
			continue
		}
		for _, e := range G[u] {
			v := e.to
			if d[v] > d[u]+e.wt {
				d[v] = d[u] + e.wt
				heap.Push(Q, pair{d[v], v})
			}
		}
	}
	return d
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
