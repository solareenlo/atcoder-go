package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	g := make([][]pair, n)
	for m > 0 {
		m--
		var u, v, c int
		fmt.Fscan(in, &u, &v, &c)
		u--
		v--
		g[u] = append(g[u], pair{v, c})
		g[v] = append(g[v], pair{u, c})
	}
	dp := make([]int, n)
	par := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1 << 30
		par[i] = -2
	}
	pq := &HeapPair{}
	dp[n-1] = 0
	heap.Push(pq, pair{0, n - 1})
	for pq.Len() != 0 {
		tmp := heap.Pop(pq).(pair)
		d := tmp.x
		v := tmp.y
		if dp[v] > d {
			continue
		}
		for _, p := range g[v] {
			u := p.x
			c := p.y
			if d+c >= dp[u] {
				continue
			}
			dp[u] = d + c
			par[u] = v
			heap.Push(pq, pair{dp[u], u})
		}
	}
	v := 0
	for v != -2 {
		if v != 0 {
			fmt.Print(" ")
		}
		fmt.Print(v + 1)
		v = par[v]
	}
	fmt.Println()
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
