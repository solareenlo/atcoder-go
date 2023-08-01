package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = int(1e18)

	var N, M, s, g int
	fmt.Fscan(in, &N, &M, &s, &g)

	var a, b, c, d [100005]int
	for i := 1; i <= N; i++ {
		fmt.Fscan(in, &a[i], &b[i], &c[i])
		d[i] = INF
	}

	adj := make([][]Pair, 100005)
	for M > 0 {
		M--
		var u, v, w int
		fmt.Fscan(in, &u, &v, &w)
		adj[u] = append(adj[u], Pair{v, w})
		adj[v] = append(adj[v], Pair{u, w})
	}

	q := &HeapPair{}
	d[s] = 0
	heap.Push(q, Pair{0, s})
	for q.Len() > 0 {
		tmp := heap.Pop(q).(Pair)
		dist := tmp.x
		u := tmp.y
		if -dist != d[u] {
			continue
		}
		if u == g {
			fmt.Println(d[g])
			return
		}
		if d[u] < c[u] {
			d[u] = c[u]
		} else {
			d[u] -= c[u]
			rem := d[u] % (a[u] + b[u])
			if rem >= a[u] {
				d[u] = d[u] - rem + a[u] + b[u]
			}
			d[u] += c[u]
		}
		for _, t := range adj[u] {
			v := t.x
			w := t.y
			if d[u]+w < d[v] {
				d[v] = d[u] + w
				heap.Push(q, Pair{-d[v], v})
			}
		}
	}
}

type Pair struct {
	x, y int
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
