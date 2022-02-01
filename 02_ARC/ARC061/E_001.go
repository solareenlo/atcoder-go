package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type edge struct{ t, c int }

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 100005
	const mx = 1000000

	var n, m int
	fmt.Fscan(in, &n, &m)

	w := make([][]edge, N)
	for i := 0; i < m; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		w[a] = append(w[a], edge{b, c})
		w[b] = append(w[b], edge{a, c})
	}

	d := make([]int, n+1)
	d[1] = 0
	for i := 2; i < n+1; i++ {
		d[i] = mx
	}

	q := &Heap{}
	heap.Push(q, pair{-d[1], 1})
	s := make([]map[int]bool, N)
	for i := range s {
		s[i] = make(map[int]bool, 1)
	}
	for q.Len() > 0 {
		u := (*q)[0].y
		cost := -(*q)[0].x
		heap.Pop(q)
		if d[u] < cost {
			continue
		}
		for _, j := range w[u] {
			dist := d[u]
			if _, ok := s[u][j.c]; !ok {
				dist++
			}
			if dist < d[j.t] {
				d[j.t] = dist
				for k := range s[j.t] {
					delete(s[j.t], k)
				}
				s[j.t][j.c] = true
				heap.Push(q, pair{-d[j.t], j.t})
			} else if dist == d[j.t] {
				s[j.t][j.c] = true
			}
		}
	}
	if d[n] >= mx {
		fmt.Println(-1)
	} else {
		fmt.Println(d[n])
	}
}

type pair struct{ x, y int }

type Heap []pair

func (h Heap) Len() int { return len(h) }
func (h Heap) Less(i, j int) bool {
	return (h[i].x > h[j].x) || (h[i].x == h[j].x && h[i].y > h[j].y)
}
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(pair)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
