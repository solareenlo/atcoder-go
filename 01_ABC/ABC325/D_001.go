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

	const N = 200005

	var n int
	fmt.Fscan(in, &n)
	a := make([]pair, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i].x, &a[i].y)
		a[i].y += a[i].x
	}
	sortPair(a[1:])
	p := 1
	k := 0
	q := &Heap{}
	heap.Init(q)
	ans := 0
	for p <= n || q.Len() > 0 {
		if q.Len() == 0 {
			k = max(k, a[p].x)
		}
		for p <= n && a[p].x == k {
			heap.Push(q, -a[p].y)
			p++
		}
		for q.Len() > 0 {
			x := -heap.Pop(q).(int)
			if k <= x {
				k++
				ans++
				break
			}
		}
	}
	fmt.Println(ans)
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

type pair struct {
	x, y int
}

func sortPair(tmp []pair) {
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].x == tmp[j].x {
			return tmp[i].y < tmp[j].y
		}
		return tmp[i].x < tmp[j].x
	})
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
