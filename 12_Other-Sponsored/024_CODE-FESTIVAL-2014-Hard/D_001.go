package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

type pair struct {
	x, y float64
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, V int
	var xs, ys, xg, yg float64
	fmt.Fscan(in, &n, &V, &xs, &ys, &xg, &yg)
	x := make([]float64, n+2)
	y := make([]float64, n+2)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x[i], &y[i])
	}
	x[n] = xs
	y[n] = ys
	x[n+1] = xg
	y[n+1] = yg
	v := make([][]int, n+2)
	w := make([][]int, n-1)
	for i := 0; i < n+2; i++ {
		for j := 0; j+1 < n; j++ {
			k := len(x)
			v[i] = append(v[i], k)
			w[j] = append(w[j], k)
			p := calc(1.0/float64(V), x[i], y[i], x[j], y[j], x[j+1], y[j+1])
			x = append(x, p.x)
			y = append(y, p.y)
			k = len(x)
			v[i] = append(v[i], k)
			w[j] = append(w[j], k)
			p = calc(1.0/float64(V), x[i], y[i], x[j+1], y[j+1], x[j], y[j])
			x = append(x, p.x)
			y = append(y, p.y)
		}
	}
	k := len(x)
	G := make([][]edge, k)
	for i := 0; i < n+2; i++ {
		for _i := i + 1; _i < n+2; _i++ {
			add_edge(i, _i, 1, x, y, G)
		}
	}
	for i := 0; i < n+2; i++ {
		for l := 0; l < len(v[i]); l++ {
			_i := v[i][l]
			add_edge(i, _i, 1, x, y, G)
		}
	}
	for j := 0; j+1 < n; j++ {
		w[j] = append(w[j], j)
		w[j] = append(w[j], j+1)
		ws := len(w[j])
		for l := 0; l < ws; l++ {
			for _l := l + 1; _l < ws; _l++ {
				add_edge(w[j][l], w[j][_l], V, x, y, G)
			}
		}
	}
	d := make([]float64, k)
	dijkstra(k, G, n, d)
	fmt.Println(d[n+1])
}

type edge struct {
	u, v int
	w    float64
}

func hoge(x1, y1, x2, y2 float64) float64 {
	return (x1*x2 + y1*y2) / math.Hypot(x1, y1) / math.Hypot(x2, y2)
}

func calc(c, x0, y0, x1, y1, x2, y2 float64) pair {
	xl := x1
	yl := y1
	xu := x2
	yu := y2
	for k := 0; k < 100; k++ {
		xm := (xl + xu) / 2
		ym := (yl + yu) / 2
		_c := hoge(x2-x1, y2-y1, x0-xm, y0-ym)
		if _c > c {
			xl = xm
			yl = ym
		} else {
			xu = xm
			yu = ym
		}
	}
	return pair{xl, yl}
}

func add_edge(u, v, V int, x, y []float64, G [][]edge) {
	w := math.Hypot(x[u]-x[v], y[u]-y[v]) / float64(V)
	e := edge{u, v, w}
	_e := edge{v, u, w}
	G[u] = append(G[u], e)
	G[v] = append(G[v], _e)
}

func dijkstra(n int, G [][]edge, s int, d []float64) {
	for i := range d {
		d[i] = 1e18
	}
	d[s] = 0
	q := &HeapPair{}
	heap.Push(q, Pair{0, s})
	for q.Len() > 0 {
		p := heap.Pop(q).(Pair)
		u := p.y
		if p.x > d[u] {
			continue
		}
		for i := 0; i < len(G[u]); i++ {
			e := G[u][i]
			if d[e.v] > d[u]+e.w {
				d[e.v] = d[u] + e.w
				heap.Push(q, Pair{d[e.v], e.v})
			}
		}
	}
}

type Pair struct {
	x float64
	y int
}

type HeapPair []Pair

func (h HeapPair) Len() int            { return len(h) }
func (h HeapPair) Less(i, j int) bool  { return h[i].x < h[j].x }
func (h HeapPair) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapPair) Push(x interface{}) { *h = append(*h, x.(Pair)) }

func (h *HeapPair) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
