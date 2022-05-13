package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)

	t := make([]int, n+1)
	y := make([]int, n+1)
	t[0] = 1
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &t[i], &y[i])
	}

	sum := 0
	q := &Heap{}
	ans := -1 << 60
	for i := n; 0 <= i; i-- {
		if t[i] == 2 {
			if y[i] >= 0 {
				sum += y[i]
			} else {
				heap.Push(q, y[i])
			}
		} else {
			ans = max(ans, sum+y[i])
			k--
		}
		if k < 0 {
			break
		}
		for q.Len() > 0 && q.Len() > k {
			sum += (*q)[0]
			heap.Pop(q)
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
