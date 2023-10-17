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

	a := make([]pair, 333333)
	ans := 0
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i].x, &a[i].y)
		a[i].x *= -1
		ans += a[i].y
	}
	tmp := a[1 : n+1]
	sortPair(tmp)
	pq := make(Heap, 0)
	heap.Init(&pq)
	for i := 1; i <= n; i++ {
		heap.Push(&pq, -a[i].y)
		heap.Push(&pq, -a[i].y)
		ans += heap.Pop(&pq).(int)
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
