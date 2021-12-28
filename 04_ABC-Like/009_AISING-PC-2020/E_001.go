package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

type pair struct{ x, y int }

func S(a []pair) int {
	sort.Slice(a, func(i, j int) bool {
		return a[i].x < a[j].x
	})
	A := 0
	pq := &Heap{}
	for _, p := range a {
		heap.Push(pq, p.y)
		A += p.y
		for pq.Len() > p.x {
			A -= heap.Pop(pq).(int)
		}
	}
	return A
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(in, &n)
		sum := 0
		L := []pair{}
		R := []pair{}
		for i := 0; i < n; i++ {
			var k, l, r int
			fmt.Fscan(in, &k, &l, &r)
			if l > r {
				L = append(L, pair{k, l - r})
				sum += r
			} else if r > l {
				R = append(R, pair{n - k, r - l})
				sum += l
			} else {
				sum += l
			}
		}
		fmt.Fprintln(out, sum+S(L)+S(R))
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
