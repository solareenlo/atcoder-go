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
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	const N = 100000

	var T int
	fmt.Fscan(in, &T)
	q := &Heap{}
	q1 := &Heap1{}
	var b [N + 5]int
	a := make([]int, N+5)
	for T > 0 {
		T--
		for q.Len() > 0 {
			heap.Pop(q)
		}
		for q1.Len() > 0 {
			heap.Pop(q1)
		}
		var n, m int
		fmt.Fscan(in, &n, &m)
		for i := 1; i <= n; i++ {
			fmt.Fscan(in, &a[i])
		}
		for i := 1; i <= m; i++ {
			fmt.Fscan(in, &b[i])
			heap.Push(q, b[i])
		}
		tmp := a[1 : n+1]
		sort.Ints(tmp)
		for i := n; i > 0; i-- {
			if q.Len() == 0 {
				t := heap.Pop(q1).(int)
				heap.Push(q, t/2)
				heap.Push(q1, t-t/2)
			}
			for (*q)[0] > a[i] {
				t := heap.Pop(q).(int)
				heap.Push(q, t/2)
				heap.Push(q, t-t/2)
				if q.Len()+q1.Len() > n {
					break
				}
			}
			heap.Push(q1, (*q)[0])
			heap.Pop(q)
		}
		if q.Len() != 0 {
			fmt.Fprintln(out, "No")
		} else {
			fmt.Fprintln(out, "Yes")
		}
	}
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

type Heap1 []int

func (h Heap1) Len() int            { return len(h) }
func (h Heap1) Less(i, j int) bool  { return h[i] < h[j] }
func (h Heap1) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap1) Push(x interface{}) { *h = append(*h, x.(int)) }

func (h *Heap1) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
