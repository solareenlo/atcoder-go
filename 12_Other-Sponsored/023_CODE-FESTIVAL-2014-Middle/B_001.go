package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MAX = 100000

	var n, m int
	fmt.Fscan(in, &n, &m)
	d := make([][]int, MAX+1)
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		d[x] = append(d[x], y)
	}
	cnt := make([]int, MAX+1)
	for i := 0; i < m; i++ {
		var a int
		fmt.Fscan(in, &a)
		cnt[a]++
	}
	res := 0
	pq := &Heap{}
	for i := 1; i <= MAX; i++ {
		for _, p := range d[i] {
			heap.Push(pq, p)
		}
		for pq.Len() != 0 && (*pq)[0] < i {
			heap.Pop(pq)
		}
		for j := 0; j < cnt[i]; j++ {
			if pq.Len() != 0 {
				res++
				heap.Pop(pq)
			}
		}
	}
	fmt.Println(res)
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
