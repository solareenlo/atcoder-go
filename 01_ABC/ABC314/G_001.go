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

	const N = 300005

	var d, dis [N]int
	var vis [N]bool

	var n, m, h int
	fmt.Fscan(in, &n, &m, &h)
	q1 := make(HeapNode1, 0)
	heap.Init(&q1)
	q2 := make(HeapNode2, 0)
	heap.Init(&q2)
	for i := 1; i <= n; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		d[i] = d[i-1]
		dis[b] += a
		if vis[b] {
			heap.Push(&q2, node2{dis[b], b})
			continue
		}
		h -= a
		for q2.Len() != 0 && dis[q2[0].id] != q2[0].d {
			heap.Pop(&q2)
		}
		if q2.Len() != 0 && dis[b] > q2[0].d {
			heap.Push(&q1, node1{q2[0].d, q2[0].id})
			vis[q2[0].id] = false
			h -= q2[0].d
			heap.Pop(&q2)
			heap.Push(&q2, node2{dis[b], b})
			vis[b] = true
			h += dis[b]
		} else {
			heap.Push(&q1, node1{dis[b], b})
		}
		for h <= 0 {
			for vis[q1[0].id] || dis[q1[0].id] != q1[0].d {
				heap.Pop(&q1)
			}
			h += q1[0].d
			vis[q1[0].id] = true
			heap.Push(&q2, node2{q1[0].d, q1[0].id})
			d[i]++
			heap.Pop(&q1)
		}
	}
	for i := 0; i < m+1; i++ {
		fmt.Fprintf(out, "%d ", upperBound(d[1:n+1], i))
	}
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}

type node1 struct {
	d, id int
}

type HeapNode1 []node1

func (h HeapNode1) Len() int            { return len(h) }
func (h HeapNode1) Less(i, j int) bool  { return h[i].d > h[j].d }
func (h HeapNode1) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapNode1) Push(x interface{}) { *h = append(*h, x.(node1)) }

func (h *HeapNode1) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type node2 struct {
	d, id int
}

type HeapNode2 []node2

func (h HeapNode2) Len() int            { return len(h) }
func (h HeapNode2) Less(i, j int) bool  { return h[i].d < h[j].d }
func (h HeapNode2) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapNode2) Push(x interface{}) { *h = append(*h, x.(node2)) }

func (h *HeapNode2) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
