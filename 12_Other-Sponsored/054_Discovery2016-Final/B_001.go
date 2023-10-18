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

	type pair struct {
		x, y int
	}

	var n, x int
	fmt.Fscan(in, &n, &x)
	a := make([]pair, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i].x)
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i].y)
	}
	sort.Slice(a, func(i, j int) bool {
		if a[i].x == a[j].x {
			return a[i].y < a[j].y
		}
		return a[i].x < a[j].x
	})
	pq := make(Heap, 0)
	heap.Init(&pq)
	for i := 0; i < n; i++ {
		heap.Push(&pq, a[i].y)
		for len(pq) > a[i].x {
			heap.Pop(&pq)
		}
	}
	vec := make([]int, 0)
	for pq.Len() > 0 {
		vec = append(vec, heap.Pop(&pq).(int))
	}
	reverseOrderInt(vec)
	for i := 1; i <= len(vec); i++ {
		x -= vec[i-1]
		if x <= 0 {
			fmt.Println(i)
			return
		}
	}
	fmt.Println(-1)
}

func reverseOrderInt(a []int) {
	n := len(a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
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
