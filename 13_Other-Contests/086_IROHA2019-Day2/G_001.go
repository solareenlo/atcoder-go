package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	G := make([][]tuple, 1000)
	for i := 0; i < m; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		a--
		b--
		G[a] = append(G[a], tuple{b, 0, c})
		G[b] = append(G[b], tuple{a, 0, c})
	}
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		G[i] = append(G[i], tuple{i, x, y})
	}
	Q := &HeapTuple{}
	heap.Push(Q, tuple{0, 0, 0})
	var D [1000][1000]int
	for i := range D {
		for j := range D[i] {
			D[i][j] = int(1e18)
		}
	}
	D[0][0] = 0
	for Q.Len() > 0 {
		tmp := heap.Pop(Q).(tuple)
		d, u, z := tmp.x, tmp.y, tmp.z
		if u == n-1 && z == k {
			fmt.Println(d)
			return
		}
		if D[u][z] < d {
			continue
		}
		for _, g := range G[u] {
			v, x, y := g.x, g.y, g.z
			w := min(z+x, k)
			if D[v][w] > d+y {
				D[v][w] = d + y
				heap.Push(Q, tuple{d + y, v, w})
			}
		}
	}
	fmt.Println(-1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type tuple struct {
	x, y, z int
}

type HeapTuple []tuple

func (h HeapTuple) Len() int            { return len(h) }
func (h HeapTuple) Less(i, j int) bool  { return h[i].x < h[j].x }
func (h HeapTuple) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapTuple) Push(x interface{}) { *h = append(*h, x.(tuple)) }

func (h *HeapTuple) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
