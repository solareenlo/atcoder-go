package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

type edge struct {
	nxt, to, val int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	N := 200005
	e := make([]edge, N*2)
	tot := 1
	fst := make([]int, N)
	for i := 0; i < m; i++ {
		var u, v, w int
		fmt.Fscan(in, &u, &v, &w)
		tot++
		e[tot].nxt = fst[u]
		e[tot].to = v
		e[tot].val = w
		fst[u] = tot
		tot++
		e[tot].nxt = fst[v]
		e[tot].to = u
		e[tot].val = w
		fst[v] = tot
	}

	dis := make([]int, N)
	for i, _ := range dis {
		dis[i] = math.MaxInt64
	}
	dis[1] = 0

	p := make([]int, N)
	que := &Heap{}
	heap.Init(que)
	heap.Push(que, 1)
	for que.Len() > 0 {
		u := heap.Pop(que).(int)
		for i := fst[u]; i > 0; i = e[i].nxt {
			v := e[i].to
			if dis[v] > dis[u]+e[i].val {
				p[v] = i >> 1
				dis[v] = dis[u] + e[i].val
				heap.Push(que, v)
			}
		}
	}

	for i := 2; i <= n; i++ {
		fmt.Fprintf(out, "%d ", p[i])
	}
}

type Heap []int

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i] > h[j] }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(int)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
