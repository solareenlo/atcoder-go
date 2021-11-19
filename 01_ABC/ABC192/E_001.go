package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m, x, y int
	fmt.Fscan(in, &n, &m, &x, &y)
	x--
	y--

	type tuple struct{ to, time, k int }
	g := make([][]tuple, n)

	for i := 0; i < m; i++ {
		var u, v, t, k int
		fmt.Fscan(in, &u, &v, &t, &k)
		u--
		v--
		g[u] = append(g[u], tuple{v, t, k})
		g[v] = append(g[v], tuple{u, t, k})
	}

	dist := make([]int, n)
	for i := range dist {
		dist[i] = 1 << 60
	}
	dist[x] = 0

	q := &Heap{}
	heap.Init(q)
	heap.Push(q, pair{0, x})
	for q.Len() > 0 {
		now := heap.Pop(q).(pair)
		if dist[now.to] < now.dist {
			continue
		}
		for _, v := range g[now.to] {
			var ndist = dist[now.to] + v.time + (v.k-dist[now.to]%v.k)%v.k
			if dist[v.to] > ndist {
				dist[v.to] = ndist
				heap.Push(q, pair{ndist, v.to})
			}
		}
	}

	if dist[y] == 1<<60 {
		dist[y] = -1
	}
	fmt.Println(dist[y])
}

type pair struct{ dist, to int }
type Heap []pair

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i].dist < h[j].dist }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(pair)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
