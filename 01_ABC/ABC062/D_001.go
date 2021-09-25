package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, 3*n)
	for i := 0; i < 3*n; i++ {
		fmt.Fscan(in, &a[i])
	}

	q := &Heap{}
	sum := 0
	maxi := make([]int, 3*n)
	for i := 0; i < 3*n; i++ {
		heap.Push(q, -a[i])
		sum += a[i]
		if n < q.Len() {
			sum += heap.Pop(q).(int)
		}
		maxi[i] = sum
	}

	q = &Heap{}
	sum = 0
	mini := make([]int, 3*n)
	for i := 3*n - 1; i >= 0; i-- {
		heap.Push(q, a[i])
		sum += a[i]
		if n < q.Len() {
			sum -= heap.Pop(q).(int)
		}
		mini[i] = sum
	}

	res := -int(1e16)
	for i := n; i <= 2*n; i++ {
		res = max(res, maxi[i-1]-mini[i])
	}
	fmt.Println(res)
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
