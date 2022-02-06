package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	v := make([][]int, 200200)
	for i := n; i > 0; i-- {
		var x, y int
		fmt.Fscan(in, &x, &y)
		v[x] = append(v[x], y)
	}

	q := &Heap{}
	t := 0
	for i := m; i >= 0; i-- {
		for _, o := range v[i] {
			heap.Push(q, o)
		}
		if i != 0 && q.Len() != 0 {
			heap.Pop(q)
			t++
		}
	}
	ans := t
	y := m
	for q.Len() > 0 {
		x := (*q)[0]
		heap.Pop(q)
		if x <= y && y > t {
			ans++
			y--
		}
	}
	fmt.Println(n - ans)
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
