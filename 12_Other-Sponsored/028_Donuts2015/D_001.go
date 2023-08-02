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

	var n, k int
	fmt.Fscan(in, &n, &k)
	c := make([]int, n)
	for i := range c {
		fmt.Fscan(in, &c[i])
	}
	ord := make([]int, n)
	for i := range ord {
		ord[i] = i
	}
	sort.Slice(ord, func(i, j int) bool {
		return c[ord[i]] < c[ord[j]]
	})

	pq := &HeapTuple{}
	rpq := &HeapTupleR{}
	li := make([]int, n)
	ri := make([]int, n)

	for i := 0; i < n; i++ {
		if i == 0 {
			li[ord[i]] = -1
		} else {
			li[ord[i]] = ord[i-1]
		}
		if i+1 == n {
			ri[ord[i]] = -1
		} else {
			ri[ord[i]] = ord[i+1]
		}
	}
	for i := 1; i < n; i++ {
		i1 := ord[i-1]
		i2 := ord[i]
		heap.Push(pq, tuple{c[i2] - c[i1], i1, i2})
	}

	ans := 0
	inside := make([]int, n)
	outside := make([]int, n)
	for i := 0; i < n; i++ {
		inside[i] = -1
		outside[i] = -1
	}

	var set_io func()
	set_io = func() {
		for pq.Len() > 0 {
			tmp := heap.Pop(pq).(tuple)
			c := tmp.x
			a := tmp.y
			b := tmp.z
			if ri[a] != b || li[b] != a {
				continue
			}
			inside[b] = a
			outside[a] = b
			heap.Push(rpq, tuple{c, a, b})
			ans += c
			break
		}
	}
	for i := 0; i < n-k; i++ {
		set_io()
	}

	fmt.Println(ans)

	var q int
	fmt.Fscan(in, &q)
	for q > 0 {
		q--
		var idx int
		fmt.Fscan(in, &idx)
		idx--
		l := li[idx]
		r := ri[idx]
		if l != -1 {
			ri[l] = ri[idx]
		}
		if r != -1 {
			li[r] = li[idx]
		}
		if l != -1 && r != -1 {
			heap.Push(pq, tuple{c[r] - c[l], l, r})
		}
		li[idx] = -1
		ri[idx] = -1
		if inside[idx] == -1 && outside[idx] == -1 {
			for rpq.Len() > 0 {
				tmp := heap.Pop(rpq).(tuple)
				c := tmp.x
				a := tmp.y
				b := tmp.z
				if inside[b] != a {
					continue
				}
				ans -= c
				inside[b] = -1
				outside[a] = -1
				heap.Push(pq, tuple{c, a, b})
				break
			}
		} else if inside[idx] == -1 {
			inside[outside[idx]] = -1
			ans -= c[outside[idx]] - c[idx]
		} else if outside[idx] == -1 {
			outside[inside[idx]] = -1
			ans -= c[idx] - c[inside[idx]]
		} else {
			a := inside[idx]
			b := outside[idx]
			inside[b] = -1
			outside[a] = -1
			ans -= c[b] - c[a]
			set_io()
		}
		inside[idx] = -1
		outside[idx] = -1
		fmt.Println(ans)
	}
}

type tuple struct {
	x, y, z int
}

type HeapTuple []tuple

func (h HeapTuple) Len() int            { return len(h) }
func (h HeapTuple) Less(i, j int) bool  { return h[i].x < h[j].x }
func (h HeapTuple) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapTuple) Push(x interface{}) { *h = append(*h, x.(tuple)) }

func (h *HeapTuple) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type HeapTupleR []tuple

func (h HeapTupleR) Len() int            { return len(h) }
func (h HeapTupleR) Less(i, j int) bool  { return h[i].x > h[j].x }
func (h HeapTupleR) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapTupleR) Push(x interface{}) { *h = append(*h, x.(tuple)) }

func (h *HeapTupleR) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
