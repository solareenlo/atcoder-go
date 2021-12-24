package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

const MX = 100005
const INF = 1000000000000000

var (
	n    int
	dist = [2][MX]int{}
	ans  = [MX]int{}
	pq   = &Heap{}
	G    = [2][MX][]P{}
)

func dijkstra(s, f int) {
	for i := 0; i < n; i++ {
		dist[f][i] = INF
	}
	dist[f][s] = 0
	heap.Push(pq, P{0, s})
	for pq.Len() > 0 {
		ac := (*pq)[0].x
		from := (*pq)[0].y
		heap.Pop(pq)
		if dist[f][from] == ac {
			for _, it := range G[f][from] {
				cost := it.x
				to := it.y
				cost += ac
				if dist[f][to] > cost {
					dist[f][to] = cost
					heap.Push(pq, P{cost, to})
				}
			}
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var m int
	s := [2]int{}
	fmt.Fscan(in, &n, &m, &s[0], &s[1])
	s[0]--
	s[1]--
	for i := 0; i < m; i++ {
		var u, v int
		a := [2]int{}
		fmt.Fscan(in, &u, &v, &a[0], &a[1])
		u--
		v--
		for j := 0; j < 2; j++ {
			G[j][u] = append(G[j][u], P{a[j], v})
			G[j][v] = append(G[j][v], P{a[j], u})
		}
	}
	for i := 0; i < 2; i++ {
		dijkstra(s[i], i)
	}
	for i := n - 1; i >= 0; i-- {
		ans[i] = max(ans[i+1], INF-dist[0][i]-dist[1][i])
	}

	for i := 0; i < n; i++ {
		fmt.Fprintln(out, ans[i])
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type P struct{ x, y int }
type Heap []P

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i].x < h[j].x }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(P)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
