package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type Sushi struct {
	t, d int
}

func main() {
	var N, k int
	fmt.Scan(&N, &k)

	sushi := make([]Sushi, N)
	for i := range sushi {
		fmt.Scan(&sushi[i].t, &sushi[i].d)
	}
	sort.Slice(sushi, func(i, j int) bool {
		return sushi[i].d > sushi[j].d
	})

	u := make([]bool, N+1)
	q := &Heap{}
	sum, n := 0, 0
	for i := 0; i < k; i++ {
		sum += sushi[i].d
		if u[sushi[i].t] {
			heap.Push(q, sushi[i].d)
		} else {
			u[sushi[i].t] = true
			n++
		}
	}

	point := sum + n*n
	for i := k; i < N; i++ {
		if !u[sushi[i].t] {
			if q.Len() == 0 {
				break
			}
			u[sushi[i].t] = true
			sum += sushi[i].d
			sum -= heap.Pop(q).(int)
			n++
			point = max(point, sum+n*n)
		}
	}
	fmt.Println(point)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Heap []int

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i] < h[j] }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(int)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
