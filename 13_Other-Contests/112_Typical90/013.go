package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type P struct {
	x, y int
}

var N int
var g [100000][]P

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var M int
	fmt.Fscan(in, &N, &M)
	for M > 0 {
		M--
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		a--
		b--
		g[a] = append(g[a], P{c, b})
		g[b] = append(g[b], P{c, a})
	}
	v1 := dijkstra(0)
	v2 := dijkstra(N - 1)
	for i := 0; i < N; i++ {
		fmt.Fprintln(out, v1[i]+v2[i])
	}
}

func dijkstra(f int) []int {
	d := make([]int, N)
	for i := range d {
		d[i] = -1
	}
	que := &HeapPair{}
	heap.Push(que, P{0, f})
	for que.Len() > 0 {
		tmp := heap.Pop(que).(P)
		c := tmp.x
		x := tmp.y
		if d[x] >= 0 {
			continue
		}
		d[x] = c
		for _, ne := range g[x] {
			ne.x += c
			heap.Push(que, ne)
		}
	}
	return d
}

type HeapPair []P

func (h HeapPair) Len() int            { return len(h) }
func (h HeapPair) Less(i, j int) bool  { return h[i].x < h[j].x }
func (h HeapPair) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapPair) Push(x interface{}) { *h = append(*h, x.(P)) }

func (h *HeapPair) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
