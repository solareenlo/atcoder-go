package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

type pair struct {
	x float64
	y int
}

type P struct {
	w    int
	l, r float64
}

func main() {
	in := bufio.NewReader(os.Stdin)

	const M = 10009

	var n, m int
	fmt.Fscan(in, &n, &m)
	var w [M]int
	var x, v [M]float64
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &w[i], &x[i], &v[i])
	}

	ans := 0
	var p [M]P
	for i := 1; i <= n; i++ {
		num, sum := 0, 0
		q := &Heap{}
		for j := 1; j <= n; j++ {
			var l, r float64
			if x[i] != x[j] && v[i] == v[j] {
				if x[j]-x[i] >= 0 && x[j]-x[i] <= float64(m) {
					sum += w[j]
				}
				continue
			}
			if x[i] == x[j] && v[i] == v[j] {
				l = 0
				r = 2e4
			} else {
				l = (x[j] - x[i]) / (v[i] - v[j])
				r = (x[j] - x[i] - float64(m)) / (v[i] - v[j])
			}
			if l > r {
				l, r = r, l
			}
			if r < 0 {
				continue
			}
			num++
			p[num] = P{w[j], l, r}
		}
		tmp := p[1 : num+1]
		sort.Slice(tmp, func(i, j int) bool {
			return tmp[i].r > tmp[j].r
		})
		for j := 1; j <= num; j++ {
			sum += p[j].w
			heap.Push(q, pair{p[j].l, p[j].w})
			for (*q)[0].x > p[j].r {
				sum -= (*q)[0].y
				heap.Pop(q)
			}
			if p[j].r >= 0 {
				ans = max(ans, sum)
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
