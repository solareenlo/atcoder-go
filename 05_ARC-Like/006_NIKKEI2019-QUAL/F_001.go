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

	var n, m int
	fmt.Fscan(in, &n, &m)

	a := make([][]int, 200005)
	for i := 1; i <= n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		a[x] = append(a[x], y)
	}

	p := make([]*HeapPair, 3)
	for i := range p {
		p[i] = &HeapPair{}
	}
	for i := 1; i <= m; i++ {
		sort.Sort(sort.Reverse(sort.IntSlice(a[i])))
		heap.Push(p[0], pair{a[i][0] + a[i][1], i})
		if len(a[i]) > 2 {
			heap.Push(p[2], pair{a[i][0] + a[i][1] + a[i][2], i})
		}
	}

	q := make([]*HeapInt, 2)
	for i := range q {
		q[i] = &HeapInt{}
	}
	ans := 0
	vis := make([]bool, n)
	for x, sz := 1, 0; x <= n; x++ {
		if sz == x {
			fmt.Fprintln(out, ans)
			continue
		}
		v0, c0, v1 := 0, 0, 0
		if p[0].Len() != 0 {
			v0 = (*p[0])[0].x
			c0 = (*p[0])[0].y
		}
		if p[1].Len() != 0 {
			v1 = (*p[1])[0].x
		}
		if v0 > v1+v1 {
			if x-sz < 2 {
				tmp := -1 << 60
				if p[1].Len() != 0 {
					tmp = (*p[1])[0].x
				}
				if p[0].Len() != 0 && q[1].Len() != 0 {
					tmp = max(tmp, (*p[0])[0].x+(*q[1])[0])
				}
				for p[2].Len() != 0 && vis[(*p[2])[0].y] {
					heap.Pop(p[2])
				}
				if p[2].Len() != 0 && q[0].Len() != 0 {
					tmp = max(tmp, (*p[2])[0].x+(*q[0])[0])
				}
				if tmp == -1<<60 {
					fmt.Fprintln(out, -1)
				} else {
					fmt.Fprintln(out, ans+tmp)
				}
			}
			vis[c0] = true
			heap.Pop(p[0])
			heap.Push(q[0], -v0)
			for i := 2; i < len(a[c0]); i++ {
				heap.Push(p[1], pair{a[c0][i], c0})
			}
			ans += v0
			sz += 2
		} else {
			heap.Pop(p[1])
			heap.Push(q[1], -v1)
			ans += v1
			sz++
			fmt.Fprintln(out, ans)
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type HeapInt []int

func (h HeapInt) Len() int            { return len(h) }
func (h HeapInt) Less(i, j int) bool  { return h[i] > h[j] }
func (h HeapInt) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapInt) Push(x interface{}) { *h = append(*h, x.(int)) }

func (h *HeapInt) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type pair struct{ x, y int }
type HeapPair []pair

func (h HeapPair) Len() int            { return len(h) }
func (h HeapPair) Less(i, j int) bool  { return h[i].x > h[j].x }
func (h HeapPair) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapPair) Push(x interface{}) { *h = append(*h, x.(pair)) }

func (h *HeapPair) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
