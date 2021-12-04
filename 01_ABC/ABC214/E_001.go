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

	var t int
	fmt.Fscan(in, &t)

	type pair struct{ l, r int }
	for k := 0; k < t; k++ {
		var n int
		fmt.Fscan(in, &n)
		a := make([]pair, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a[i].l, &a[i].r)
		}
		a = append(a, pair{1000000007, 1000000007})
		sort.Slice(a, func(i, j int) bool {
			return a[i].l < a[j].l || (a[i].l == a[j].l && a[i].r < a[j].r)
		})
		q := &Heap{}
		nw := 1
		ok := true
		for i := 0; i < n+1; i++ {
			for nw < a[i].l && q.Len() > 0 {
				if -(*q)[0] < nw {
					ok = false
				}
				heap.Pop(q)
				nw++
			}
			nw = a[i].l
			heap.Push(q, -a[i].r)
		}
		if ok {
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
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
