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

	var n, m int
	fmt.Fscan(in, &n, &m)

	ga := make([][]int, n)
	gb := make([][]int, n)
	for i := 0; i < m; i++ {
		var t, u, v int
		fmt.Fscan(in, &t, &u, &v)
		if t != 0 {
			gb[u] = append(gb[u], v)
			gb[v] = append(gb[v], u)
		} else {
			ga[u] = append(ga[u], v)
			ga[v] = append(ga[v], u)
		}
	}

	d := make([]int, n)
	for i := range d {
		d[i] = 1 << 60
	}

	q := &Heap{}
	heap.Push(q, T{0, 0, 0})
	for q.Len() > 0 {
		bn := (*q)[0].bn
		cost := (*q)[0].cost
		v := (*q)[0].v
		heap.Pop(q)
		if cost >= d[v] {
			continue
		}
		d[v] = cost
		for _, to := range ga[v] {
			heap.Push(q, T{bn, cost + 1, to})
		}
		for _, to := range gb[v] {
			heap.Push(q, T{bn + 1, cost + bn + 1, to})
		}
	}
	for i := 0; i < n; i++ {
		fmt.Fprintln(out, d[i])
	}
}

type T struct{ bn, cost, v int }
type Heap []T

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i].bn < h[j].bn }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(T)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
