package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var h, w int
	fmt.Fscan(in, &h, &w)
	r := h * w

	g := NewWeightedGraph(2 * r)
	for i := 0; i < h; i++ {
		for j := 0; j < w-1; j++ {
			var a int
			fmt.Fscan(in, &a)
			g.add_edge(i*w+j, i*w+j+1, 0, a)
			g.add_edge(i*w+j+1, i*w+j, 0, a)
		}
	}
	for i := 0; i < h-1; i++ {
		for j := 0; j < w; j++ {
			var b int
			fmt.Fscan(in, &b)
			g.add_edge(i*w+j, i*w+w+j, 0, b)
			g.add_edge(i*w+j+w, i*w+j+r, 0, 1)
			g.add_edge(i*w+j+w+r, i*w+j+r, 0, 1)
			g.add_edge(i*w+j+r, i*w+j, 0, 1)
		}
	}
	fmt.Println(g.flow(0, r-1))
}

type pair struct{ dist, p int }
type Heap []pair

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i].dist < h[j].dist }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(pair)) }
func (h *Heap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return x
}

type edge struct {
	from, to   int
	capa, flow int
	cost       int
}
type _edge struct {
	to, rev int
	capa    int
	cost    int
}
type WeightedGraph struct {
	n      int
	_edges [][]edge
}

func NewWeightedGraph(n int) WeightedGraph {
	g := WeightedGraph{n, make([][]edge, n)}
	return g
}
func (g *WeightedGraph) add_edge(from, to, capa, cost int) int {
	g._edges[from] = append(g._edges[from], edge{to: to, cost: cost})
	return 1
}
func (g *WeightedGraph) flow(s, t int) int {
	n := g.n
	INF := 1 << 60
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = INF
	}
	pq := new(Heap)
	heap.Push(pq, pair{s, 0})
	for pq.Len() != 0 {
		cur := heap.Pop(pq).(pair)
		if dp[cur.p] < cur.dist {
			continue
		}
		for i := 0; i < len(g._edges[cur.p]); i++ {
			nv := g._edges[cur.p][i].to
			ndist := cur.dist + g._edges[cur.p][i].cost
			if dp[nv] > ndist {
				dp[nv] = ndist
				heap.Push(pq, pair{ndist, nv})
			}
		}
	}
	return dp[t]
}
