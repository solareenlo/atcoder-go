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

	var n, m int
	fmt.Fscan(in, &n, &m)
	p := make([]int, n)
	ans := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &p[i])
		ans += p[i]
	}
	sort.Ints(p)
	a := make([]pair, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &a[i].x)
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &a[i].y)
	}
	sortPair(a)
	pq := make(Heap, 0)
	heap.Init(&pq)
	j := 0
	for _, x := range p {
		for ; j < m && x >= a[j].x; j++ {
			heap.Push(&pq, a[j].y)
		}
		if pq.Len() > 0 {
			ans -= heap.Pop(&pq).(int)
		}
	}
	fmt.Println(ans)
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
