package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 100010
	var dx = []int{0, 0, -1, 1}
	var dy = []int{1, -1, 0, 0}

	var n, t int
	fmt.Fscan(in, &n, &t)

	dmap := make([]int, 256)
	for i := 0; i < 4; i++ {
		dmap["UDLR"[i]] = i
	}

	x := make([]int, n)
	y := make([]int, n)
	d := make([]int, n)
	vp := make([]PPI, n)
	for i := 0; i < n; i++ {
		var c string
		fmt.Fscan(in, &x[i], &y[i], &c)
		vp[i] = PPI{pair{x[i], y[i]}, i}
		d[i] = dmap[c[0]]
	}

	var nxt [N][4]int
	for i := range nxt {
		for j := range nxt[i] {
			nxt[i][j] = -1
		}
	}
	sort.Slice(vp, func(i, j int) bool {
		if vp[i].x.x == vp[j].x.x {
			return vp[i].x.y < vp[j].x.y
		}
		return vp[i].x.x < vp[j].x.x
	}) // sort by x
	for i := 0; i < n-1; i++ {
		p := vp[i].x
		q := vp[i+1].x
		a := vp[i].y
		b := vp[i+1].y
		if p.x == q.x {
			nxt[a][0] = b // U
			nxt[b][1] = a // D
		}
	}

	sort.Slice(vp, func(a, b int) bool { // sort by y
		if vp[a].x.y != vp[b].x.y {
			return vp[a].x.y < vp[b].x.y
		}
		return vp[a].x.x < vp[b].x.x
	})

	for i := 0; i < n-1; i++ {
		p := vp[i].x
		q := vp[i+1].x
		a := vp[i].y
		b := vp[i+1].y
		if p.y == q.y {
			nxt[a][3] = b // R
			nxt[b][2] = a // L
		}
	}

	var get_dist func(int, int) int
	get_dist = func(i, j int) int { return abs(x[i]-x[j]) + abs(y[i]-y[j]) }

	var dist [N][4]int
	for i := 0; i < N; i++ {
		for j := 0; j < 4; j++ {
			dist[i][j] = t
		}
	}
	dist[0][d[0]] = 0
	q := &HeapPair{}
	heap.Push(q, pair{0, d[0]})
	for q.Len() > 0 {
		p := heap.Pop(q).(pair)
		tt := p.x
		id := p.y / 10
		dir := p.y % 10
		if dist[id][dir] < tt {
			continue
		}
		if nxt[id][dir] < 0 {
			continue
		}
		nn := nxt[id][dir]
		for i := 0; i < 2; i++ {
			if tt+get_dist(id, nn) < dist[nn][dir] {
				dist[nn][dir] = tt + get_dist(id, nn)
				heap.Push(q, pair{dist[nn][dir], nn*10 + dir})
			}
			dir = d[nn]
		}
	}

	for i := 0; i < n; i++ {
		dt := max(0, t-dist[i][d[i]])
		xx := x[i] + dt*dx[d[i]]
		yy := y[i] + dt*dy[d[i]]
		fmt.Println(xx, yy)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

type pair struct {
	x, y int
}

type PPI struct {
	x pair
	y int
}

type HeapPair []pair

func (h HeapPair) Len() int { return len(h) }
func (h HeapPair) Less(i, j int) bool {
	if h[i].x == h[j].x {
		return h[i].y < h[j].y
	}
	return h[i].x < h[j].x
}
func (h HeapPair) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapPair) Push(x interface{}) { *h = append(*h, x.(pair)) }

func (h *HeapPair) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
