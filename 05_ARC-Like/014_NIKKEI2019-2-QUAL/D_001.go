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

	a := make([][]pair, 100005)
	for i := 0; i < m; i++ {
		var l, r, c int
		fmt.Fscan(in, &l, &r, &c)
		l--
		r--
		a[l] = append(a[l], pair{c, r})
	}

	dis := 0
	qu := &Heap{}
	for i := 0; i < n; i++ {
		for qu.Len() != 0 && (*qu)[0].y < i {
			heap.Pop(qu)
		}
		if i != 0 && qu.Len() == 0 {
			fmt.Println(-1)
			return
		}
		if i != 0 {
			dis = (*qu)[0].x
		} else {
			dis = 0
		}
		for _, p := range a[i] {
			heap.Push(qu, pair{p.x + dis, p.y})
		}
	}
	fmt.Println(dis)
}

type pair struct{ x, y int }
type Heap []pair

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i].x < h[j].x }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(pair)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
