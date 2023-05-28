package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

var in = bufio.NewReader(os.Stdin)

type P struct {
	x, y int
}

func main() {

	var t int
	fmt.Fscan(in, &t)
	for t > 0 {
		t--
		solve()
	}
}

func solve() {
	var n int
	fmt.Fscan(in, &n)
	var s string
	fmt.Fscan(in, &s)
	st := make([]P, 0)
	sm := make([]int, 4*n+1)
	for i := 0; i < 2*n; i++ {
		var l, r int
		fmt.Fscan(in, &l, &r)
		l--
		r--
		if s[l] == s[r] {
			if s[l] == '(' {
				sm[l+1]++
			} else {
				sm[r+1]--
			}
		} else {
			if s[l] == '(' {
				sm[l+1]++
			} else {
				sm[r+1]++
			}
			st = append(st, P{l, r})
		}
	}
	for i := 0; i < 4*n; i++ {
		sm[i+1] += sm[i]
	}

	if sm[4*n] < 0 || sm[4*n]%2 == 1 {
		fmt.Println("No")
		return
	}
	x := sm[4*n] / 2
	if x > len(st) {
		fmt.Println("No")
		return
	}
	for i := 0; i <= 2*n; i++ {
		sm[i] -= x
	}
	for i := 2*n + 1; i <= 4*n; i++ {
		sm[i] -= 2 * x
	}
	sort.Slice(st, func(i, j int) bool {
		return st[i].y > st[j].y
	})
	id := 0
	pq := &Heap{}
	tmp := 0
	p := make([]int, 4*n+1)
	for i := 4 * n; i >= 2*n+1; i-- {
		for id < len(st) && st[id].y == i {
			heap.Push(pq, st[id].x)
			id++
		}
		for sm[i]+tmp < 0 {
			if pq.Len() == 0 {
				fmt.Println("No")
				return
			}
			L := heap.Pop(pq).(int)
			p[L]++
			tmp++
			x--
			if x < 0 {
				fmt.Println("No")
				return
			}
		}
		sm[i] += tmp
	}
	for id < len(st) {
		heap.Push(pq, st[id].x)
		id++
	}
	for x != 0 {
		L := heap.Pop(pq).(int)
		p[L]++
		tmp++
		x--
	}
	for i := 2 * n; i >= 0; i-- {
		p[i] += p[i+1]
	}
	for i := 0; i <= 4*n; i++ {
		if p[i]+sm[i] < 0 {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
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
