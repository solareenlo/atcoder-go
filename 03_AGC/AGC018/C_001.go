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

	var xx, yy, zz int
	fmt.Fscan(in, &xx, &yy, &zz)

	n := xx + yy + zz
	sum := 0
	type pair struct{ x, y int }
	a := make([]pair, n+1)
	for i := 1; i <= n; i++ {
		var x, y, z int
		fmt.Fscan(in, &x, &y, &z)
		sum += x
		a[i] = pair{y - x, z - x}
	}

	tmp := make([]pair, n)
	copy(tmp, a[1:])
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].x-tmp[i].y > tmp[j].x-tmp[j].y
	})
	copy(a[1:], tmp)

	q := &Heap{}
	suf := make([]int, 100100)
	for i := 1; i <= n; i++ {
		heap.Push(q, -a[i].x)
		suf[i] = suf[i-1] + a[i].x
		if q.Len() > yy {
			suf[i] += (*q)[0]
			heap.Pop(q)
		}
	}

	for q.Len() != 0 {
		heap.Pop(q)
	}

	ans := 0
	for i := n; i > yy; i-- {
		heap.Push(q, -a[i].y)
		sum += a[i].y
		if q.Len() > zz {
			sum += (*q)[0]
			heap.Pop(q)
		}
		if i+zz-1 <= n {
			ans = max(ans, sum+suf[i-1])
		}
	}

	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
