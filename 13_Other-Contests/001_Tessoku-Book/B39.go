package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, D int
	fmt.Fscan(in, &N, &D)
	tb := make([][]int, D)
	for N > 0 {
		N--
		var x, y int
		fmt.Fscan(in, &x, &y)
		x--
		tb[x] = append(tb[x], y)
	}
	pq := &Heap{}
	ans := 0
	for i := 0; i < D; i++ {
		for _, v := range tb[i] {
			heap.Push(pq, v)
		}
		if pq.Len() != 0 {
			ans += heap.Pop(pq).(int)
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
