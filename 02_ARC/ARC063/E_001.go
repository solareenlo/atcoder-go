package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	const N = 233333
	e := make([][]int, N)
	for i := 1; i < n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		e[x] = append(e[x], y)
		e[y] = append(e[y], x)
	}

	a := make([]int, N)
	for i := range a {
		a[i] = -1
	}

	var k int
	fmt.Fscan(in, &k)
	q := &Heap{}
	for i := 0; i < k; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		a[x] = y
		heap.Push(q, node{a[x], x})
	}

	for q.Len() > 0 {
		x := (*q)[0].y
		y := (*q)[0].x
		heap.Pop(q)
		for _, i := range e[x] {
			if a[i] < 0 {
				a[i] = y + 1
				heap.Push(q, node{a[i], i})
			} else if a[i] != y+1 && a[i] != y-1 {
				fmt.Fprintln(out, "No")
				return
			}
		}
	}

	fmt.Fprintln(out, "Yes")
	for i := 1; i < n+1; i++ {
		fmt.Fprintln(out, a[i])
	}
}

type node struct{ x, y int }

type Heap []node

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i].x < h[j].x }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(node)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
