package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 100001

	var n int
	a := make([]int, N)
	fmt.Fscan(in, &n, &a[1])
	v := make([][]int, N)
	dep := make([]int, N)
	for i := 2; i <= n; i++ {
		var f int
		fmt.Fscan(in, &f, &a[i])
		v[f] = append(v[f], i)
		dep[i] = dep[f] + 1
	}
	q := make([]*HEAP, N)
	for i := range q {
		q[i] = &HEAP{}
	}
	ans := 0
	for i := n; i > 0; i-- {
		for _, j := range v[i] {
			if q[j].Len() > q[i].Len() {
				q[i], q[j] = q[j], q[i]
			}
			for q[j].Len() > 0 {
				heap.Push(q[i], heap.Pop(q[j]))
			}
		}
		if q[i].Len() > 0 && (*q[i])[0] > a[i]+dep[i] {
			ans += heap.Pop(q[i]).(int) - a[i] - dep[i]
			heap.Push(q[i], a[i]+dep[i])
		}
		heap.Push(q[i], a[i]+dep[i])
	}
	fmt.Println(ans)
}

type HEAP []int

func (h HEAP) Len() int            { return len(h) }
func (h HEAP) Less(i, j int) bool  { return h[i] > h[j] }
func (h HEAP) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HEAP) Push(x interface{}) { *h = append(*h, x.(int)) }

func (h *HEAP) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
