package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type pair struct {
	x, y int
}

type node struct {
	x, y, z int
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m, s int
	fmt.Fscan(in, &n, &m, &s)

	const N = 2e5 + 10
	a := make([][]pair, N)
	b := make([][]pair, N)
	var out [N][2]int
	for m > 0 {
		var u, v, w int
		fmt.Fscan(in, &u, &v, &w)
		a[u] = append(a[u], pair{v, w})
		b[v] = append(b[v], pair{u, w})
		out[u][0]++
		out[u][1]++
		m--
	}

	var f [N][2]int
	for i := 1; i <= n; i++ {
		f[i][0] = 9e18
	}

	q := &Heap{}
	heap.Init(q)
	for i := 1; i <= n; i++ {
		if out[i][0] == 0 {
			f[i][0] = 0
			f[i][1] = 0
			heap.Push(q, node{i, 0, 0})
			heap.Push(q, node{i, 1, 0})
		}
	}

	var vis [N][2]bool
	for q.Len() > 0 {
		xx := heap.Pop(q).(node)
		x := xx.x
		t := xx.y
		if vis[x][t] {
			continue
		}
		vis[x][t] = true
		for i, _ := range b[x] {
			v := b[x][i].x
			w := b[x][i].y
			out[v][t^1]--
			if t == 1 {
				f[v][0] = min(f[v][0], f[x][1]+w)
			} else if t == 0 {
				f[v][1] = max(f[v][1], f[x][0]+w)
			}
			if out[v][t^1] == 0 {
				heap.Push(q, node{v, t ^ 1, f[v][t^1]})
			} else if t == 1 {
				heap.Push(q, node{v, t ^ 1, f[v][t^1]})
			}
		}
	}

	if !vis[s][0] {
		fmt.Println("INFINITY")
	} else {
		fmt.Println(f[s][0])
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Heap []node

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i].z < h[j].z }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(node)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
