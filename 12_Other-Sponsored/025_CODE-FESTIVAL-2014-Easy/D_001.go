package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	mx := 100003
	bucket := make([][]int, mx)
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		bucket[x] = append(bucket[x], y)
	}

	abucket := make([]int, mx)
	for i := 0; i < m; i++ {
		var a int
		fmt.Fscan(in, &a)
		abucket[a]++
	}

	pq := &Heap{}

	ans := 0
	for i := 0; i < mx; i++ {
		for _, j := range bucket[i] {
			heap.Push(pq, j)
		}
		cnt := 0
		for pq.Len() > 0 && cnt < abucket[i] {
			x := heap.Pop(pq).(int)
			if x >= i {
				cnt++
			}
		}
		ans += cnt
	}
	fmt.Println(ans)
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
