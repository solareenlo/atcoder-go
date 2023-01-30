package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

const N = 200005

type pair struct {
	x, y int
}

type edge struct {
	next, to int
}

var num int
var e []edge
var head [N]int

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	var a [N]int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	e = make([]edge, N<<1)
	var val [N]int
	for i := 1; i <= m; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		add_edge(u, v)
		add_edge(v, u)
		val[u] += a[v]
		val[v] += a[u]
	}
	var vis [N]bool
	q := &Heap{}
	for i := 1; i <= n; i++ {
		vis[i] = false
		heap.Push(q, pair{-val[i], i})
	}
	ans := 0
	for q.Len() > 0 {
		x := heap.Pop(q).(pair)
		u := x.y
		if vis[u] {
			continue
		}
		vis[u] = true
		ans = max(ans, -x.x)
		for i := head[u]; i > 0; i = e[i].next {
			v := e[i].to
			if !vis[v] {
				val[v] -= a[u]
				heap.Push(q, pair{-val[v], v})
			}
		}
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func add_edge(from, to int) {
	num++
	e[num].next = head[from]
	e[num].to = to
	head[from] = num
}

type Heap []pair

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i].x > h[j].x }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(pair)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
