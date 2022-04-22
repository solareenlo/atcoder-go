package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	type Point struct {
		x, y float64
		c    int
	}
	p := make([]Point, 40)
	for i := 0; i < n+m; i++ {
		fmt.Fscan(in, &p[i].x, &p[i].y, &p[i].c)
	}

	ans := 1e9
	for bt := 0; bt < 1<<m; bt++ {
		used := make([]bool, 40)
		for i := 0; i < m; i++ {
			if (bt>>i)&1 != 0 {
				used[i+n] = true
			}
		}
		ttcost := 0.0
		qu := &Heap{}
		heap.Push(qu, P{0, 0})
		for qu.Len() > 0 {
			cost := (*qu)[0].x
			t := (*qu)[0].y
			heap.Pop(qu)
			if used[t] {
				continue
			}
			ttcost += float64(cost)
			used[t] = true
			for i := 0; i < n+m; i++ {
				if !used[i] {
					tmp := 10.0
					if p[t].c == p[i].c {
						tmp = 1.0
					}
					heap.Push(qu, P{math.Hypot(p[t].x-p[i].x, p[t].y-p[i].y) * tmp, i})
				}
			}
		}
		ans = min(ans, ttcost)
	}
	fmt.Println(ans)
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

type P struct {
	x float64
	y int
}
type Heap []P

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i].x < h[j].x }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(P)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
