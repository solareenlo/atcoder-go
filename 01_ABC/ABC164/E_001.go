package main

import (
	"container/heap"
	"fmt"
)

func main() {
	var n, m, s int
	fmt.Scan(&n, &m, &s)

	type E struct{ to, coin, time int }

	edge := make([][]E, 50)
	for i := 0; i < m; i++ {
		var u, v, a, b int
		fmt.Scan(&u, &v, &a, &b)
		u--
		v--
		edge[u] = append(edge[u], E{v, a, b})
		edge[v] = append(edge[v], E{u, a, b})
	}

	c := make([]int, n)
	d := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&c[i], &d[i])
	}

	dp := [50][5001]int{}
	res := [50]int{}
	que := &Heap{}
	heap.Push(que, P{1, 0, min(5000, s)})
	for que.Len() > 0 {
		p := (*que)[0]
		heap.Pop(que)
		if dp[p.pos][p.re] != 0 {
			continue
		}
		dp[p.pos][p.re] = p.time
		if res[p.pos] == 0 {
			res[p.pos] = p.time
		}
		for _, q := range edge[p.pos] {
			if p.re >= q.coin && dp[q.to][p.re-q.coin] == 0 {
				heap.Push(que, P{p.time + q.time, q.to, p.re - q.coin})
			}
		}
		if p.re+c[p.pos] <= 5000 {
			heap.Push(que, P{p.time + d[p.pos], p.pos, p.re + c[p.pos]})
		}
	}
	for i := 1; i < n; i++ {
		fmt.Println(res[i] - 1)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type P struct{ time, pos, re int }
type Heap []P

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i].time < h[j].time }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(P)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
