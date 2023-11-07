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

	var a [1001000]int

	var n, k, Q int
	fmt.Fscan(in, &n, &k, &Q)
	sum := 0
	if k == n {
		for Q > 0 {
			Q--
			var x, y int
			fmt.Fscan(in, &x, &y)
			sum -= a[x]
			a[x] = y
			sum += a[x]
			fmt.Fprintln(out, sum)
		}
		return
	}
	q := make(HeapL, 0)
	qq := make(HeapL, 0)
	p := make(HeapG, 0)
	pp := make(HeapG, 0)
	for i := 1; i <= k; i++ {
		heap.Push(&p, 0)
	}
	for i := 1; i <= n-k; i++ {
		heap.Push(&q, 0)
	}
	for Q > 0 {
		Q--
		var x, y int
		fmt.Fscan(in, &x, &y)
		if a[x] <= q[0] {
			heap.Push(&qq, a[x])
			for qq.Len() != 0 && q[0] == qq[0] {
				heap.Pop(&q)
				heap.Pop(&qq)
			}
		} else {
			sum -= a[x]
			heap.Push(&pp, a[x])
			for pp.Len() != 0 && p[0] == pp[0] {
				heap.Pop(&p)
				heap.Pop(&pp)
			}
			sum += q[0]
			heap.Push(&p, heap.Pop(&q).(int))
			for qq.Len() != 0 && q[0] == qq[0] {
				heap.Pop(&q)
				heap.Pop(&qq)
			}
		}
		a[x] = y
		if y <= p[0] {
			heap.Push(&q, y)
		} else {
			heap.Push(&p, y)
			sum += y - p[0]
			heap.Push(&q, heap.Pop(&p).(int))
			for pp.Len() != 0 && p[0] == pp[0] {
				heap.Pop(&p)
				heap.Pop(&pp)
			}
		}
		fmt.Fprintln(out, sum)
	}
}

type HeapG []int

func (h HeapG) Len() int            { return len(h) }
func (h HeapG) Less(i, j int) bool  { return h[i] < h[j] }
func (h HeapG) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapG) Push(x interface{}) { *h = append(*h, x.(int)) }

func (h *HeapG) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type HeapL []int

func (h HeapL) Len() int            { return len(h) }
func (h HeapL) Less(i, j int) bool  { return h[i] > h[j] }
func (h HeapL) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapL) Push(x interface{}) { *h = append(*h, x.(int)) }

func (h *HeapL) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
