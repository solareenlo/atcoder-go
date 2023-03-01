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

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)

	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	t := make([]pair, m+1)
	for i := 1; i <= m; i++ {
		t[i] = pair{a[i], i}
	}
	tmp := t[1:]
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].x == tmp[j].x {
			return tmp[i].y < tmp[j].y
		}
		return tmp[i].x < tmp[j].x
	})

	q := &HeapQ{}
	sum := 0
	for i := 1; i <= k; i++ {
		heap.Push(q, t[i])
		sum += t[i].x
	}
	p := &HeapP{}
	for i := k + 1; i <= m; i++ {
		heap.Push(p, t[i])
	}
	fmt.Printf("%d ", sum)

	var lessThan func(pair, pair) bool
	lessThan = func(a, b pair) bool {
		if a.x == b.x {
			return a.y <= b.y
		}
		return a.x <= b.x
	}
	s := k
	u := n - k
	for i := m + 1; i <= n; i++ {
		j := i - m
		b := pair{a[j], j}
		if lessThan(b, (*q)[0]) {
			s--
			sum -= a[j]
		} else {
			u--
		}
		b = pair{a[i], i}
		if !lessThan(b, (*q)[0]) {
			u++
			heap.Push(p, b)
		} else {
			s++
			heap.Push(q, b)
			sum += a[i]
		}
		for s > k {
			x := heap.Pop(q).(pair)
			if x.y <= j {
				continue
			} else {
				heap.Push(p, x)
				s--
				u++
				sum -= x.x
			}
		}
		for s < k {
			x := heap.Pop(p).(pair)
			if x.y <= j {
				continue
			} else {
				heap.Push(q, x)
				u--
				s++
				sum += x.x
			}
		}
		fmt.Printf("%d ", sum)
	}
}

type pair struct {
	x, y int
}

type HeapQ []pair

func (h HeapQ) Len() int            { return len(h) }
func (h HeapQ) Less(i, j int) bool  { return h[i].x > h[j].x }
func (h HeapQ) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapQ) Push(x interface{}) { *h = append(*h, x.(pair)) }

func (h *HeapQ) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type HeapP []pair

func (h HeapP) Len() int            { return len(h) }
func (h HeapP) Less(i, j int) bool  { return h[i].x < h[j].x }
func (h HeapP) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapP) Push(x interface{}) { *h = append(*h, x.(pair)) }

func (h *HeapP) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
