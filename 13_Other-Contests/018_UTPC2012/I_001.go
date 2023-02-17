package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

var m, n, q int
var a [][]int
var cur [11]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &m, &n, &q)
	a = make([][]int, 10005)
	for i := range a {
		a[i] = make([]int, 11)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &a[i][j])
		}
	}
	dijkstra()
	var seg seg
	seg.init(0, n-1, 1)
	for q > 0 {
		q--
		var sx, sy, tx, ty int
		fmt.Fscan(in, &sy, &sx, &ty, &tx)
		sx--
		sy--
		tx--
		ty--
		if sx > tx {
			sx, tx = tx, sx
			sy, ty = ty, sy
		}
		for i := range cur {
			cur[i] = int(1e18)
		}
		cur[sy] = 0
		seg.query(sx, tx, 0, n-1, 1)
		fmt.Fprintln(out, cur[ty])
	}
}

var dist [10005][11][11]int

var pq *Heap

type node struct {
	x, ys, ye int
	dist      int
}

func dijkstra() {
	for i := range dist {
		for j := range dist[i] {
			for k := range dist[i][j] {
				dist[i][j][k] = int(1e18)
			}
		}
	}
	update := func(n node) {
		if dist[n.x][n.ys][n.ye] > n.dist {
			dist[n.x][n.ys][n.ye] = n.dist
			heap.Push(pq, n)
		}
	}
	pq = &Heap{}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			update(node{i, j, j, a[i][j]})
		}
	}
	for pq.Len() > 0 {
		x := heap.Pop(pq).(node)
		if dist[x.x][x.ys][x.ye] != x.dist {
			continue
		}
		for i := 0; i < x.ys; i++ {
			update(node{x.x, i, x.ye, x.dist + dist[x.x][i][x.ys] - a[x.x][x.ys]})
		}
		for i := x.ye + 1; i < m; i++ {
			update(node{x.x, x.ys, i, x.dist + dist[x.x][x.ye][i] - a[x.x][x.ye]})
		}
		if x.ys > 0 {
			update(node{x.x, x.ys - 1, x.ye, x.dist + a[x.x][x.ys-1]})
		}
		if x.ye+1 < m {
			update(node{x.x, x.ys, x.ye + 1, x.dist + a[x.x][x.ye+1]})
		}
		if x.ys != x.ye {
			update(node{x.x, x.ys + 1, x.ye, x.dist + a[x.x][x.ys+1]})
			update(node{x.x, x.ys, x.ye - 1, x.dist + a[x.x][x.ye-1]})
		}
		if x.x > 0 {
			update(node{x.x - 1, x.ys, x.ye, x.dist + a[x.x-1][x.ys] + a[x.x-1][x.ye]})
		}
		if x.x+1 < n {
			update(node{x.x + 1, x.ys, x.ye, x.dist + a[x.x+1][x.ys] + a[x.x+1][x.ye]})
		}
	}
}

type Heap []node

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i].dist < h[j].dist }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(node)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type NODE struct {
	adj [11][11]int
}

type seg struct {
	tree [33000]NODE
}

func (seg *seg) init(s, e, p int) {
	if s == e {
		for i := 0; i < m; i++ {
			for j := 0; j < m; j++ {
				if i <= j {
					seg.tree[p].adj[i][j] = dist[s][i][j]
				} else {
					seg.tree[p].adj[i][j] = dist[s][j][i]
				}
			}
		}
		return
	}
	m := (s + e) / 2
	seg.init(s, m, 2*p)
	seg.init(m+1, e, 2*p+1)
	seg.tree[p] = seg.merge(seg.tree[2*p], seg.tree[2*p+1])
}

func (seg *seg) merge(a, b NODE) NODE {
	var ret NODE
	for i := range ret.adj {
		for j := range ret.adj[i] {
			ret.adj[i][j] = int(1e18)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < m; k++ {
				ret.adj[i][j] = min(ret.adj[i][j], a.adj[i][k]+b.adj[k][j])
			}
		}
	}
	return ret
}

func (seg *seg) query(s, e, ps, pe, p int) {
	if e < ps || pe < s {
		return
	}
	if s <= ps && pe <= e {
		var prev [11]int
		for i := range prev {
			prev[i] = cur[i]
		}
		for i := range cur {
			cur[i] = int(1e18)
		}
		for i := 0; i < m; i++ {
			for j := 0; j < m; j++ {
				cur[j] = min(cur[j], prev[i]+seg.tree[p].adj[i][j])
			}
		}
		return
	}
	pm := (ps + pe) / 2
	seg.query(s, e, ps, pm, 2*p)
	seg.query(s, e, pm+1, pe, 2*p+1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
