package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

type pair struct {
	x, y int
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	v := make([]pair, 0)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		v = append(v, pair{b, a})
	}
	sort.Slice(v, func(i, j int) bool {
		if v[i].x == v[j].x {
			return v[i].y < v[j].y
		}
		return v[i].x < v[j].x
	})

	tmp := 0
	p := &Heap{}
	heap.Init(p)
	for i := 0; i < n; i++ {
		tmp += v[i].y
		heap.Push(p, v[i].y)
		if tmp > v[i].x {
			tmp = tmp - heap.Pop(p).(int)
		}
	}
	fmt.Println(p.Len())
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
