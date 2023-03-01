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

	const MX = 100005
	const INF = 1000000005

	type pair struct {
		x, y int
	}

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)

	p := make([]pair, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &p[i].x, &p[i].y)
	}
	sort.Slice(p, func(i, j int) bool {
		if p[i].x == p[j].x {
			return p[i].y < p[j].y
		}
		return p[i].x < p[j].x
	})

	l := 0
	r := INF
	var bfs func(int) bool
	bfs = func(c int) bool {
		x, h, pi, last := 0, 0, 0, 0
		q := &Heap{}
		for x+c <= m {
			for pi < n && p[pi].x <= x+c {
				heap.Push(q, -p[pi].y)
				pi++
			}
			if q.Len() > 0 {
				nx := -heap.Pop(q).(int)
				last = max(last, nx)
				if nx > x {
					x = min(x+c, nx)
				}
			} else {
				if x == last {
					return false
				}
				x = last
				h++
			}
		}
		return (h <= k)
	}
	for l+1 < r {
		c := (l + r) >> 1
		if bfs(c) {
			r = c
		} else {
			l = c
		}
	}
	fmt.Println(r)
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
