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

	a := [1001][1001]int{}
	for i := 1; i < n+1; i++ {
		for j := 1; j < n; j++ {
			fmt.Fscan(in, &a[i][j])
		}
	}

	q := make([]*Heap, 2)
	q[0] = &Heap{}
	q[1] = &Heap{}
	for i := 1; i < n+1; i++ {
		q[0].Push(i)
	}

	t, p := [1001]int{}, [1001]int{}
	now := 1
	for day, tot := 1, n*(n-1)/2; day <= n*(n-1)/2; day++ {
		now ^= 1
		for q[now].Len() > 0 {
			x := (*q[now])[0]
			heap.Pop(q[now])
			t[x]++
			y := a[x][t[x]]
			if p[y] == x {
				tot--
				heap.Push(q[1-now], x)
				heap.Push(q[1-now], y)
				p[x], p[y] = 0, 0
			} else {
				p[x] = y
			}
		}
		if tot == 0 {
			fmt.Println(day)
			return
		}
	}
	fmt.Println(-1)
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
