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

	const MX = 20005

	type pair struct {
		x, y int
	}

	var n int
	fmt.Fscan(in, &n)
	var t, a, b [MX]int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &t[i], &a[i], &b[i])
	}
	l, r := -int(1e9), int(1e9)
	for l+1 < r {
		ok := true
		m := (l + r) / 2
		p := make([]pair, n)
		qu := &Heap{}
		for i := 0; i < n; i++ {
			k := (a[i] - m) / b[i]
			if a[i]-m < 0 {
				ok = false
			}
			p[i] = pair{t[i] - k, t[i] + k}
		}
		sort.Slice(p, func(i, j int) bool {
			if p[i].x == p[j].x {
				return p[i].y < p[j].y
			}
			return p[i].x < p[j].x
		})
		for i, c := 1, 0; i <= n; i++ {
			for c < n && p[c].x <= i {
				heap.Push(qu, p[c].y)
				c++
			}
			if qu.Len() == 0 || (*qu)[0] < i {
				ok = false
			} else {
				heap.Pop(qu)
			}
		}
		if ok {
			l = m
		} else {
			r = m
		}
	}
	fmt.Println(l)
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
