package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

const INF = 1 << 30
const MAXN = 100000

var G [MAXN][]E

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, D int
	fmt.Fscan(in, &N, &D)
	for i := 0; i < N-1; i++ {
		var A, B, C int
		fmt.Fscan(in, &A, &B, &C)
		A--
		B--
		G[A] = append(G[A], E{B, C, 1})
		G[B] = append(G[B], E{A, C, -1})
	}

	d := make([]bool, MAXN)
	for i := range d {
		d[i] = true
	}
	q := &Heap{}
	var bfs func(int, []bool) int
	bfs = func(s int, d []bool) int {
		_max := -1
		idx := -1
		heap.Push(q, F{s, 0, -1})
		for q.Len() > 0 {
			tmp := heap.Pop(q).(F)
			to := tmp.a
			co := tmp.b
			pre := tmp.c
			t := false
			if co <= D {
				t = true
			}
			if d[to] && t {
				d[to] = true
			} else {
				d[to] = false
			}
			if co > _max {
				idx = to
			}
			for _, e := range G[to] {
				if e.to != pre {
					heap.Push(q, F{e.to, co + e.co, to})
				}
			}
		}
		return idx
	}
	bfs(bfs(bfs(0, d), d), d)

	cnt := 0
	ans := INF
	heap.Push(q, F{0, 0, -1})
	for q.Len() > 0 {
		tmp := heap.Pop(q).(F)
		to := tmp.a
		di := tmp.b
		pre := tmp.c
		if di == 1 {
			cnt++
		}
		for _, e := range G[to] {
			if e.to != pre {
				heap.Push(q, F{e.to, e.di, to})
			}
		}
	}
	heap.Push(q, F{0, cnt, -1})
	for q.Len() > 0 {
		tmp := heap.Pop(q).(F)
		to := tmp.a
		ti := tmp.b
		pre := tmp.c
		if d[to] && ans > ti {
			ans = ti
		}
		for _, e := range G[to] {
			if e.to != pre {
				heap.Push(q, F{e.to, ti - e.di, to})
			}
		}
	}
	if ans == INF {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}
}

type E struct {
	to, co, di int
}

type F struct {
	a, b, c int
}

type Heap []F

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i].a < h[j].a }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(F)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
