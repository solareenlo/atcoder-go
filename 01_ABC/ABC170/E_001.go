package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

const mx = 200005

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	a := [mx]int{}
	b := [mx]int{}
	for i := 1; i < n+1; i++ {
		fmt.Fscan(in, &a[i], &b[i])
	}

	c := [mx]int{}
	q := make([]Heap, mx)
	for i := 1; i < n+1; i++ {
		heap.Push(&q[b[i]], pai{a[i], i})
		c[b[i]] = max(c[b[i]], a[i])
	}

	Q := &Heap2{}
	for i := 1; i < 200001; i++ {
		if c[i] == 0 {
			continue
		}
		heap.Push(Q, pai{c[i], i})
	}

	for i := 1; i < m+1; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		if b[x] == y {
			continue
		}
		heap.Push(&q[y], pai{a[x], x})
		c[y] = max(c[y], a[x])
		heap.Push(Q, pai{c[y], y})
		z := b[x]
		b[x] = y
		for q[z].Len() > 0 && b[q[z][0].idx] != z {
			heap.Pop(&q[z])
		}
		if q[z].Len() > 0 {
			c[z] = q[z][0].a
		} else {
			c[z] = 0
		}
		if c[z] != 0 {
			heap.Push(Q, pai{c[z], z})
		}
		for Q.Len() > 0 && c[(*Q)[0].idx] != (*Q)[0].a {
			heap.Pop(Q)
		}
		fmt.Fprintln(out, (*Q)[0].a)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type pai struct{ a, idx int }

type Heap []pai

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i].a > h[j].a }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(pai)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type Heap2 []pai

func (h Heap2) Len() int            { return len(h) }
func (h Heap2) Less(i, j int) bool  { return h[i].a < h[j].a }
func (h Heap2) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap2) Push(x interface{}) { *h = append(*h, x.(pai)) }

func (h *Heap2) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
