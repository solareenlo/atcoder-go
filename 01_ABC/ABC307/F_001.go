package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	const N = 300300

	type node struct {
		to, w, nxt int
	}
	var e [N << 1]node
	var head [N]int
	cnt := 0
	var add func(int, int, int)
	add = func(x, y, w int) {
		cnt++
		e[cnt] = node{y, w, head[x]}
		head[x] = cnt
	}

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 1; i <= m; i++ {
		var x, y, w int
		fmt.Fscan(in, &x, &y, &w)
		add(x, y, w)
		add(y, x, w)
	}
	var k int
	fmt.Fscan(in, &k)
	var dis [N]int
	for i := range dis {
		dis[i] = -1
	}
	p := make(HeapDot, 0)
	heap.Init(&p)
	for i := 1; i <= k; i++ {
		var x int
		fmt.Fscan(in, &x)
		dis[x] = 0
		for j := head[x]; j > 0; j = e[j].nxt {
			heap.Push(&p, dot{e[j].to, e[j].w})
		}
	}
	var d int
	fmt.Fscan(in, &d)
	q := make(HeapDot, 0)
	heap.Init(&q)
	for i := 1; i <= d; i++ {
		var now int
		fmt.Fscan(in, &now)
		for p.Len() > 0 && p[0].k <= now {
			if dis[p[0].x] == -1 {
				heap.Push(&q, p[0])
			}
			heap.Pop(&p)
		}
		for q.Len() > 0 {
			tmp := heap.Pop(&q).(dot)
			x := tmp.x
			k := tmp.k
			if dis[x] != -1 {
				continue
			}
			dis[x] = i
			for j := head[x]; j > 0; j = e[j].nxt {
				to := e[j].to
				if dis[to] != -1 {
					continue
				}
				if k+e[j].w <= now {
					heap.Push(&q, dot{to, k + e[j].w})
				} else {
					heap.Push(&p, dot{to, e[j].w})
				}
			}
		}
	}
	for i := 1; i <= n; i++ {
		fmt.Fprintln(out, dis[i])
	}
}

type dot struct {
	x, k int
}

type HeapDot []dot

func (h HeapDot) Len() int            { return len(h) }
func (h HeapDot) Less(i, j int) bool  { return h[i].k < h[j].k }
func (h HeapDot) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapDot) Push(x interface{}) { *h = append(*h, x.(dot)) }

func (h *HeapDot) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
