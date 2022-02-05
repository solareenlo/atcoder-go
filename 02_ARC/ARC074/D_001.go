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
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	lr := [2][100001]int{}
	for j := 0; j < 2; j++ {
		for i := 0; i < 3*n; i++ {
			a[i] *= -1
		}
		p := &Heap{}
		s := 0
		for i := 0; i < n; i++ {
			heap.Push(p, a[i])
			s += a[i]
		}
		lr[j][0] = s
		for i := n; i < 2*n; i++ {
			heap.Push(p, a[i])
			s += a[i]
			s -= (*p)[0]
			lr[j][i+1-n] = s
			heap.Pop(p)
		}
		a = reverseOrderInt(a)
	}
	x := -1 << 60
	for i := 0; i <= n; i++ {
		x = max(x, -lr[0][i]-lr[1][n-i])
	}
	fmt.Println(x)
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func reverseOrderInt(a []int) []int {
	n := len(a)
	res := make([]int, n)
	n = copy(res, a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
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
