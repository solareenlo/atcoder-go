package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	type pair struct {
		x, y int
	}

	var n, m int
	fmt.Fscan(in, &n, &m)
	var P [100]pair
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &P[i].x, &P[i].y)
	}
	G := make([]map[int]int, 100)
	for i := range G {
		G[i] = make(map[int]int)
	}
	for i := 0; i < m; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		a--
		b--
		G[a][b] = c
		G[b][a] = c
	}
	Q := &HeapTuple{}
	heap.Push(Q, tuple{0, 0, 0})
	var S [100][100]bool
	for Q.Len() > 0 {
		tmp := heap.Pop(Q).(tuple)
		c, y, x := tmp.x, tmp.y, tmp.z
		if S[y][x] {
			continue
		}
		S[y][x] = true
		if y == 1 {
			fmt.Println(c)
			return
		}
		for z, d := range G[y] {
			p, q := P[x].x, P[x].y
			r, s := P[y].x, P[y].y
			t, u := P[z].x, P[z].y
			if (p-r)*(t-r)+(q-s)*(u-s) <= 0 {
				heap.Push(Q, tuple{c + d, z, y})
			}
		}
	}
	fmt.Println(-1)
}

type tuple struct {
	x, y, z int
}

type HeapTuple []tuple

func (h HeapTuple) Len() int            { return len(h) }
func (h HeapTuple) Less(i, j int) bool  { return h[i].x < h[j].x }
func (h HeapTuple) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapTuple) Push(x interface{}) { *h = append(*h, x.(tuple)) }

func (h *HeapTuple) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
