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

	var a [200005]int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	en := n + 1

	q := &Heap{}
	ans := 0
	for i := n; i > 0 && en > 1; i-- {
		heap.Push(q, a[i])
		if en > i {
			en -= heap.Pop(q).(int)
			ans++
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
