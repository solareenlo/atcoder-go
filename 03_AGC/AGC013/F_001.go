package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

var (
	n int
	x = make([]int, n)
)

func down(x, y int) int {
	if x > y {
		x = y
	}
	return x
}

func pos(val int) int {
	return lowerBound(x, val)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	const N = 100100
	const INF = 1 << 60
	fmt.Fscan(in, &n)
	n++

	a := make([]pair, n)
	for i := 1; i < n; i++ {
		fmt.Fscan(in, &a[i].x, &a[i].y)
	}

	x = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x[i])
	}
	sort.Ints(x)

	for i := 1; i < n; i++ {
		a[i].y = pos(a[i].y)
		a[i].x = pos(a[i].x)
		a[i].y = down(a[i].y, a[i].x)
	}

	c := make([]int, N)
	for i := range c {
		c[i] = -1
	}

	for i := 1; i < n; i++ {
		c[a[i].x]++
		a[i].x, a[i].y = a[i].y, a[i].x
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i].x < a[j].x || (a[i].x == a[j].x && a[i].y < a[j].y)
	})

	cur := 0
	cov := 0
	q := &Heap{}
	C := make([]int, N)
	last := make([]int, N)
	flag := false
	for j, i := 1, 0; i < n; i++ {
		cur += c[i]
		for ; j < n && a[j].x == i; j++ {
			heap.Push(q, (a[j].y))
		}
		r := INF
		for q.Len() > 0 && (*q)[0] > i && cur < 0 {
			r = (*q)[0]
			heap.Pop(q)
			cur++
			cov++
			c[r]--
		}
		for q.Len() > 0 && (*q)[0] <= i {
			heap.Pop(q)
		}
		C[i] = cur
		last[i] = r
		if cur < -1 {
			flag = true
			break
		}
	}

	var Q int
	fmt.Fscan(in, &Q)
	if flag {
		for ; Q > 0; Q-- {
			fmt.Fprintln(out, -1)
		}
		return
	}

	ans := make([]int, N)
	var j int
	for r, i := n-1, n-1; i >= 0; i-- {
		if C[i] >= 0 {
			if last[i] <= r {
				r = i
				cov--
			}
		} else {
			j = i
			r = j
		}
		ans[i] = cov
	}
	for i := j + 1; i < N; i++ {
		ans[i] = 1 << 60
	}

	for i := 0; i < Q; i++ {
		var d, e, MIN int
		fmt.Fscan(in, &d, &e)
		MIN = min(ans[pos(d)], ans[pos(e)]+1)
		if MIN >= INF {
			fmt.Fprintln(out, -1)
		} else {
			fmt.Fprintln(out, n-MIN)
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

type pair struct{ x, y int }
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
